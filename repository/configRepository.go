package repository

import (
	"awesomeProject/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetTicket() (ret int) {
	filter := bson.D{{"title", "Count"}}
	var result model.Config
	err = configCol.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	ret = result.Ticket
	return
}

func UpdateTicket() {
	filter := bson.D{{"title", "Count"}}
	update := bson.D{
		{"$inc", bson.D{
			{"ticket", 1},
		}},
	}
	_, err = configCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
