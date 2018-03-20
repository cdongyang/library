package http_test

import (
	"unsafe"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func ContextTimeout() {

}

func TestTimeout(t *testing.T) {
	http.HandleFunc("/30ms", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 30)
	})
	http.HandleFunc("/100ms", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Microsecond * 100)
	})
	fmt.Println(unsafe.Sizeof(struct{}{}))
}
