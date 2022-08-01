package services

import (
	"context"
	"errors"
	"time"

	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewPostService(collection *mongo.Collection, ctx context.Context) PostService {
	return &PostServiceImpl{collection, ctx}
}

func (p *PostServiceImpl) CreatePost(post *models.CreatePostRequest) (*models.DBPost, error) {

	var tempPost models.DBPost
	tempPost.Title = post.Title
	tempPost.Content = post.Content
	tempPost.Image = post.Image
	tempPost.User = post.User
	tempPost.CreatedAt = time.Now()
	tempPost.UpdatedAt = tempPost.CreatedAt
	res, err := p.collection.InsertOne(p.ctx, tempPost)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("post with that title already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := p.collection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for title")
	}

	var newPost *models.DBPost
	query := bson.M{"_id": res.InsertedID}
	if err = p.collection.FindOne(p.ctx, query).Decode(&newPost); err != nil {
		return nil, err
	}

	return newPost, nil
}

func (p *PostServiceImpl) UpdatePost(id string, data *models.UpdatePost) (*models.DBPost, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.collection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *models.DBPost

	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil
}

func (p *PostServiceImpl) FindPostById(id string) (*models.DBPost, error) {
	// convert string id to objectID
	obId, _ := primitive.ObjectIDFromHex(id)

	// create a query command by id
	query := bson.M{"_id": obId}

	// create container
	var post *models.DBPost

	// find one post by query command
	if err := p.collection.FindOne(p.ctx, query).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return post, nil
}

func (p *PostServiceImpl) FindPosts(page int, limit int) ([]*models.DBPost, error) {

	// page return product
	if page == 0 {
		page = 1
	}

	// limit data return
	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	// create a query command
	query := bson.M{}

	// find all posts with optional data
	cursor, err := p.collection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	// create container for data
	var posts []*models.DBPost

	// with data find out, we will decode them and append to array
	for cursor.Next(p.ctx) {
		post := &models.DBPost{}
		err := cursor.Decode(post)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	// if any item error, return err
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// if data is empty, return nil
	if len(posts) == 0 {
		return []*models.DBPost{}, nil
	}

	return posts, nil
}

func (p *PostServiceImpl) DeletePost(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.collection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
