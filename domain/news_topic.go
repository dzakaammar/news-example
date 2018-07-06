package domain

type NewsTopic struct {
	NewsID  uint `gorm:"not null"`
	TopicID uint `gorm:"not null"`

	News  *News  `gorm:"ForeignKey:NewsID"`
	Topic *Topic `gorm:"ForeignKey:TopicID"`
}

func (NewsTopic) TableName() string {
	return "news_topic"
}
