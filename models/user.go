package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Fullname string
	Email    string
	Password string
}
