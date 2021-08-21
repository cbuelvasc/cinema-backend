package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Room struct {
	*RoomInput `bson:",inline"`
	ID         primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type RoomInput struct {
	Name      string    `json:"name,omitempty" xml:"name,omitempty" bson:"name" validate:"required"`
	Capacity  string    `json:"capacity,omitempty" xml:"capacity,omitempty" bson:"capacity" validate:"required"`
	Format    string    `json:"format,omitempty" xml:"format,omitempty" bson:"format" validate:"required"`
	CinemaId  string    `json:"cinemaId,omitempty" xml:"cinemaId,omitempty" bson:"cinemaId" validate:"required"`
	Schedules []string  `json:"schedules,omitempty" xml:"schedules,omitempty" bson:"schedules"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedRoom struct {
	Data     []Room                         `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
