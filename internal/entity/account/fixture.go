package account

import (
	"time"

	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

func NewFixtureUserAccount() *Account {
	return &Account{
		ID:        id.NewID(),
		Email:     "user11@cheese.com",
		Password:  "manchego",
		Actions:   "R",
		CreatedAt: time.Now(),
	}
}

func NewFixtureAdminAccount() *Account {
	return &Account{
		ID:        id.NewID(),
		Email:     "admin1@cheese.com",
		Password:  "SmokedGouda",
		Actions:   "RW",
		CreatedAt: time.Now(),
	}
}
