package service

import (
	"project2/app/internal"

	"github.com/jayaraj/project1/pkg"
	"github.com/pkg/errors"
)

type project2Service struct {
	project1 pkg.Project1
}

func NewProject2Service(project1 pkg.Project1) *project2Service {
	return &project2Service{
		project1: project1,
	}
}

func (p *project2Service) GetTopTenUsedWords(request internal.GetTopTenUsedWordsRequest) (response internal.GetTopTenUsedWordsResponse, err error) {
	req := pkg.TopTenUsedWordsRequest{
		Text: request.Text,
	}
	resp, err := p.project1.GetTopTenUsedWords(req)
	if err != nil {
		return internal.GetTopTenUsedWordsResponse{}, errors.Wrapf(err, "get top ten used words from project failed")
	}
	response.Data = make([]internal.Occurance, len(resp.Data))
	for i, o := range resp.Data {
		response.Data[i] = internal.Occurance{Word: o.Word, Count: o.Count}
	}
	return response, nil
}
