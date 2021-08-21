package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Country struct {
	*CountryInput `bson:",inline"`
	ID            primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type CountryInput struct {
	Name      string    `json:"name,omitempty" xml:"name,omitempty" bson:"name" validate:"required"`
	States    []string  `json:"states,omitempty" xml:"states,omitempty" bson:"states"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedCountry struct {
	Data     []Country                      `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
