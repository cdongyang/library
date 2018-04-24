package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type BaseHandler struct {
	Filters []Filter
	Handler func(w http.ResponseWriter, r *http.Request)
}

type Filter interface {
	BeforeServeHTTP(w http.ResponseWriter, r *http.Request) bool
	AfterServeHTTP(w http.ResponseWriter, r *http.Request) bool
}

func NewHandler(handleFunc func(w http.ResponseWriter, r *http.Request), filters ...Filter) BaseHandler {
	return BaseHandler{Filters: filters, Handler: handleFunc}
}

func CopyRequestBody(r *http.Request) ([]byte, error) {
	bs, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bs))
	return bs, nil
}

type ResponseWriter struct {
	http.ResponseWriter
	Buffer bytes.Buffer
}

func (w *ResponseWriter) Write(bs []byte) (int, error) {
	w.Buffer.Write(bs)
	return w.ResponseWriter.Write(bs)
}

func (b BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = &ResponseWriter{ResponseWriter: w}
	if b.Filters != nil {
		for _, filter := range b.Filters {
			if filter.BeforeServeHTTP(w, r) {
				return
			}
		}
	}

	b.Handler(w, r)

	if b.Filters != nil {
		for _, filter := range b.Filters {
			if filter.AfterServeHTTP(w, r) {
				return
			}
		}
	}
}

type MethodFilter struct {
	methods []string
}

func NewMethodFilter(methods []string) MethodFilter {
	return MethodFilter{methods: methods}
}

func (f MethodFilter) BeforeServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	for _, method := range f.methods {
		if r.Method == method {
			return false
		}
	}
	return true
}

func (f MethodFilter) AfterServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	return false
}

func WriteJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	bytes, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(`{
			"success": false,
			"timestamp": ` + fmt.Sprint(time.Now().Unix()) + `,
			"message": "marshal json error: ` + err.Error() + `"
		}`))
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "write response error:", err.Error())
	}
}

type Result struct {
	Success   bool        `json:"success"`
	Res       interface{} `json:"result,omitempty"`
	Msg       string      `json:"message,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
}

func Success(w http.ResponseWriter, result interface{}) {
	WriteJSON(w, Result{
		Success:   true,
		Res:       result,
		Timestamp: time.Now().Unix(),
	})
}

func Fail(w http.ResponseWriter, message string) {
	WriteJSON(w, Result{
		Success:   false,
		Msg:       message,
		Timestamp: time.Now().Unix(),
	})
}
