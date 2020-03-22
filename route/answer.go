package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
	"wirsindcorona/model"
	"wirsindcorona/reponse"
)

func InitAnswerRoute(route *gin.RouterGroup) {
	route.PUT(":auid", answer)
}

func answer(gc *gin.Context) {
	// TODO Validate Answer

	answer := model.Answer{}
	answer.Auid = gc.Param("auid")
	answer.Time = time.Now()

	_, ier := answerColl.InsertOne(context.Background(), &answer)
	if ier != nil {
		gc.JSON(http.StatusInternalServerError, reponse.ReturnError(ier))
		return
	}

	count, cerr := getAnswerCount(answer.Auid)
	if cerr != nil {
		gc.JSON(http.StatusInternalServerError, reponse.ReturnError(cerr))
		return
	}

	gc.JSON(http.StatusCreated, reponse.ReturnData(count))
}

func getAnswerCount(auid string) (int64, error) {
	return answerColl.CountDocuments(context.Background(), bson.D{{"auid", auid}})
}
