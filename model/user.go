package model

import (
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	*UserInput `bson:",inline"`
	ID         primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type UserInput struct {
	Name      string    `json:"name,omitempty" xml:"name,omitempty" bson:"name" validate:"required"`
	Lastname  string    `json:"lastname,omitempty" xml:"lastname,omitempty" bson:"lastname" validate:"required"`
	BirthDate time.Time `json:"birthDate,omitempty" xml:"birthDate,omitempty" bson:"birthDate"`
	Email     string    `json:"email" xml:"email" bson:"email" validate:"required,email"`
	Password  string    `json:"password,omitempty" xml:"password,omitempty" bson:"password"`
	Avatar    string    `json:"avatar,omitempty" xml:"avatar,omitempty" bson:"avatar"`
	Banner    string    `json:"banner,omitempty" xml:"banner,omitempty" bson:"banner"`
	Biography string    `json:"biography,omitempty" xml:"biography,omitempty" bson:"biography"`
	Location  string    `json:"location,omitempty" xml:"location,omitempty" bson:"location"`
	WebSite   string    `json:"webSite,omitempty" xml:"webSite,omitempty" bson:"webSite"`
	CreatedAt time.Time `json:"created_at,omitempty" xml:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty" bson:"updated_at"`
}

type SignInInput struct {
	Email    string `json:"email" xml:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" bson:"password" validate:"required"`
}

type PagedUser struct {
	Data     []User                         `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
