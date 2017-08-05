package utils

import (
	"fmt"
	"net/http"
)

// Adapter this.
type Adapter func(http.Handler) http.Handler

// Adapt this.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	if h == nil {
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//TODO: this needs to send back a dynamic/default message. Somehow we need to append to the w.
			fmt.Println("end == nil")
			return
		})
	}
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
