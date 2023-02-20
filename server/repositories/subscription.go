package repositories

import (
	"wayshub/models"

	"gorm.io/gorm"
)

type SubscribeRepository interface {
	FindSubscribes(ChannelID int) ([]models.Subscription, error)
	GetSubscribe(ChannelID int) (models.Subscription, error)

	GetSubscribeByOther(OtherID int) (models.Subscription, error)

	GetOther(ID int) (models.Channel, error)

	CreateSubscribe(subscribe models.Subscription) (models.Subscription, error)

	DeleteSubscribe(subscribe models.Subscription) (models.Subscription, error)

	GetSubscription(ChannelID int) ([]models.Subscription, error)
}

func RepositorySubscribe(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetOther(ID int) (models.Channel, error) {
	var channel models.Channel
	err := r.db.Preload("Subscription.Channel").Preload("Video.Channel").Preload("Video").First(&channel, ID).Error

	return channel, err
}

func (r *repository) FindSubscribes(ChannelID int) ([]models.Subscription, error) {
	var subscribes []models.Subscription
	err := r.db.Preload("Channel").Where("channel_id = ?", ChannelID).Find(&subscribes).Error

	return subscribes, err
}

func (r *repository) GetSubscribe(ChannelID int) (models.Subscription, error) {
	var subscribe models.Subscription
	err := r.db.Preload("Channel").Where("channel_id = ?", ChannelID).First(&subscribe).Error

	return subscribe, err
}

func (r *repository) GetSubscribeByOther(OtherID int) (models.Subscription, error) {
	var subscribe models.Subscription
	err := r.db.Preload("Channel").Where("other_id = ?", OtherID).First(&subscribe).Error

	return subscribe, err
}

func (r *repository) CreateSubscribe(subscribe models.Subscription) (models.Subscription, error) {
	err := r.db.Create(&subscribe).Error

	return subscribe, err
}

// func (r *repository) DeleteSubscribe(OtherID int) (models.Subscription, error) {
// 	var subscribe models.Subscription
// 	err := r.db.Where("other_id = ?", OtherID).Delete(&subscribe).Error

// 	return subscribe, err
// }

func (r *repository) DeleteSubscribe(subscribe models.Subscription) (models.Subscription, error) {
	err := r.db.Delete(&subscribe).Error

	return subscribe, err
}

func (r *repository) GetSubscription(ChannelID int) ([]models.Subscription, error) {
	var subscription []models.Subscription
	err := r.db.Preload("Channel").Preload("OtherVideo.Channel").Find(&subscription).Error

	return subscription, err
}

func (r *repository) GetOtherVid(ID int) (models.Video, error) {
	var channel models.Video
	err := r.db.First(&channel, ID).Error

	return channel, err
}

func (r *repository) CreateSubscription(subscribe models.Subscription) (models.Subscription, error) {
	err := r.db.Create(&subscribe).Error

	return subscribe, err
}
