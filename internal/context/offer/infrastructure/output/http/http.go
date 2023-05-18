package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain/ports"
	"net/http"
	"time"
)

var url = "http://localhost:5000/v1/api/offers"

var _ ports.HttpRequester = Requester{}

type Requester struct{}

// Request method that sends requests to different endpoints
func (h Requester) Request(offers domain.Offers) error {
	data, err := json.Marshal(offers)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if http.StatusOK != res.StatusCode {
		return fmt.Errorf("status code was expected %d, but it was obtained %d", http.StatusOK, res.StatusCode)
	}
	return nil
}
