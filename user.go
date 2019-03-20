package rainbox

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
)

type User struct {
	UserID string `firestore:"userId"`
	DisplayName string `firestore:"displayName"`
}

func NewUser(fuser *auth.UserInfo) *User {
	return &User{
		UserID: fuser.UID,
		DisplayName: fuser.DisplayName,
	}
}

type UserRespotiroy struct {
	store *firestore.Client
}

func NewUserRepository(store *firestore.Client) (*UserRespotiroy, error) {
	return &UserRespotiroy{
		store: store,
	}, nil
}

func (ur *UserRespotiroy) SignIn(ctx context.Context, fuser *auth.UserInfo) (*User, error) {
	user := NewUser(fuser)
	if _, err := ur.store.Collection("users").Doc(fuser.UID).Set(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}