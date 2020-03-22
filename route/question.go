package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"wirsindcorona/model"
	"wirsindcorona/reponse"
)

func InitQuestionRoute(route *gin.RouterGroup) {
	route.POST("", createQuestion)
	route.GET("", getQuestions)
	route.GET(":quid", getQuestion)
	route.PUT("", updateQuestion)
	route.DELETE(":quid", deleteQuestion)
}

func createQuestion(gc *gin.Context) {
	question := model.Question{}

	jerr := gc.ShouldBindJSON(&question)
	if jerr != nil {
		gc.JSON(http.StatusUnprocessableEntity, reponse.ReturnError(jerr))
		return
	}

	question.GenerateQuid()

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

func getQuestion(gc *gin.Context) {
	quid := gc.Param("quid")

	res := questionsColl.FindOne(context.TODO(), bson.D{{"quid", quid}})

	var question model.Question
	derr := res.Decode(&question)
	if derr != nil {
		gc.JSON(http.StatusInternalServerError, reponse.ReturnError(derr))
		return
	}

	gc.JSON(http.StatusOK, reponse.ReturnData(question))
}

func updateQuestion(gc *gin.Context) {
	question := model.Question{}
	jerr := gc.ShouldBindJSON(&question)
	if jerr != nil {
		gc.JSON(http.StatusUnprocessableEntity, reponse.ReturnError(jerr))
		return
	}

	find := bson.D{
		{"quid", question.Quid},
	}

	_, ier := questionsColl.ReplaceOne(context.TODO(), find, &question)
	if ier != nil {
		gc.JSON(http.StatusInternalServerError, reponse.ReturnError(ier))
		return
	}

	gc.JSON(http.StatusOK, reponse.ReturnData(question))
}

func deleteQuestion(gc *gin.Context) {
	quid := gc.Param("quid")

	find := bson.D{
		{"quid", quid},
	}

	delRes, err := questionsColl.DeleteOne(context.TODO(), find)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, reponse.ReturnError(err))
		return
	}

	gc.JSON(http.StatusOK, reponse.ReturnData(delRes.DeletedCount))
}
