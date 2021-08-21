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

type UserRepository interface {
	GetAllUser(ctx context.Context, page int64, limit int64) (*model.PagedUser, error)
	SaveUser(ctx context.Context, user *model.User) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userRepositoryImpl struct {
	Connection *mongo.Database
}

func NewUserRepository(Connection *mongo.Database) UserRepository {
	return &userRepositoryImpl{Connection: Connection}
}

func (userRepository *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var existingUser model.User
	filter := bson.M{"email": email}
	err := userRepository.Connection.Collection("users").FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		return nil, err
	}
	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) GetAllUser(ctx context.Context, page int64, limit int64) (*model.PagedUser, error) {
	var users []model.User

	filter := bson.M{}

	collection := userRepository.Connection.Collection("users")

	projection := bson.D{
		{"id", 1},
		{"name", 1},
		{"lastname", 1},
		{"birthDate", 1},
		{"email", 1},
		{"avatar", 1},
		{"banner", 1},
		{"biography", 1},
		{"location", 1},
		{"webSite", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedUser{
		Data:     users,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (userRepository *userRepositoryImpl) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	err := userRepository.Connection.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.Password = ""
	return &user, nil
}

func (userRepository *userRepositoryImpl) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.ID = primitive.NewObjectID()

	_, err := userRepository.Connection.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	registry := make(map[string]interface{})
	if len(user.Name) > 0 {
		registry["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		registry["lastname"] = user.Lastname
	}
	registry["birthDate"] = user.BirthDate
	if len(user.Avatar) > 0 {
		registry["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registry["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		registry["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		registry["location"] = user.Location
	}
	if len(user.WebSite) > 0 {
		registry["webSite"] = user.WebSite
	}

	filter := bson.M{"_id": bson.M{"$eq": objectId}}

	updateString := bson.M{
		"$set": registry,
	}

	result, err := userRepository.Connection.Collection("users").UpdateOne(ctx, filter, updateString)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	user.Password = ""
	return user, nil
}

func (userRepository *userRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	result, err := userRepository.Connection.Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return exception.ResourceNotFoundException("User", "id", id)
	}

	return nil
}
