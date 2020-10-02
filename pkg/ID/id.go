// package id is used for generating new id
package id

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func IDToString(id ID) string {
	s := ID.String(id)
	return s
}