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

type TweetRepository interface {
	GetAllTweets(ctx context.Context, page int64, limit int64, userId string) (*model.PagedTweet, error)
	GetTweet(ctx context.Context, id string) (*model.Tweet, error)
	SaveTweet(ctx context.Context, tweet *model.Tweet) (*model.Tweet, error)
	DeleteTweet(ctx context.Context, id string, userId string) error
}

type tweetRepositoryImpl struct {
	Connection *mongo.Database
}

func NewTeewtRepository(Connection *mongo.Database) TweetRepository {
	return &tweetRepositoryImpl{Connection: Connection}
}

func (tweetRepository *tweetRepositoryImpl) GetAllTweets(ctx context.Context, page int64, limit int64, userId string) (*model.PagedTweet, error) {
	var tweets []model.Tweet

	filter := bson.M{
		"userId": userId,
	}

	collection := tweetRepository.Connection.Collection("tweets")

	projection := bson.D{
		{"id", 1},
		{"userId", 1},
		{"message", 1},
		{"created_at", 1},
	}

	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&tweets).Find()
	if err != nil {
		return nil, err
	}

	if tweets == nil {
		return nil, exception.ResourceNotFoundException("Tweets", "userId", userId)
	}

	return &model.PagedTweet{
		Data:     tweets,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (tweetRepository *tweetRepositoryImpl) GetTweet(ctx context.Context, id string) (*model.Tweet, error) {
	var tweet model.Tweet
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}

	err := tweetRepository.Connection.Collection("tweets").FindOne(ctx, filter).Decode(&tweet)
	if err != nil {
		return nil, exception.ResourceNotFoundException("Tweet", "id", id)
	}
	return &tweet, nil
}

func (tweetRepository *tweetRepositoryImpl) SaveTweet(ctx context.Context, tweet *model.Tweet) (*model.Tweet, error) {
	tweet.ID = primitive.NewObjectID()

	_, err := tweetRepository.Connection.Collection("tweets").InsertOne(ctx, tweet)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (tweetRepository *tweetRepositoryImpl) DeleteTweet(ctx context.Context, id string, userId string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":    objectId,
		"userId": userId,
	}

	result, err := tweetRepository.Connection.Collection("tweets").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return exception.NotFoundRequestException(fmt.Sprintf("Tweet not found with id: %s and userId: %s", id, userId))
	}

	return nil
}
