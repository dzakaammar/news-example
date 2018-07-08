package repository

import (
	"errors"
	"fmt"

	"github.com/dzakaammar/news-example/domain"

	"github.com/dzakaammar/news-example/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	NotFoundErr       = errors.New("Cannot Find The Data")
	CannotRetrieveErr = errors.New("Cannot Retrieve Data")
	CannotStoreErr    = errors.New("Cannot Store The Data")
	CannotRemoveErr   = errors.New("Cannot Remove The Data")
	CannotUpdateErr   = errors.New("Cannot Update The Data")
)

func Init() (*gorm.DB, error) {
	var arg string

	switch config.Conf.Database.Driver {
	case "mysql":
		arg = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", config.Conf.Database.Username, config.Conf.Database.Password, config.Conf.Database.Host, config.Conf.Database.Name)
	case "sqlite3":
		arg = "./news-example.db"
	default:
		return nil, errors.New("Unknown Driver")
	}

	db, err := gorm.Open(config.Conf.Database.Driver, arg)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func Migrate(db *gorm.DB) error {
	defer fmt.Println("Migrating is done.")

	err := handleMigrate(db, &domain.News{}, &domain.Topic{}, &domain.NewsTopic{})
	if err != nil {
		return err
	}

	if err := seeding(db); err != nil {
		return err
	}

	return nil
}

func handleMigrate(db *gorm.DB, models ...interface{}) error {
	fmt.Println("Migrating tables to database...")
	for _, model := range models {
		if db.HasTable(model) {
			if err := db.DropTableIfExists(model).Error; err != nil {
				return err
			}
		}

		if err := db.AutoMigrate(model).Error; err != nil {
			return err
		}
	}

	return nil
}

func seeding(db *gorm.DB) error {
	defer fmt.Println("Seeding is done")

	fmt.Println("Starting to seeding database...")
	//Create news
	txn := db.Begin()

	news := &domain.News{
		Title:   "Contoh 1",
		Status:  "draft",
		Content: "Contoh 1",
	}

	err := handleSeeding(txn, news, &domain.Topic{ID: 1, Name: "politik"}, &domain.Topic{ID: 2, Name: "ekonomi"}, &domain.Topic{ID: 3, Name: "sosial"}, &domain.Topic{ID: 4, Name: "budaya"})
	if err != nil {
		txn.Rollback()
		return err
	}

	news2 := &domain.News{
		Title:   "Contoh 2",
		Status:  "published",
		Content: "Contoh 2",
	}

	err = handleSeeding(txn, news2, &domain.Topic{ID: 5, Name: "pendidikan"}, &domain.Topic{ID: 1, Name: "politik"}, &domain.Topic{ID: 6, Name: "teknologi"}, &domain.Topic{ID: 4, Name: "budaya"})
	if err != nil {
		txn.Rollback()
		return err
	}

	txn.Commit()
	return nil
}

func handleSeeding(txn *gorm.DB, news *domain.News, topics ...*domain.Topic) error {
	for _, v := range topics {
		if err := txn.FirstOrCreate(v).Error; err != nil {
			return err
		}
	}

	if err := txn.Create(news).Error; err != nil {
		return err
	}

	if err := txn.Model(news).Association("Topics").Append(topics).Error; err != nil {
		return err
	}
	return nil
}
