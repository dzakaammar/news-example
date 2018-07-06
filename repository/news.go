package repository

import (
	"github.com/dzakaammar/news-example/domain"
	"github.com/jinzhu/gorm"
)

type newsRepository struct {
	*gorm.DB
}

func NewNewsRepository(db *gorm.DB) domain.NewsRepository {
	return &newsRepository{db}
}

func (n *newsRepository) List(filterBy map[string][]string) ([]*domain.News, error) {
	var news []*domain.News

	db := n.DB.Preload("Topics")
	if len(filterBy) > 0 {
		for keys, v := range filterBy {
			switch keys {
			case "status":
				db = filterByStatus(db, v)
			case "topics":
				db = filterByTopics(db, v)
			default:
			}
		}
	}

	if err := db.Find(&news).Error; err != nil {
		return nil, CannotRetrieveErr
	}

	return news, nil
}

func filterByStatus(db *gorm.DB, status []string) *gorm.DB {
	return db.Where("status in (?)", status)
}

func filterByTopics(db *gorm.DB, topics []string) *gorm.DB {
	return db.Preload("Topics", "name in (?)", topics)
}

func (n *newsRepository) Find(id uint) (*domain.News, error) {
	news := &domain.News{}

	if err := n.DB.Preload("Topics").Where("id = ?", id).First(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}

func (n *newsRepository) Store(news *domain.News, topics []string) error {
	var t []*domain.Topic
	txn := n.DB.Begin()

	for _, v := range topics {
		tt := &domain.Topic{}
		if err := txn.FirstOrCreate(&tt, &domain.Topic{Name: v}).Error; err != nil {
			return CannotRetrieveErr
		}

		t = append(t, tt)
	}
	if err := txn.Create(news).Error; err != nil {
		txn.Rollback()
		return CannotStoreErr
	}

	if err := txn.Model(news).Association("Topics").Append(t).Error; err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (n *newsRepository) Remove(news *domain.News) error {
	txn := n.DB.Begin()

	if err := txn.Model(&news).Update("status", "deleted").Error; err != nil {
		txn.Rollback()
		return CannotUpdateErr
	}

	if err := txn.Delete(news).Error; err != nil {
		txn.Rollback()
		return CannotRemoveErr
	}

	txn.Commit()
	return nil
}
