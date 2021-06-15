package helpers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	ContentTypeApplicationJson           = "application/json"
	ContentTypeApplicationFormUrlEncoded = "application/x-www-form-urlencoded"
)

type HttpConfig struct {
	HttpClient *http.Client
	Timeout    *time.Duration
	Host       string
	Port       string
}

func (c HttpConfig) BuildBaseUrl() string {
	if len(c.Port) > 0 {
		return fmt.Sprintf("%s:%s", c.Host, c.Port)
	}

	return c.Host
}

type HttpRequestPayload struct {
	Method      string
	Url         string
	QueryParams map[string]string
	Body        interface{}
	Result      interface{}
	Logger      logs.Log
	Client      *http.Client
	TimeoutReq  *time.Duration // default 10s
	Header      *http.Header
}

func HttpRequest(payload HttpRequestPayload) error {
	var bodyByte []byte

	if payload.Body != nil {
		reqBody, err := json.Marshal(payload.Body)

		if err != nil {
			return err
		}

		bodyByte = reqBody
	}

	req, err := http.NewRequest(payload.Method, payload.Url, bytes.NewBuffer(bodyByte))

	if err != nil {
		return err
	}

	if payload.QueryParams != nil {
		q := req.URL.Query()

		for key, val := range payload.QueryParams {
			q.Set(key, val)
		}

		req.URL.RawQuery = q.Encode()
	}

	if payload.Header == nil {
		req.Header.Set("Content-Type", ContentTypeApplicationJson)
	} else {
		req.Header = *payload.Header
	}

	// --- do request

	timeout := 10 * time.Second

	if payload.TimeoutReq != nil {
		timeout = *payload.TimeoutReq
	}

	newClient := http.Client{
		Timeout: timeout,
	}

	if payload.Client != nil {
		newClient = *payload.Client
	}

	resp, err := newClient.Do(req)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return errors.BadRequest("request timeout.")
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		readResp, _ := ioutil.ReadAll(resp.Body)
		payload.Logger.Error(string(readResp))
		payload.Logger.Error(fmt.Sprintf("%v", resp.StatusCode))

		return errors.BadRequest("request error.")
	}

	readResp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(readResp, &payload.Result); err != nil {
		return err
	}

	return nil
}

func GenerateBasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}