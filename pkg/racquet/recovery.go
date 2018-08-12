package racquet

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func DisplayErr(h http.Handler, renderStackTrace bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				log.Println(err)
				log.Println(string(stack))

				if renderStackTrace {
					fmt.Fprintf(w, "<h1>Error: %s</h1><pre>%s</pre>", err, string(stack))
				} else {
					http.Error(w, "Something went wrong", http.StatusInternalServerError)
				}
			}
		}()
		h.ServeHTTP(w, r)
	}
}
