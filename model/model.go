package model

import {
	"go.mongodb.org/mongo-driver/v2/bson/primtive"
}

type Netflix struct {
	ID primtive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string	`json:"movie,omitempty"`
	Watched bool	`json:"watched,omitempty"`
}