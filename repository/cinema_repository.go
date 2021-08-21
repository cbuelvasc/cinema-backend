package repository

import (
	"context"
	"github.com/cbuelvasc/cinema-backend/exception"
	"github.com/cbuelvasc/cinema-backend/model"
	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CinemaRepository interface {
	GetAllCinemas(ctx context.Context, page int64, limit int64) (*model.PagedCinema, error)
	GetCinemaById(ctx context.Context, id string) (*model.Cinema, error)
	GetCinemaByCity(ctx context.Context, cityId string) (*model.Cinema, error)
	SaveCinema(ctx context.Context, cinema *model.Cinema) (*model.Cinema, error)
	UpdateCinema(ctx context.Context, id string, cinemaId *model.Cinema) (*model.Cinema, error)
	DeleteCinema(ctx context.Context, id string, cinemaId string) error
}

type cinemaRepositoryImpl struct {
	Connection *mongo.Database
}

func NewCinemaRepository(Connection *mongo.Database) CinemaRepository {
	return &cinemaRepositoryImpl{Connection: Connection}
}

func (cinemaRepository *cinemaRepositoryImpl) GetAllCinemas(ctx context.Context, page int64, limit int64) (*model.PagedCinema, error) {
	var cinemas []model.Cinema

	filter := bson.M{
	}

	collection := cinemaRepository.Connection.Collection("cinemas")

	projection := bson.D{
		{"id", 1},
		{"name", 1},
		{"cityId", 1},
		{"premieres", 1},
		{"rooms", 1},
		{"created_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&cinemas).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedCinema{
		Data:     cinemas,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (cinemaRepository *cinemaRepositoryImpl) GetCinemaById(ctx context.Context, id string) (*model.Cinema, error) {
	var cinema model.Cinema
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	err := cinemaRepository.Connection.Collection("cinemas").FindOne(ctx, filter).Decode(&cinema)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Cinema", "id", id)
	}
	return &cinema, nil
}

func (cinemaRepository *cinemaRepositoryImpl) GetCinemaByCity(ctx context.Context, cityId string) (*model.Cinema, error) {
	panic("implement me")
}

func (cinemaRepository *cinemaRepositoryImpl) SaveCinema(ctx context.Context, cinema *model.Cinema) (*model.Cinema, error) {
	panic("implement me")
}

func (cinemaRepository *cinemaRepositoryImpl) UpdateCinema(ctx context.Context, id string, cinemaId *model.Cinema) (*model.Cinema, error) {
	panic("implement me")
}

func (cinemaRepository *cinemaRepositoryImpl) DeleteCinema(ctx context.Context, id string, cinemaId string) error {
	panic("implement me")
}