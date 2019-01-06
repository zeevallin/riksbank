package swea

import (
	"context"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/davecgh/go-spew/spew"
)

const (
	scheme      = "http"
	host        = "swea.riksbank.se"
	path        = "/sweaWS/services/SweaWebServiceHttpSoap12Endpoint"
	contentType = "text/xml"
)

// Config represents the configuration for the Riksbank API client
type Config struct {
	HTTPClient *http.Client
}

// New constructs and returns a new Swea client
func New(config Config) *Swea {
	// Setup the HTTP client
	client := http.DefaultClient
	if config.HTTPClient != nil {
		client = config.HTTPClient
	}
	return &Swea{
		client: client,
		url: &url.URL{
			Scheme: scheme,
			Host:   host,
			Path:   path,
		},
	}
}

// Swea is a collection of the methods for the Riksbank API
type Swea struct {
	client *http.Client
	url    *url.URL
}

func (s *Swea) call(ctx context.Context, body io.Reader, v interface{}) error {
	// Build the request
	req, err := http.NewRequest(http.MethodPost, s.url.String(), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", contentType)
	req.WithContext(ctx)

	// Perform the request
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	utf8 := NewValidUTF8Reader(res.Body)

	// Read the response
	bts, err := ioutil.ReadAll(utf8)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bts, v)
	return err
}
