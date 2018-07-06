package domain

import "time"

type Topic struct {
	ID        uint
	Name      string     `gorm:"type:varchar(50);unique;not null" json:"name"`
	News      []*News    `gorm:"many2many:news_topic;association_jointable_foreignkey:news_id;jointable_foreignkey:topic_id;" json:"topics,omitempty"`
	DeletedAt *time.Time `sql:"index"`
}

func (Topic) TableName() string {
	return "topic"
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name: name,
	}
}

type TopicRepository interface {
	List() ([]*Topic, error)
	Find(ID uint) (*Topic, error)
	Store(topic *Topic) error
	Remove(topic *Topic) error
}
