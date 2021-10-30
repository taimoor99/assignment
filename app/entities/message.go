package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Messages struct {
		ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
		Message   string             `bson:"message" json:"message"`
		CreatedAt time.Time          `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	}
	MessagesCreateInput struct {
		Message string `bson:"message" json:"message"`
	}
	MessageDetails struct {
		IsPalindrome bool `bson:"is_palindrome" json:"is_palindrome"`
		Messages
	}
)

func (Messages) MessagesCollection() string {
	return "messages"
}
