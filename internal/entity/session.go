package entity

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Session struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID       string        `json:"userID" bson:"user_id"`
	RefreshToken string        `json:"refreshToken" bson:"refresh_token"`
	ExpiresIn    time.Time     `json:"expiresIn" bson:"expires_in"`
}
