package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileMetadata struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FileHash  string             `bson:"file_hash"`
	UserId    primitive.ObjectID `bson:"user_id"`
	Name      string             `bson:"name"`
	Content   string             `bson:"content"`
	Timestamp int64              `bson:"timestamp"`
}
