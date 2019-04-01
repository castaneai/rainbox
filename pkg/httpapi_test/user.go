package httpapi

import (
	"fmt"

	"github.com/castaneai/rainbox/pkg/rainbox"
	"github.com/google/uuid"
)

func NewRandomUserID() rainbox.UserID {
	return rainbox.UserID(uuid.Must(uuid.NewRandom()).String())
}

type User struct {
	ID      rainbox.UserID
	IDToken string
}

func newUser() *User {
	uid := NewRandomUserID()
	return &User{
		ID:      uid,
		IDToken: fmt.Sprintf("token-%s", uid),
	}
}
