package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type City struct {
	*CityInput `bson:",inline"`
	ID         primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type CityInput struct {
	Name      string    `json:"name,omitempty" xml:"name,omitempty" bson:"name" validate:"required"`
	StateId   string    `json:"stateId,omitempty" xml:"stateId,omitempty" bson:"stateId" validate:"required"`
	Cinemas   []string  `json:"cinemas,omitempty" xml:"cinemas,omitempty" bson:"cinemas" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedCity struct {
	Data     []City                         `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
