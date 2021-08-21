package model

import (
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	*TweetInput `bson:",inline"`
	ID          primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type TweetInput struct {
	UserId    string    `json:"userId,omitempty" xml:"userId,omitempty" bson:"userId" validate:"required"`
	Message   string    `json:"message,omitempty" xml:"message,omitempty" bson:"message" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedTweet struct {
	Data     []Tweet                        `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
