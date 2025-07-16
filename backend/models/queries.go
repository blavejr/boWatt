package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryHistory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Query     string             `bson:"query"`
	FileHash  string             `bson:"file_hash"`
	UserId    primitive.ObjectID `bson:"user_id"`
	Timestamp int64              `bson:"timestamp"`
}
