package models

type UserObject struct {
	Name   string `bson:"name" json:"name"`
	Age    int    `bson:"age" json:"age"`
	Email  string `bson:"email" json:"email"`
	Maried bool   `bson:"maried" json:"maried"`
}
