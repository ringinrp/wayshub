package models

type Subscription struct {
	ID         int             `json:"id" gorm:"primary_key:auto_increment"`
	OtherID    int             `json:"other_id"`
	OtherName  string          `json:"otherName" gorm:"type: varchar(255)"`
	OtherPhoto string          `json:"otherPhoto" gorm:"type: varchar(255)"`
	ChannelID  int             `json:"channel_id"`
	Channel    ChannelResponse `json:"channels" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SubscriptionResponse struct {
	ID        int             `json:"id"`
	ChannelID int             `json:"channel_id"`
	Channel   ChannelResponse `json:"-"`
}

type SubscriptionChannelResponse struct {
	ID        int             `json:"id"`
	ChannelID int             `json:"channel_id"`
	Channel   ChannelResponse `json:"-"`
}

func (SubscriptionResponse) TableName() string {
	return "subscribers"
}

func (SubscriptionChannelResponse) TableName() string {
	return "subscribers"
}
