package in_memory

import (
	"gravitational_full_stack_challenge/internal/entity/tenant"
	"gravitational_full_stack_challenge/pkg/ID"
)

//tenantRepoMem in memory repo
type tenantRepoMem struct {
	m map[id.ID]*tenant.Tenant
}

//NewTenantRepoMem create new repository
func NewTenantRepoMem() *tenantRepoMem {
	var m = map[id.ID]*tenant.Tenant{}
	return &tenantRepoMem{
		m: m,
	}
}

//Create an tenant
func (r *tenantRepoMem) Create(e *tenant.Tenant) (id.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an tenant
func (r *tenantRepoMem) Get(id id.ID) (*tenant.Tenant, error) {
	if r.m[id] == nil {
		return nil, tenant.ErrNotFound
	}
	return r.m[id], nil
}

//Update an tenant
func (r *tenantRepoMem) Update(e *tenant.Tenant) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//List tenants
func (r *tenantRepoMem) List() ([]*tenant.Tenant, error) {
	var d []*tenant.Tenant
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an tenant
func (r *tenantRepoMem) Delete(id id.ID) error {
	if r.m[id] == nil {
		return tenant.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
