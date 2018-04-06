package webutils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//从request中获取request data struct,Content-Type须为application/json
func GetReqJSON(r *http.Request, reqJson interface{}) error {
	if strings.Index(r.Header.Get("Content-Type"), "application/json") == -1 {
		return errors.New("Content-Type should be application/json,Content-Type: " + r.Header.Get("Content-Type"))
	}
	var body []byte
	var err error
	if r.Method == "GET" {
		query, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			return err
		}
		body = []byte(query)
	} else {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(body, &reqJson)
}
