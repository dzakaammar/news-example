package domain

import "github.com/jinzhu/gorm"

type News struct {
	gorm.Model
	Title   string   `gorm:"type:varchar(255);not null" json:"title"`
	Status  string   `gorm:"type:varchar(20);not null" json:"status"`
	Content string   `gorm:"type:text" json:"content"`
	Topics  []*Topic `gorm:"many2many:news_topic;association_jointable_foreignkey:topic_id;jointable_foreignkey:news_id;" json:"topics,omitempty"`
}

func (News) TableName() string {
	return "news"
}

func NewNews(title, status, content string) *News {
	return &News{
		Title:   title,
		Content: content,
		Status:  status,
	}
}

type NewsRepository interface {
	List(FilterBy map[string][]string) ([]*News, error)
	Find(ID uint) (*News, error)
	Store(news *News, Topics []string) error
	Remove(news *News) error
}
