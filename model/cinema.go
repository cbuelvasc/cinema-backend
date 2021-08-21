package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cinema struct {
	*CinemaInput `bson:",inline"`
	ID           primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type CinemaInput struct {
	Name      string    `json:"name,omitempty" xml:"name,omitempty" bson:"name" validate:"required"`
	CityId    string    `json:"cityId,omitempty" xml:"cityId,omitempty" bson:"cityId" validate:"required"`
	Premieres []string  `json:"premieres,omitempty" xml:"premieres,omitempty" bson:"premieres"`
	Rooms     []string  `json:"rooms,omitempty" xml:"rooms,omitempty" bson:"rooms"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedCinema struct {
	Data     []Cinema                       `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
