package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/johnsoncwb/mongo_gin/internal/db"
	"github.com/johnsoncwb/mongo_gin/internal/models"
	"github.com/johnsoncwb/mongo_gin/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetById(c *gin.Context) {
	id := c.Param("id")
	ginContext := c.Copy()
	ctx := context.TODO()

	coll, client, err := db.ConnectMongo()
	if err != nil {
		utils.HandleError(ginContext, http.StatusInternalServerError, err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.HandleError(ginContext, http.StatusInternalServerError, err)
	}

	filter := bson.D{{
		Key: "_id",
		Value: bson.D{{
			Key:   "$eq",
			Value: objId,
		}},
	}}

	result := models.UserObject{}

	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.HandleError(ginContext, http.StatusNotFound, err)
			return
		}
		utils.HandleError(ginContext, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
