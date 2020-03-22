package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"wirsindcorona/database"
	"wirsindcorona/model"
	"wirsindcorona/reponse"
)

var questionsColl *mongo.Collection

func InitQuestionRoute(route *gin.RouterGroup) {
	questionsColl = database.GetDatabase().Collection("questionsColl")

	route.POST("", createQuestion)
	route.GET("", getQuestions)
}

func createQuestion(gc *gin.Context) {
	question := model.Question{}
	jerr := gc.ShouldBindJSON(&question)
	if jerr != nil {
		gc.JSON(http.StatusUnprocessableEntity, reponse.ReturnError(jerr))
		return
	}

	_, ier := questionsColl.InsertOne(context.TODO(), &question)
	if ier != nil {
		gc.JSON(http.StatusInternalServerError, reponse.ReturnError(ier))
		return
	}

	gc.JSON(http.StatusCreated, reponse.ReturnData(question))
}

func getQuestions(gc *gin.Context) {
	cur, ier := questionsColl.Find(context.TODO(), bson.D{{}}, nil)
	if ier != nil {
		gc.JSON(http.StatusInternalServerError, reponse.ReturnError(ier))
		return
	}

	questions := make([]*model.Question, 0, 0)
	for cur.Next(context.TODO()) {
		var question model.Question
		err := cur.Decode(&question)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, reponse.ReturnError(err))
			return
		}

		questions = append(questions, &question)
	}

	gc.JSON(http.StatusOK, reponse.ReturnData(questions))
}
