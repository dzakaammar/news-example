package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	Database DatabaseConfig `yaml:"database"`
	App      AppConfig      `yaml:"app"`
}

var (
	Conf     *Configuration
	syncOnce sync.Once
)

func Init(path string) error {
	if Conf == nil {
		errChan := make(chan error, 1)
		syncOnce.Do(func() {
			viper.SetConfigFile(path)
			err := viper.ReadInConfig()
			if err != nil {
				fmt.Errorf("Error : %v", err)
				errChan <- err
				return
			}

			err = viper.Unmarshal(&Conf)
			if err != nil {
				fmt.Errorf("Error : %v", err)
				errChan <- err
				return
			}

			errChan <- nil
		})

		err := <-errChan
		if err != nil {
			return err
		}
		close(errChan)
	}

	return nil
}
