package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
)

type Rest struct {
	request  *http.Request
	Response *http.Response
}

const (
	NewRelicTxnKey = "NewRelicTxnKey"
)

//RestClient interface for rest client
type RestClient interface {
	Do(c *gin.Context, o interface{}) error
}

func NewRequest(method, url, token string, body interface{}) (RestClient, error) {
	var buf io.ReadWriter

	if body != nil {
		b, err := json.Marshal(body)
		buf = bytes.NewBuffer(b)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if len(token) > 0 {
		req.Header.Add("Authorization", token)
	}

	return &Rest{request: req}, nil
}

func (r *Rest) Do(c *gin.Context, o interface{}) error {
	if nr, ok := c.Get(NewRelicTxnKey); ok {
		txn := nr.(newrelic.Transaction)
		segment := newrelic.StartExternalSegment(txn, r.request)
		defer segment.End()
	}

	client := &http.Client{}
	resp, err := client.Do(r.request)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, o); err == nil {
		r.Response = resp
	}

	return err
}
