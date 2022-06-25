package oriflame

import "net/http"

type Client struct {
	HTTPClient *http.Client
}
