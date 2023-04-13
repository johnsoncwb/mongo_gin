package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/johnsoncwb/mongo_gin/internal/db"
	"github.com/johnsoncwb/mongo_gin/internal/models"
	"github.com/johnsoncwb/mongo_gin/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func GetAllByCondition(c *gin.Context) {
	condition := c.Param("maried")
	ginContext := c.Copy()
	ctx := context.TODO()

	coll, client, err := db.ConnectMongo()
	if err != nil {
		utils.HandleError(ginContext, http.StatusInternalServerError, err)
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	var finalCondition bool
	if condition == "false" {
		finalCondition = false
	} else {
		finalCondition = true
	}

	filter := bson.D{{
		Key: "maried",
		Value: bson.D{{
			Key:   "$eq",
			Value: finalCondition,
		}},
	}}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.HandleError(ginContext, http.StatusNotFound, err)
			return
		}
		utils.HandleError(ginContext, http.StatusInternalServerError, err)
		return
	}

	var results []models.UserObject
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, value := range results {
		if value.Age > 5 {
			fmt.Println(value)
		}
	}

	c.JSON(http.StatusOK, results)

}
