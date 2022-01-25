package httpservice

import (
	"net/http"
	"project2/app/internal"
	"project2/app/internal/serviceerror"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TopTenUsedWords struct {
	Text string `json:"text" validate:"required"`
}

func CreateTopTenUsedWordsHandler(project2 internal.Project2Service) gin.HandlerFunc {
	mapGetTopTenUsedWordsRequest := func(request TopTenUsedWords) internal.GetTopTenUsedWordsRequest {
		return internal.GetTopTenUsedWordsRequest{
			Text: request.Text,
		}
	}
	return func(c *gin.Context) {
		var request TopTenUsedWords
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		v := validator.New()
		if err := v.Struct(request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		response, err := project2.GetTopTenUsedWords(mapGetTopTenUsedWordsRequest(request))
		if err != nil {
			serviceerror.AbortOnError(c, err)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
