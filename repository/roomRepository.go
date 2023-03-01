package repository

import (
	"awesomeProject/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ModifyRoomConfig(room model.Room) {
	filter := bson.D{{"rid", room.Rid}}
	update := bson.D{{"$set", bson.D{
		{"capacity", room.Capacity},
		{"roomname", room.RoomName},
		{"mapid", room.MapID},
		{"weathertype", room.WeatherType},
		{"lighttype", room.LightType},
		{"ifpublic", room.IfPublic},
		{"password", room.Password},
		{"type", room.Type},
		{"state", 1},
	}}}
	_, err = roomCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ModifyRoom(room model.Room) {
	filter := bson.D{{"rid", room.Rid}}
	update := bson.D{{"$set", room}}
	_, err = roomCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func AddRoomPlayerCount(roomName string) {
	filter := bson.D{{"roomname", roomName}}
	update := bson.D{
		{"$inc", bson.D{
			{"playercount", 1},
		}},
	}
	_, err = roomCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func SubRoomPlayerCount(roomName string) {
	filter := bson.D{{"roomname", roomName}}
	update := bson.D{
		{"$inc", bson.D{
			{"playercount", -1},
		}},
	}
	_, err = roomCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateRoom(room model.Room) {
	_, err = roomCol.InsertOne(context.TODO(), room)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetRoomByRoomName(roomName string) (ret model.Room) {
	filter := bson.D{{"roomname", roomName}}
	err = roomCol.FindOne(context.TODO(), filter).Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetRidByRoomName(roomName string) (ret int) {
	filter := bson.D{{"roomname", roomName}}
	var result model.Room
	err = roomCol.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result.Rid
}

func GetRoomByRid(rid int) (ret model.Room) {
	filter := bson.D{{"rid", rid}}
	err = roomCol.FindOne(context.TODO(), filter).Decode(&ret)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetRoomList() (roomList []*model.Room) {
	findOptions := options.Find()
	cur, _ := roomCol.Find(context.TODO(), bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		var elem model.Room
		err = cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		roomList = append(roomList, &elem)
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
