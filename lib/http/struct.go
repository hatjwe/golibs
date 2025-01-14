package http

const (
	Post = "Post"
	Get  = "Get"
)

type HttpRespone struct {
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
	Error      error  `json:"error"`
}
