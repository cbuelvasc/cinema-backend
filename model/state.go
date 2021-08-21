package model

import (
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type State struct {
	*StateInput `bson:",inline"`
	ID          primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type StateInput struct {
	Name      string    `json:"name,omitempty" xml:"name,omitempty" bson:"name" validate:"required"`
	CountryId string    `json:"countryId,omitempty" xml:"countryId,omitempty" bson:"countryId" validate:"required"`
	Cities    []string  `json:"cities,omitempty" xml:"cities,omitempty" bson:"cities"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type StateResponse struct {
	Name      string    `json:"name,omitempty" xml:"name,omitempty" bson:"name" validate:"required"`
	Country   Country   `json:"country,omitempty" xml:"country,omitempty" bson:"country"`
	Cities    []string  `json:"cities,omitempty" xml:"cities,omitempty" bson:"cities"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedState struct {
	Data     []State                        `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
