package repository

import (
	"awesomeProject/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetUserByUsername(username string) (ret model.User) {
	filter := bson.D{{"username", username}}
	err = userCol.FindOne(context.TODO(), filter).Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetUserByUid(uid int) (ret model.User) {
	filter := bson.D{{"uid", uid}}
	err = userCol.FindOne(context.TODO(), filter).Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetIDByUsername(username string) (ret int) {
	filter := bson.D{{"username", username}}
	var user model.User
	err = userCol.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	ret = user.Uid
	return
}

func CreateUser(user model.User) {
	_, err = userCol.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ModifyUser(user model.User) {
	filter := bson.D{{"uid", user.Uid}}
	update := bson.D{{"$set", user}}
	_, err = userCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return
}
