package httpservice_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"project2/app/internal"
	"project2/app/internal/httpservice"
	"project2/app/internal/mock"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	sampleText = `I’m now convinced that the returns can be far greater in NFTs than even cryptocurrency at large, and that these returns are, far from being baseless and rooted in mere greater-fool-theory hype and bubbly fadness, in fact extraordinarily justified for a few select investments.`
)

func TestCreateTopTenUsedWordsHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	projectService := mock.NewMockProject2Service(mockCtrl)
	router := gin.Default()
	router.POST("/", httpservice.CreateTopTenUsedWordsHandler(projectService))

	tests := []struct {
		name     string
		request  string
		status   int
		response string
		setup    func()
	}{
		{
			name:     "get top words successfully",
			request:  fmt.Sprintf(`{"text":"%s"}`, sampleText),
			status:   http.StatusOK,
			response: `{"data":[{"word":"test","count":1}]}`,
			setup: func() {
				request := internal.GetTopTenUsedWordsRequest{
					Text: sampleText,
				}
				projectService.EXPECT().GetTopTenUsedWords(request).Return(internal.GetTopTenUsedWordsResponse{
					Data: []internal.Occurance{{Word: "test", Count: 1}},
				}, nil).Times(1)
			},
		},
		{
			name:     "fail on body unmarshal",
			request:  `{"text":"test",`,
			status:   http.StatusBadRequest,
			response: `{"message":"unexpected EOF"}`,
			setup:    func() {},
		},
		{
			name:     "fail on missing text",
			request:  `{}`,
			status:   http.StatusBadRequest,
			response: `{"message":"Key: 'TopTenUsedWords.Text' Error:Field validation for 'Text' failed on the 'required' tag"}`,
			setup:    func() {},
		},
		{
			name:     "fail on unknown error",
			request:  fmt.Sprintf(`{"text":"%s"}`, sampleText),
			status:   http.StatusInternalServerError,
			response: `{"message":"test"}`,
			setup: func() {
				request := internal.GetTopTenUsedWordsRequest{
					Text: sampleText,
				}
				projectService.EXPECT().GetTopTenUsedWords(request).Return(internal.GetTopTenUsedWordsResponse{}, errors.New("test")).Times(1)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/", strings.NewReader(test.request))
			test.setup()
			router.ServeHTTP(recorder, req)
			response, err := ioutil.ReadAll(recorder.Body)
			assert.NoError(t, err)
			assert.Equal(t, test.status, recorder.Code)
			assert.Equal(t, test.response, string(response))
		})
	}
}
