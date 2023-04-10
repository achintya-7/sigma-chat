package models

type Chat struct {
	ChatId       string `json:"chatId" bson:"chatId"`
	Profile1     string `json:"profile1" bson:"profile1"`
	Profile1Name string `json:"profile1name" bson:"profile1name"`
	Profile2     string `json:"profile2" bson:"profile2"`
	Profile2Name string `json:"profile2name" bson:"profile2name"`
	Message      string `json:"message" bson:"message"`
	CreatedOn    int32  `json:"createdOn" bson:"createdOn"`
	Sender       string `json:"sender" bson:"sender"`
}
