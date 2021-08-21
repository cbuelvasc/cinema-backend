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

type MovieRepository interface {
	GetAllMovies(ctx context.Context, page int64, limit int64) (*model.PagedMovie, error)
	GetMovie(ctx context.Context, id string) (*model.Movie, error)
	SaveMovie(ctx context.Context, movie *model.Movie) (*model.Movie, error)
	UpdateMovie(ctx context.Context, id string, movie *model.Movie) (*model.Movie, error)
	DeleteMovie(ctx context.Context, id string, movieId string) error
}

type movieRepositoryImpl struct {
	Connection *mongo.Database
}

func NewMovieRepository(Connection *mongo.Database) MovieRepository {
	return &movieRepositoryImpl{Connection: Connection}
}

func (movieRepository *movieRepositoryImpl) GetAllMovies(ctx context.Context, page int64, limit int64) (*model.PagedMovie, error) {
	var movies []model.Movie

	filter := bson.M{
	}

	collection := movieRepository.Connection.Collection("movies")

	projection := bson.D{
		{"id", 1},
		{"title", 1},
		{"format", 1},
		{"releaseYear", 1},
		{"releaseMonth", 1},
		{"releaseDay", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&movies).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedMovie{
		Data:     movies,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (movieRepository *movieRepositoryImpl) GetMovie(ctx context.Context, id string) (*model.Movie, error) {
	var movie model.Movie
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	err := movieRepository.Connection.Collection("movies").FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Movie", "id", id)
	}
	return &movie, nil
}

func (movieRepository *movieRepositoryImpl) SaveMovie(ctx context.Context, movie *model.Movie) (*model.Movie, error) {
	movie.ID = primitive.NewObjectID()

	_, err := movieRepository.Connection.Collection("movies").InsertOne(ctx, movie)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (movieRepository *movieRepositoryImpl) UpdateMovie(ctx context.Context, id string, movie *model.Movie) (*model.Movie, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	registry := make(map[string]interface{})
	if len(movie.Title) > 0 {
		registry["title"] = movie.Title
	}
	if len(movie.Format) > 0 {
		registry["format"] = movie.Format
	}
	registry["releaseYear"] = movie.ReleaseYear
	registry["releaseMonth"] = movie.ReleaseMonth
	registry["releaseDay"] = movie.ReleaseDay

	filter := bson.M{"_id": bson.M{"$eq": objectId}}

	updateString := bson.M{
		"$set": registry,
	}

	result, err := movieRepository.Connection.Collection("movies").UpdateOne(ctx, filter, updateString)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("Movie", "id", id)
	}

	movie.ID = objectId
	return movie, nil
}

func (movieRepository *movieRepositoryImpl) DeleteMovie(ctx context.Context, id string, movieId string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":    objectId,
		"movieId": movieId,
	}

	result, err := movieRepository.Connection.Collection("movies").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return exception.NotFoundRequestException(fmt.Sprintf("Movie not found with id: %s and movieId: %s", id, movieId))
	}

	return nil
}


