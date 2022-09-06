package ipfs

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
	"io"
)

type Freezer struct {
	resty     *resty.Client
	authToken string
	url       string
}

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func New(authToken, url string) *Freezer {
	client := resty.New()
	return &Freezer{
		resty:     client,
		authToken: authToken,
		url:       url,
	}
}

type IPFSResponse struct {
	CID string `json:"cid"`
}

type Response struct {
	OK    bool          `json:"ok"`
	Value *IPFSResponse `json:"value"`
}

func (s *Freezer) StoreWithMeta(name, desc string, image io.Reader) (string, string, error) {
	resp, err := s.upload(image)
	if err != nil {
		return "", "", err
	}
	if resp.Value == nil {
		return "", "", errors.New("no ipfs resp")
	}

	metadata := &Metadata{
		Name:        name,
		Description: desc,
		Image:       fmt.Sprintf("ipfs://%s", resp.Value.CID),
	}

	result, err := s.upload(metadata)
	if err != nil {
		return "", "", err
	}

	return metadata.Image, fmt.Sprintf("ipfs://%s", result.Value.CID), nil
}

func (s *Freezer) Store(image io.Reader) (string, error) {
	resp, err := s.upload(image)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("ipfs://%s", resp.Value.CID), nil
}

func (s *Freezer) upload(body interface{}) (*Response, error) {
	resp, err := s.resty.R().
		SetAuthToken(s.authToken).
		SetBody(body).
		Post(s.url)
	if err != nil {
		return nil, fmt.Errorf("resty request: %w", err)
	}

	r := &Response{}

	if err := json.Unmarshal(resp.Body(), r); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return r, nil
}
