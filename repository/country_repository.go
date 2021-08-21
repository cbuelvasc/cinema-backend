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

type CountryRepository interface {
	GetAllCountries(ctx context.Context, page int64, limit int64) (*model.PagedCountry, error)
	GetCountryById(ctx context.Context, id string) (*model.Country, error)
	SaveCountry(ctx context.Context, country *model.Country) (*model.Country, error)
	UpdateCountry(ctx context.Context, id string, country *model.Country) (*model.Country, error)
	DeleteCountry(ctx context.Context, id string) error
}

type countryRepositoryImpl struct {
	Connection *mongo.Database
}

func NewCountryRepository(Connection *mongo.Database) CountryRepository {
	return &countryRepositoryImpl{Connection: Connection}
}

func (countryRepository *countryRepositoryImpl) GetAllCountries(ctx context.Context, page int64, limit int64) (*model.PagedCountry, error) {
	var countries []model.Country

	filter := bson.M{}

	collection := countryRepository.Connection.Collection("countries")

	projection := bson.D{
		{"id", 1},
		{"name", 1},
		{"states", 1},
		{"created_at", 1},
		{"updated_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&countries).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedCountry{
		Data:     countries,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (countryRepository *countryRepositoryImpl) GetCountryById(ctx context.Context, id string) (*model.Country, error) {
	var country model.Country
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	err := countryRepository.Connection.Collection("countries").FindOne(ctx, filter).Decode(&country)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Country", "id", id)
	}
	return &country, nil
}

func (countryRepository *countryRepositoryImpl) SaveCountry(ctx context.Context, country *model.Country) (*model.Country, error) {
	country.ID = primitive.NewObjectID()

	_, err := countryRepository.Connection.Collection("countries").InsertOne(ctx, country)
	if err != nil {
		return nil, err
	}

	return country, nil
}

func (countryRepository *countryRepositoryImpl) UpdateCountry(ctx context.Context, id string, country *model.Country) (*model.Country, error) {
	country.ID = primitive.NewObjectID()

	_, err := countryRepository.Connection.Collection("countries").InsertOne(ctx, country)
	if err != nil {
		return nil, err
	}

	return country, nil
}

func (countryRepository *countryRepositoryImpl) DeleteCountry(ctx context.Context, id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	result, err := countryRepository.Connection.Collection("countries").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return exception.NotFoundRequestException(fmt.Sprintf("Country not found with id: %s", id))
	}

	return nil
}
