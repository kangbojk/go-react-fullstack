package router

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"

	"github.com/quasoft/memstore"

	"github.com/kangbojk/go-react-fullstack/pkg/usecase"
)

var store = memstore.NewMemStore(
	[]byte("authkey123"),
	[]byte("enckey12341234567890123456789012"),
)

func NewRouter(srv usecase.Service) http.Handler {

	m := mux.NewRouter()

	// messageChan used to pass tenant users
	messageChan := make(chan int, 1)

	m.HandleFunc("/", GetIndex).Methods("GET", "OPTIONS")
	m.HandleFunc("/api/login", LoginWithEmail(srv)).Methods("POST")

	m.HandleFunc("/api/accounts", CreateAccount(srv)).Methods("POST")
	m.HandleFunc("/api/tenants", CreateTenant(srv)).Methods("POST")

	a := m.PathPrefix("/auth").Subrouter()
	a.Use(isLogin)

	a.HandleFunc("/api/account", GetCurrentAccount(srv)).Methods("GET")
	a.HandleFunc("/api/accounts/{id}", GetAccount(srv)).Methods("GET")
	a.HandleFunc("/api/accounts/tenant", AssignTenantToAccount(srv)).Methods("POST")

	a.HandleFunc("/api/tenants/{id}", GetTenant(srv)).Methods("GET")
	a.HandleFunc("/api/tenants/plan", UpgradePlan(srv)).Methods("POST")
	a.HandleFunc("/api/tenants/users", AddTenantUsers(srv, messageChan)).Methods("POST", "PUT")

	a.HandleFunc("/api/logout", Logout(srv)).Methods("POST")
	a.HandleFunc("/ws/tenantUsers", wsEndpoint(messageChan))

	// if(config.production)
	// 	origins :=
	// else
	// 	origins := "*"

	c := cors.New(cors.Options{
		
		AllowedOrigins:   []string{"http://localhost:*", "https://kangbojk.github.io"},
		// AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Set-Cookie"},
	})

	handler := c.Handler(m)
	return handler
}
