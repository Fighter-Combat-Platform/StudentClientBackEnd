package repository

import (
	"awesomeProject/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateRoomPlayer(player model.RoomPlayer) {
	_, err = roomPlayerCol.InsertOne(context.TODO(), player)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteRoomPlayer(player model.RoomPlayer) {
	filter := bson.D{{"rid", player.Rid}, {"uid", player.Uid}}
	_, err = roomPlayerCol.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteAllRoomPlayers(rid int) {
	filter := bson.D{{"rid", rid}}
	_, err = roomPlayerCol.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetRoomPlayerByIds(player model.RoomPlayer) (ret model.RoomPlayer) {
	filter := bson.D{{"rid", player.Rid}, {"uid", player.Uid}}
	err = roomPlayerCol.FindOne(context.TODO(), filter).Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetRoomPlayersByRid(rid int) (ret []*model.RoomPlayer) {
	filter := bson.D{{"rid", rid}}
	findOptions := options.Find()
	cur, _ := roomPlayerCol.Find(context.TODO(), filter, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.RoomPlayer
		err = cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, &elem)
	}
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}
	err = cur.Close(context.TODO())
	if err != nil {
		return nil
	}
	return
}

func ModifyRoomPlayer(player model.RoomPlayer) {
	filter := bson.D{{"rid", player.Rid}, {"uid", player.Uid}}
	update := bson.D{{"$set", player}}
	_, err = roomPlayerCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return
}
