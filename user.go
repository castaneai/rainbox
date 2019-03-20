package rainbox

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

type User struct {
	UserID      string `firestore:"userId"`
	DisplayName string `firestore:"displayName"`
}

func NewUser(fuser *auth.UserInfo) *User {
	return &User{
		UserID:      fuser.UID,
		DisplayName: fuser.DisplayName,
	}
}

type UserRespotiroy struct {
	store *firestore.Client
}

func NewUserRepository(store *firestore.Client) *UserRespotiroy {
	return &UserRespotiroy{
		store: store,
	}
}

func (ur *UserRespotiroy) SignIn(ctx context.Context, fuser *auth.UserInfo) (*User, error) {
	ref := ur.store.Collection("users").Doc(fuser.UID)
	ds, err := ref.Get(ctx)
	if err != nil && status.Code(err) == codes.NotFound {
		user := NewUser(fuser)
		if _, err := ref.Create(ctx, user); err != nil {
			return nil, err
		}
		return user, nil
	}
	if err != nil {
		return nil, err
	}

	var user *User
	if err := ds.DataTo(&user); err != nil {
		return nil, err
	}
	return user, nil
}
