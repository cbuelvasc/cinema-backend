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

type RoomRepository interface {
	GetAllRooms(ctx context.Context, page int64, limit int64) (*model.PagedRoom, error)
	GetRoomById(ctx context.Context, id string) (*model.Room, error)
	GetRoomByCinema(ctx context.Context, cityId string) (*model.Room, error)
	SaveRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	UpdateRoom(ctx context.Context, id string, roomId *model.Room) (*model.Room, error)
	DeleteRoom(ctx context.Context, id string, roomId string) error
}

type roomRepositoryImpl struct {
	Connection *mongo.Database
}

func NewRoomRepository(Connection *mongo.Database) RoomRepository {
	return &roomRepositoryImpl{Connection: Connection}
}

func (roomRepository *roomRepositoryImpl) GetAllRooms(ctx context.Context, page int64, limit int64) (*model.PagedRoom, error) {
	var rooms []model.Room

	filter := bson.M{
	}

	collection := roomRepository.Connection.Collection("rooms")

	projection := bson.D{
		{"id", 1},
		{"name", 1},
		{"capacity", 1},
		{"format", 1},
		{"schedules", 1},
		{"created_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&rooms).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedRoom{
		Data:     rooms,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (roomRepository *roomRepositoryImpl) GetRoomById(ctx context.Context, id string) (*model.Room, error) {
	var room model.Room
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	err := roomRepository.Connection.Collection("rooms").FindOne(ctx, filter).Decode(&room)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Room", "id", id)
	}
	return &room, nil
}

func (roomRepository *roomRepositoryImpl) GetRoomByCinema(ctx context.Context, cinemaId string) (*model.Room, error) {
	panic("implement me")
}

func (roomRepository *roomRepositoryImpl) SaveRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	room.ID = primitive.NewObjectID()

	_, err := roomRepository.Connection.Collection("rooms").InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (roomRepository *roomRepositoryImpl) UpdateRoom(ctx context.Context, id string, room *model.Room) (*model.Room, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	registry := make(map[string]interface{})
	if len(room.Name) > 0 {
		registry["name"] = room.Name
	}
	if len(room.Capacity) > 0 {
		registry["capacity"] = room.Capacity
	}
	if len(room.Format) > 0 {
		registry["format"] = room.Format
	}
	if len(room.CinemaId) > 0 {
		registry["cinemaId"] = room.CinemaId
	}
	if len(room.Schedules) > 0 {
		registry["schedules"] = room.Schedules
	}

	filter := bson.M{"_id": bson.M{"$eq": objectId}}

	updateString := bson.M{
		"$set": registry,
	}

	result, err := roomRepository.Connection.Collection("rooms").UpdateOne(ctx, filter, updateString)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("Room", "id", id)
	}

	room.ID = objectId
	return room, nil
}

func (roomRepository *roomRepositoryImpl) DeleteRoom(ctx context.Context, id string, roomId string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":    objectId,
		"roomId": roomId,
	}

	result, err := roomRepository.Connection.Collection("rooms").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return exception.NotFoundRequestException(fmt.Sprintf("Room not found with id: %s and roomId: %s", id, roomId))
	}

	return nil
}