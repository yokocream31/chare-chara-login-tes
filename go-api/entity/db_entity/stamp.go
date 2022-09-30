package db_entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Stamp struct {
	StampId primitive.ObjectID `json:"stampId" bson:"_id"`
	StampName string `json:"stampName" bson:"stampName"`
	StampImg string `json:"stampImg" bson:"stampImg"` // スタンプのicon path
	Status string `json:"status" bson:"status"`
	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt" bson:"deletedAt"`
}

// Auto Create CreatedAt, UpdatedAt
func (s *Stamp) MarshalBSON() ([]byte, error) {
    if s.CreatedAt.IsZero() {
        s.CreatedAt = time.Now()
    }
    s.UpdatedAt = time.Now()
    
    type my Stamp
    return bson.Marshal((*my)(s))
}
