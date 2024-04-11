package middlewares

import "net/http"

func MethodNotAllowedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !isMethodAllowed(r.Method, r.URL.Path) {
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
					return
			}

			next.ServeHTTP(w, r)
	})
}

func isMethodAllowed(method string, path string) bool {

	allowedMethods := map[string][]string{
			"/insert": {"POST"},
			"/getAll": {"GET"},
	}

	methods, ok := allowedMethods[path]
	if !ok {
			return false 
	}
	for _, m := range methods {
			if m == method {
					return true 
			}
	}
	return false 
}