package oriflame

import "net/http"

type Client struct {
	HTTPClient *http.Client
}

func (client *Client) Init() {
	client.HTTPClient = new(http.Client)
}
