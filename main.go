package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"

	"github.com/dzakaammar/news-example/service"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/dzakaammar/news-example/repository"

	"github.com/dzakaammar/news-example/config"

	"github.com/go-kit/kit/log"
)

var (
	c          = flag.String("config", "./config.yaml", "path of configuration file")
	migrate    = flag.Bool("migrate", false, "True to run database migrate, false to not.")
	Logger     log.Logger
	httpLogger log.Logger
)

func main() {
	flag.Parse()
	if *c == "" {
		flag.PrintDefaults()
	}

	if err := config.Init(*c); err != nil {
		panic(err)
	}

	Logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	httpLogger = log.With(Logger, "component", "http")

	d, err := repository.Init()
	if err != nil {
		panic(err)
	}

	if *migrate {
		err = repository.Migrate(d)
		if err != nil {
			panic(err)
		}
	}

	var (
		newsRepository   = repository.NewNewsRepository(d)
		topicsRepository = repository.NewTopicRepository(d)
	)

	var s service.Service
	s = service.NewService(newsRepository, topicsRepository)
	s = service.NewLoggingService(log.With(Logger, "component", "service"), s)
	s = service.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_count",
			Help:      "Number of requests recieved",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
		s,
	)

	srv := chi.NewRouter()
	srv.Use(cors)
	srv.Mount("/", service.MakeHandler(s, httpLogger))
	srv.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:           ":" + port,
		Handler:        srv,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	errs := make(chan error, 2)
	go func() {
		Logger.Log("transport", "http", "address", port, "msg", "listening")
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	Logger.Log("terminated", <-errs)

}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		)
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
