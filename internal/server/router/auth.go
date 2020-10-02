package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kangbojk/go-react-fullstack/internal/entity/account"
	"github.com/kangbojk/go-react-fullstack/internal/entity/tenant"
	"github.com/kangbojk/go-react-fullstack/internal/usecase"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
	"github.com/kangbojk/go-react-fullstack/pkg/password"
)

// Middleware function, which will be called for each request
func isLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "sid")
		if err != nil {
			log.Printf("Error retrieving session: %v", err)
		}

		if userID, found := session.Values["userID"]; found && userID != "" {

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func LoginWithEmail(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		account_, err := srv.FindUserWithEmail(input.Email)
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err = password.Compare(account_.Password, input.Password); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(account.ErrInvalid.Error()))
			return
		}

		tenant_, err := srv.GetTenant(account_.TenantID)
		if err != nil && err != tenant.ErrNotFound {
			log.Println("GetTenant err")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		setUserSession(w, r, account_, tenant_)
	}
}

func setUserSession(w http.ResponseWriter, r *http.Request, account_ *account.Account, tenant_ *tenant.Tenant) {
	// use bearer session token instead of JWT for the time being
	session, err := store.Get(r, "sid")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error retrieving session: %v", err)
		return
	}

	if tenant_ != nil {
		session.Values["users"] = tenant_.Users
		session.Values["capacity"] = tenant_.Capacity
	}

	session.Values["userID"] = id.IDToString(account_.ID)
	session.Values["email"] = account_.Email
	session.Values["actions"] = account_.Actions
	session.Values["tenantID"] = id.IDToString(account_.TenantID)

	err = session.Save(r, w)
	if err != nil {
		log.Fatalf("Error saving session: %v", err)
	}
}

func Logout(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "sid")
		if err != nil {
			log.Printf("Error retrieving session: %v", err)
		}

		for k, _ := range session.Values {
			delete(session.Values, k)
		}

		err = session.Save(r, w)
		if err != nil {
			log.Fatalf("Error saving session: %v", err)
		}
	}
}
