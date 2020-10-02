package delivery_data

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Actions  string `json:"actions"`
	TenantID string `json:"tenant_id"`
}
