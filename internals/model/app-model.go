package model

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/adiga27/web-hosting-go/internals/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AppName     string             `json:"appName" bson:"appName"`
	AppId       string             `json:"appId" bson:"appId"`
	BranchName  string             `json:"branchName" bson:"branchName"`
	DisplayName string             `json:"displayName" bson:"displayName"`
	Url         string             `json:"url" bson:"url"`
	JobId       string             `json:"jobId" bson:"jobId"`
	Status      string             `json:"status" bson:"status"`

	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

var (
	db            *mongo.Database
	appCollection *mongo.Collection
)

func init() {
	config.LoadEnv()
	client, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("web-hosting")
	appCollection = db.Collection("apps")
}

func GetAppModel(filter interface{}) (*App, error) {
	var app *App

	err := appCollection.FindOne(context.Background(), filter).Decode(&app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no documents found")
		}
		return nil, err
	}

	return app, err
}

func GetAllAppModel() ([]*App, error) {
	var apps []*App
	cursor, err := appCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &apps); err != nil {
		return nil, err
	}

	return apps, nil
}

func (app *App) CreateAppModel() error {
	app.UpdatedAt = time.Now()
	app.CreatedAt = time.Now()
	res, err := appCollection.InsertOne(context.Background(), app)

	if err != nil {
		return err
	}

	app.Id = res.InsertedID.(primitive.ObjectID)
	return nil
}

func UpdateAppModel(filter interface{}, update interface{}) (*App, error) {
	var app *App
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := appCollection.FindOneAndUpdate(context.Background(), filter, update, &opt).Decode(&app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document found")
		}
		return nil, err
	}
	return app, nil
}

func DeleteAppModel(filter interface{}) error {
	var app *App
	err := appCollection.FindOneAndDelete(context.Background(), filter).Decode(&app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("no document found")
		}
		return err
	}
	return nil
}
