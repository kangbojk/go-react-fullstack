package account_pg

import (
	"database/sql"
	"time"

	"github.com/kangbojk/go-react-fullstack/pkg/entity/account"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

//pgRepo postgresql repo
type pgRepo struct {
	db *sql.DB
}

//NewPgRepository create new postgresql repository
func NewPgRepository(db *sql.DB) *pgRepo {
	return &pgRepo{
		db: db,
	}
}

//Get an Account
func (r *pgRepo) Get(id id.ID) (*account.Account, error) {
	stmt, err := r.db.Prepare(`select id, email, password, actions, tenant_id, created_time, updated_time from account where id = ?`)
	if err != nil {
		return nil, err
	}
	var account_ account.Account
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&account_.ID, &account_.Email, &account_.Password, &account_.Actions, &account_.TenantID, &account_.CreatedAt, &account_.UpdatedAt)
	}
	return &account_, nil
}

func (r *pgRepo) FindUserWithEmail(email string) (*account.Account, error) {
	stmt, err := r.db.Prepare(`select id, email, password, actions, tenant_id, created_time, updated_time from account where email = ?`)
	if err != nil {
		return nil, err
	}
	var account_ account.Account
	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&account_.ID, &account_.Email, &account_.Password, &account_.Actions, &account_.TenantID, &account_.CreatedAt, &account_.UpdatedAt)
	}
	return &account_, nil
}

//List Accounts
func (r *pgRepo) List() ([]*account.Account, error) {
	stmt, err := r.db.Prepare(`select id, email, password, actions, tenant_id, created_time, updated_time from account`)
	if err != nil {
		return nil, err
	}
	var accounts []*account.Account
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var account_ account.Account
		err = rows.Scan(&account_.ID, &account_.Email, &account_.Password, &account_.Actions, &account_.TenantID, &account_.CreatedAt, &account_.UpdatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account_)
	}
	if len(accounts) == 0 {
		return nil, account.ErrInvalid
	}
	return accounts, nil
}

//Create an account
func (r *pgRepo) Create(acc *account.Account) (id.ID, error) {
	stmt, err := r.db.Prepare(`
			insert into account (id, email, password, actions, tenant_id, created_time, updated_time)
			values(?,?,?,?,?,?,?)`)
	if err != nil {
		return acc.ID, err
	}

	// TODO Hash password
	_, err = stmt.Exec(
		acc.ID,
		acc.Email,
		acc.Password,
		acc.Actions,
		acc.TenantID,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return acc.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return acc.ID, err
	}
	return acc.ID, nil
}

//Update an account
func (r *pgRepo) Update(acc *account.Account) error {
	acc.UpdatedAt = time.Now()
	_, err := r.db.Exec("update account set email = ?, password = ?, actions = ?, tenant_id = ?, updated_time = ? where id = ?", acc.Email, acc.Password, acc.Actions, acc.TenantID, acc.UpdatedAt.Format("2006-01-02"), acc.ID)
	if err != nil {
		return err
	}
	return nil
}

//Delete an account
func (r *pgRepo) Delete(idx id.ID) error {
	_, err := r.db.Exec("delete from account where id = ?", idx)
	if err != nil {
		return err
	}
	return nil
}
