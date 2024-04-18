package middlewares

import (
	"net/http"
	"regexp"
)

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
	allowedRoutes := map[string][]string{
		"/insert":  {"POST"},
		"/getAll":  {"GET"},
		"/update/:id": {"PUT", "PATCH"},
	}

	for allowedPath, methods := range allowedRoutes {
		// Verifica se o path da requisição corresponde ao padrão da rota permitida
		if matchesPath(allowedPath, path) {
			for _, m := range methods {
				if m == method {
					return true
				}
			}
			return false // O método não é permitido para esta rota
		}
	}

	return false // Rota não encontrada ou método não permitido
}

// Função auxiliar para verificar se o path da requisição corresponde ao padrão da rota permitida
func matchesPath(pattern, path string) bool {
	// Substitui o parâmetro dinâmico ":id" por uma expressão regular que corresponde a qualquer sequência de caracteres
	pattern = regexp.MustCompile(`:[^/]+`).ReplaceAllString(pattern, `[^/]+`)
	pattern = "^" + pattern + "$"

	matched, err := regexp.MatchString(pattern, path)
	if err != nil {
		return false
	}
	return matched
}
