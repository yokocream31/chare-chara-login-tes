package db_entity

import (
	"time"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bear struct {
	BearId primitive.ObjectID `json:"bearId" bson:"_id"`
	BearName string `json:"bearName" bson:"bearName"`
	BearIcon string `json:"bearIcon" bson:"bearIcon"`
	Detail string `json:"detail" bson:"detail"`
	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt" bson:"deletedAt"`
}

// Auto Create CreatedAt, UpdatedAt
func (b *Bear) MarshalBSON() ([]byte, error) {
    if b.CreatedAt.IsZero() {
        b.CreatedAt = time.Now()
    }
    b.UpdatedAt = time.Now()
    
    type my Bear
    return bson.Marshal((*my)(b))
}

// BearToneのSubCollection
type Talk struct {
	Id uint `json:"id" bson:"id"`
	Response string `json:"response" bson:"response"`
	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt" bson:"deletedAt"`
}

// Auto Create CreatedAt, UpdatedAt
func (t *Talk) MarshalBSON() ([]byte, error) {
    if t.CreatedAt.IsZero() {
        t.CreatedAt = time.Now()
    }
    t.UpdatedAt = time.Now()
    
    type my Talk
    return bson.Marshal((*my)(t))
}

type BearTone struct {
	ToneId primitive.ObjectID `json:"toneId" bson:"_id"`
	ToneName string `json:"toneName" bson:"toneName"`
	Detail string `json:"detail" bson:"detail"`
	Talk []Talk `json:"talk" bson:"talk"` // SubCollection
	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt" bson:"deletedAt"`
}

// Auto Create CreatedAt, UpdatedAt
func (b *BearTone) MarshalBSON() ([]byte, error) {
    if b.CreatedAt.IsZero() {
        b.CreatedAt = time.Now()
    }
    b.UpdatedAt = time.Now()
    
    type my BearTone
    return bson.Marshal((*my)(b))
}

// CommunicationのSubCollection
type Messages struct {
	Id uint `json:"id" bson:"id"`
	Text string `json:"text" bson:"text"`
	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
}

// Auto Create CreatedAt, UpdatedAt
func (m *Messages) MarshalBSON() ([]byte, error) {
    if m.CreatedAt.IsZero() {
        m.CreatedAt = time.Now()
    }
    type my Messages
    return bson.Marshal((*my)(m))
}

type Communication struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	UserId User `json:"userId" bson:"userId"`
	Messages []Messages `json:"messages" bson:"messages"`
	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt" bson:"deletedAt"`
}

// Auto Create CreatedAt, UpdatedAt
func (c *Communication) MarshalBSON() ([]byte, error) {
    if c.CreatedAt.IsZero() {
        c.CreatedAt = time.Now()
    }
    c.UpdatedAt = time.Now()
    
    type my Communication
    return bson.Marshal((*my)(c))
}
