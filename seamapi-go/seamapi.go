package seamapi

import (
	"net/http"
	"time"
)

const (
	BaseURL = "https://connect.getseam.com"
)

type Client struct {
	BaseURL		string
	apiKey		string
	HTTPClient 	*http.Client
}

func NewClient(apiKey: string) *Client {
	return &Client{
		BaseURL: BaseURL,
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	},
}

type errorResponse struct {
	Code	int	`json:"code"`
	Message	string	`json:"message"`
}

type successResponse struct {
	Code	int		`json:"code"`
	Data	interface{}	`json:"data"`
}
