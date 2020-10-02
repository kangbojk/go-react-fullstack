package delivery_data

type Tenant struct {
	ID       string `json:"id"`
	Users    int    `json:"users"`
	Capacity int    `json:"capacity"`
}
