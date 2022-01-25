package service_test

import (
	"project2/app/internal"
	"project2/app/internal/service"
	"testing"

	"github.com/jayaraj/project1/pkg"
	"github.com/jayaraj/project1/pkg/mock"
	"github.com/pkg/errors"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetTopTenUsedWords(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	project1 := mock.NewMockProject1(mockCtrl)
	handler := service.NewProject2Service(project1)

	tests := []struct {
		name   string
		Text   string
		max    uint64
		length int
		setup  func()
		err    error
	}{
		{
			name:   "get most used words successfully",
			Text:   `some some test`,
			max:    2,
			length: 2,
			setup: func() {
				project1.EXPECT().GetTopTenUsedWords(pkg.TopTenUsedWordsRequest{
					Text: `some some test`,
				}).Return(pkg.TopTenUsedWordsResponse{
					Data: []pkg.Occurance{
						{Word: "some", Count: 2},
						{Word: "test", Count: 1},
					},
				}, nil).Times(1)
			},
			err: nil,
		},
		{
			name:   "handle error response",
			Text:   `some some test`,
			max:    0,
			length: 0,
			setup: func() {
				project1.EXPECT().GetTopTenUsedWords(pkg.TopTenUsedWordsRequest{
					Text: `some some test`,
				}).Return(pkg.TopTenUsedWordsResponse{}, errors.New("test")).Times(1)
			},
			err: errors.Wrapf(errors.New("test"), "get top ten used words from project failed"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			response, err := handler.GetTopTenUsedWords(internal.GetTopTenUsedWordsRequest{
				Text: test.Text,
			})
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			}
			assert.Len(t, response.Data, test.length)
			if test.length > 0 {
				assert.Equal(t, uint64(test.max), response.Data[0].Count)
			}
		})
	}
}
