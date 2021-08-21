package repository

import (
	"context"
	"fmt"
	"github.com/cbuelvasc/cinema-backend/exception"
	"github.com/cbuelvasc/cinema-backend/model"
	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CityRepository interface {
	GetAllCities(ctx context.Context, page int64, limit int64) (*model.PagedCity, error)
	GetCityById(ctx context.Context, id string) (*model.City, error)
	SaveCity(ctx context.Context, city *model.City) (*model.City, error)
	UpdateCity(ctx context.Context, id string, city *model.City) (*model.City, error)
	DeleteCity(ctx context.Context, id string, cityId string) error
}

type cityRepositoryImpl struct {
	Connection *mongo.Database
}

func NewCityRepository(Connection *mongo.Database) CityRepository {
	return &cityRepositoryImpl{Connection: Connection}
}

func (cityRepository *cityRepositoryImpl) GetAllCities(ctx context.Context, page int64, limit int64) (*model.PagedCity, error) {
	var cities []model.City

	filter := bson.M{
	}

	collection := cityRepository.Connection.Collection("cities")

	projection := bson.D{
		{"id", 1},
		{"name", 1},
		{"countryId", 1},
		{"cities", 1},
		{"created_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&cities).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedCity{
		Data:     cities,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (cityRepository *cityRepositoryImpl) GetCityById(ctx context.Context, id string) (*model.City, error) {
	var city model.City
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	err := cityRepository.Connection.Collection("cities").FindOne(ctx, filter).Decode(&city)
	if err != nil {
		return nil, exception.ResourceNotFoundException("City", "id", id)
	}
	return &city, nil
}

func (cityRepository *cityRepositoryImpl) SaveCity(ctx context.Context, city *model.City) (*model.City, error) {
	city.ID = primitive.NewObjectID()

	_, err := cityRepository.Connection.Collection("cities").InsertOne(ctx, city)
	if err != nil {
		return nil, err
	}

	return city, nil
}

func (cityRepository *cityRepositoryImpl) UpdateCity(ctx context.Context, id string, city *model.City) (*model.City, error) {
	city.ID = primitive.NewObjectID()

	_, err := cityRepository.Connection.Collection("cities").InsertOne(ctx, city)
	if err != nil {
		return nil, err
	}

	return city, nil
}

func (cityRepository *cityRepositoryImpl) DeleteCity(ctx context.Context, id string, cityId string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":    objectId,
		"cityId": cityId,
	}

	result, err := cityRepository.Connection.Collection("cities").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return exception.NotFoundRequestException(fmt.Sprintf("City not found with id: %s and cityId: %s", id, cityId))
	}

	return nil
}