package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Schedule struct {
	*ScheduleInput `bson:",inline"`
	ID             primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type ScheduleInput struct {
	Date          time.Time `json:"date,omitempty" xml:"date,omitempty" bson:"date"`
	RoomId        string    `json:"roomId,omitempty" xml:"roomId,omitempty" bson:"roomId" validate:"required"`
	MovieId       string    `json:"movieId,omitempty" xml:"movieId,omitempty" bson:"movieId" validate:"required"`
	SeatsEmpty    []string  `json:"seatsEmpty,omitempty" xml:"seatsEmpty,omitempty" bson:"seatsEmpty"`
	SeatsOccupied []string  `json:"seatsOccupied,omitempty" xml:"seatsOccupied,omitempty" bson:"seatsOccupied"`
	CreatedAt     time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedSchedule struct {
	Data     []Schedule                     `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
