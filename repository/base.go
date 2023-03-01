package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	MongoClient                                *mongo.Client
	db                                         *mongo.Database
	err                                        error
	roomCol, userCol, roomPlayerCol, configCol *mongo.Collection
)

func InitDBConn() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	uri := "mongodb://10.119.9.37:27017"

	//uri := "mongodb://127.0.0.1:27017"

	opt := new(options.ClientOptions)

	opt = opt.SetMaxPoolSize(10)

	du, _ := time.ParseDuration("5000")

	opt = opt.SetConnectTimeout(du)

	mt, _ := time.ParseDuration("5000")

	opt = opt.SetMaxConnIdleTime(mt)

	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri), opt)
	if err != nil {
		log.Fatal(err)
	}

	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db = MongoClient.Database("flight")

	roomCol = db.Collection("Room")
	userCol = db.Collection("User")
	roomPlayerCol = db.Collection("RoomPlayer")
	configCol = db.Collection("Config")

	return

}
