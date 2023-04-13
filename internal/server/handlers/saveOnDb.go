package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/johnsoncwb/mongo_gin/internal/db"
	"github.com/johnsoncwb/mongo_gin/internal/models"
	"github.com/johnsoncwb/mongo_gin/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func SaveOnDatabase(c *gin.Context) {

	user := models.UserObject{}
	ctx := context.TODO()

	ginContext := c.Copy()
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.HandleError(ginContext, http.StatusInternalServerError, err)
		return
	}

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

	result := models.UserObject{}

	filter := bson.D{{
		Key: "email",
		Value: bson.D{{
			Key:   "$eq",
			Value: user.Email,
		}},
	}}

	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			addedUser, err := coll.InsertOne(ctx, user)
			if err != nil {
				utils.HandleError(ginContext, http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "User successfully added!",
				"user":    addedUser,
			})
			return
		}
		utils.HandleError(ginContext, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "This e-mail is already on database",
	})

}
