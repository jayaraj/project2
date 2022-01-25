package docs

import (
	"project2/app/internal"
	"project2/app/internal/httpservice"
)

// swagger:route POST /toptenwords  project2 topTenWordsRequest
// Get top ten used words from the text.
// responses:
//   200: topTenWordsResponse
//   400: serviceError
//   500: serviceError

// swagger:response topTenWordsResponse
type TopTenWordsResponse struct {
	// in:body
	Body internal.GetTopTenUsedWordsResponse
}

// swagger:response serviceError
type ServiceErrorResponse struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:parameters topTenWordsRequest
type TopTenWordsRequest struct {
	// in:body
	Body httpservice.TopTenUsedWords
}
