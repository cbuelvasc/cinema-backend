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

type StateRepository interface {
	GetAllStates(ctx context.Context, page int64, limit int64, countryId string) (*model.PagedState, error)
	GetStateById(ctx context.Context, id string) (*model.State, error)
	SaveState(ctx context.Context, state *model.State) (*model.State, error)
	UpdateState(ctx context.Context, id string, state *model.State) (*model.State, error)
	DeleteState(ctx context.Context, id string) error
}

type stateRepositoryImpl struct {
	Connection *mongo.Database
}

func NewStateRepository(Connection *mongo.Database) StateRepository {
	return &stateRepositoryImpl{Connection: Connection}
}

func (stateRepository *stateRepositoryImpl) GetAllStates(ctx context.Context, page int64, limit int64, countryId string) (*model.PagedState, error) {
	var states []model.State

	var filter = bson.M{}
	if len(countryId) > 0 {
		filter = bson.M{
			"countryId": countryId,
		}
	}

	collection := stateRepository.Connection.Collection("states")
	projection := bson.D{
		{"id", 1},
		{"name", 1},
		{"countryId", 1},
		{"cities", 1},
		{"created_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&states).Find()
	if err != nil {
		return nil, err
	}

	if states == nil {
		return nil, exception.NotFoundRequestException("States not found")
	}

	return &model.PagedState{
		Data:     states,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (stateRepository *stateRepositoryImpl) GetStateById(ctx context.Context, id string) (*model.State, error) {
	var state model.State
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	err := stateRepository.Connection.Collection("states").FindOne(ctx, filter).Decode(&state)
	if err != nil {
		return nil, exception.ResourceNotFoundException("State", "id", id)
	}
	return &state, nil
}

func (stateRepository *stateRepositoryImpl) SaveState(ctx context.Context, state *model.State) (*model.State, error) {
	state.ID = primitive.NewObjectID()

	_, err := stateRepository.Connection.Collection("states").InsertOne(ctx, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func (stateRepository *stateRepositoryImpl) UpdateState(ctx context.Context, id string, state *model.State) (*model.State, error) {
	state.ID = primitive.NewObjectID()

	_, err := stateRepository.Connection.Collection("states").InsertOne(ctx, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func (stateRepository *stateRepositoryImpl) DeleteState(ctx context.Context, id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	result, err := stateRepository.Connection.Collection("states").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return exception.NotFoundRequestException(fmt.Sprintf("State not found with id: %s", id))
	}

	return nil
}
