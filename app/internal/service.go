package internal

//go:generate mockgen -source=service.go  -destination=mock/service.go -package=mock
type Project2Service interface {
	GetTopTenUsedWords(request GetTopTenUsedWordsRequest) (response GetTopTenUsedWordsResponse, err error)
}

type GetTopTenUsedWordsRequest struct {
	Text string `json:"text"`
}

type GetTopTenUsedWordsResponse struct {
	Data []Occurance `json:"data"`
}

type Occurance struct {
	Word  string `json:"word"`
	Count uint64 `json:"count"`
}
