package model

import (
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	*MovieInput `bson:",inline"`
	ID          primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type MovieInput struct {
	Title        string    `json:"title,omitempty" xml:"title,omitempty" bson:"title" validate:"required"`
	Format       string    `json:"format,omitempty" xml:"format,omitempty" bson:"format" validate:"required"`
	ReleaseYear  int       `json:"releaseYear,omitempty" xml:"releaseYear,omitempty" bson:"releaseYear"`
	ReleaseMonth int       `json:"releaseMonth,omitempty" xml:"releaseMonth,omitempty" bson:"releaseMonth"`
	ReleaseDay   int       `json:"releaseDay,omitempty" xml:"releaseDay,omitempty" bson:"releaseDay"`
	CreatedAt    time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type PagedMovie struct {
	Data     []Movie                        `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
