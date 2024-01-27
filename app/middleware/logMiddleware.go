package middleware

import (
	"carApp/app/logging"
	"fmt"
	"net/http"
)

func LogMiddleware(next http.Handler, log logging.ILogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.LogConsole().Info("masuk ke middleware")

		next.ServeHTTP(w, r)

		log.LogConsole().Info("keluar middleware")
		fmt.Println("")
	})
}
