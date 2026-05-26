package middlewares

import "net/http"

type AccessRole string

const (
	AccessRoleAdmin  AccessRole = "admin"
	AccessRoleBarber AccessRole = "barber"
)

func Access(role AccessRole, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.Context().Value("role")
		if role == value {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		userRole, ok := value.(AccessRole)
		if !ok {
			http.Error(w, "invalid role", http.StatusUnauthorized)
			return
		}
		if userRole != AccessRoleAdmin {
			if userRole != role {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}
		}
		next(w, r)
	}
}
