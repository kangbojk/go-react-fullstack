package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kangbojk/go-react-fullstack/internal/entity/account"
	"github.com/kangbojk/go-react-fullstack/internal/entity/tenant"
	"github.com/kangbojk/go-react-fullstack/internal/server/data"
	"github.com/kangbojk/go-react-fullstack/internal/usecase"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
	"github.com/kangbojk/go-react-fullstack/pkg/password"

	"github.com/gorilla/mux"
)

func GetCurrentAccount(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "sid")
		if err != nil {
			log.Printf("GetCurrentAccount: Error retrieving session: %v", err)
			return
		}

		toJ := &delivery_data.Account{
			ID:       session.Values["userID"].(string),
			Email:    session.Values["email"].(string),
			Actions:  session.Values["actions"].(string),
			TenantID: session.Values["tenantID"].(string),
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

// CreateAccount (aka register) returns a handler for POST /api/accounts requests
func CreateAccount(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Email    string `json:"Email"`
			Password string `json:"password"`
			Actions  string `json:"actions"`
			TenantID string `json:"tenant_id"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		hash_passwd, err := password.Generate(input.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		new_tenant := &tenant.Tenant{
			Users:    0,
			Capacity: 100,
		}

		new_tenant.ID, err = srv.CreateTenant(new_tenant)
		if err != nil {
			// log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if input.Actions == "" {
			input.Actions = "RW"
			log.Println("Default admin for register users")
		}

		new_account := &account.Account{
			Email:    input.Email,
			Password: hash_passwd,
			Actions:  input.Actions,
			TenantID: new_tenant.ID,
		}

		new_account.ID, err = srv.CreateAccount(new_account)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		toJ := &delivery_data.Account{
			ID:       id.IDToString(new_account.ID),
			Email:    new_account.Email,
			Actions:  new_account.Actions,
			TenantID: id.IDToString(new_account.TenantID),
		}

		setUserSession(w, r, new_account, new_tenant)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func GetAccount(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id_, err := id.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data, err := srv.GetAccount(id_)
		if err != nil && err != account.ErrInvalid {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		toJ := &delivery_data.Account{
			ID:       id.IDToString(data.ID),
			Email:    data.Email,
			Actions:  data.Actions,
			TenantID: id.IDToString(data.TenantID),
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func AssignTenantToAccount(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			AccountID string `json:"accountid"`
			TenantID  string `json:"tenantid"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println("AssignTenantToAccount: ", input.AccountID, input.TenantID)

		aid_, err := id.StringToID(input.AccountID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		account_, err := srv.GetAccount(aid_)
		if err != nil && err != account.ErrInvalid {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if account_ == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		tid_, err := id.StringToID(input.TenantID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		tenant_, err := srv.GetTenant(tid_)
		if err != nil && err != tenant.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if tenant_ == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		err = srv.AssignTenantToAccount(account_, tenant_)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		session, err := store.Get(r, "sid")
		if err != nil {
			log.Printf("GetCurrentAccount: Error retrieving session: %v", err)
			return
		}

		session.Values["tenantID"] = id.IDToString(tenant_.ID)
		session.Values["users"] = tenant_.Users
		session.Values["capacity"] = tenant_.Capacity
		log.Printf("assign tenant to account, users: %d, cap: %d", session.Values["users"], session.Values["capacity"])
		err = session.Save(r, w)
		if err != nil {
			log.Fatalf("Error saving session: %v", err)
		}

		toJ := &delivery_data.Account{
			ID:       session.Values["userID"].(string),
			Email:    session.Values["email"].(string),
			Actions:  session.Values["actions"].(string),
			TenantID: session.Values["tenantID"].(string),
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func GetTenant(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id_, err := id.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data, err := srv.GetTenant(id_)
		if err != nil && err != tenant.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		toJ := &delivery_data.Tenant{
			ID:       id.IDToString(data.ID),
			Users:    data.Users,
			Capacity: data.Capacity,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func CreateTenant(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Users    int `json:"users"`
			Capacity int `json:"capacity"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		new_tenant := &tenant.Tenant{
			Users:    input.Users,
			Capacity: input.Capacity,
		}

		new_tenant.ID, err = srv.CreateTenant(new_tenant)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		toJ := &delivery_data.Tenant{
			ID:       id.IDToString(new_tenant.ID),
			Users:    new_tenant.Users,
			Capacity: new_tenant.Capacity,
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func UpgradePlan(srv usecase.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "sid")
		if err != nil {
			log.Printf("UpgradePlan Error: %v", err)
			return
		}

		tid_string := session.Values["tenantID"].(string)
		tid_, err := id.StringToID(tid_string)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		tenant_, err := srv.GetTenant(tid_)
		if err != nil && err != tenant.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if tenant_ == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		log.Printf("Before UpgradePlan, users: %d(session) %d(storage), cap: %d", session.Values["users"], tenant_.Users, session.Values["capacity"])

		tenant_.Capacity = 1000
		err = srv.UpdateTenant(tenant_)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		session.Values["capacity"] = 1000
		err = session.Save(r, w)
		if err != nil {
			log.Fatalf("Error saving session: %v", err)
			return
		}

		log.Printf("After UpgradePlan, users: %d(session) %d(storage), cap: %d", session.Values["users"], tenant_.Users, session.Values["capacity"])

		toJ := &struct {
			Users    int `json:"users"`
			Capacity int `json:"capacity"`
		}{
			session.Values["users"].(int),
			session.Values["capacity"].(int),
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func AddTenantUsers(srv usecase.Service, msgChan chan<- int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "sid")
		if err != nil {
			log.Printf("Error retrieving session: %v", err)
		}

		account_id := session.Values["userID"].(string)
		tenant_id := session.Values["tenantID"].(string)
		actions := session.Values["actions"].(string)

		log.Println(account_id, tenant_id, actions)

		if actions != "RW" {
			w.Write([]byte(account.ErrNotAuthorize.Error()))
			return
		}

		var input struct {
			Users int `json:"users"`
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		aid_, err := id.StringToID(account_id)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		acc_, err := srv.GetAccount(aid_)
		if err != nil {
			log.Println("GetAccount err")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = srv.AddUser(acc_, input.Users)
		if err == tenant.ErrFull {
			toJ := &struct {
				Full bool `json:"full"`
			}{
				true,
			}
			if err := json.NewEncoder(w).Encode(toJ); err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			return
		}

		if err != nil {
			log.Println("AddUser err")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		tid_, err := id.StringToID(tenant_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		tenant_, err := srv.GetTenant(tid_)
		if err != nil && err != tenant.ErrNotFound {
			log.Println("GetTenant err")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		msgChan <- tenant_.Users

		session.Values["users"] = tenant_.Users
		err = session.Save(r, w)
		if err != nil {
			log.Fatalf("Error saving session: %v", err)
			return
		}

		log.Printf("Add %d users\n", session.Values["users"])
		fmt.Fprintf(w, "Add %d users", tenant_.Users)
	}
}

func GetIndex(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}
