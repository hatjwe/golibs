package http

const (
	Post = "POST"
	Get  = "GET"
)

type HttpRespone struct {
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
	Error      error  `json:"error"`
}
