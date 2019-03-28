package pkg

import (
	"context"
	"testing"

	firebase "firebase.google.com/go"

	"github.com/stretchr/testify/assert"

	"firebase.google.com/go/auth"

	"cloud.google.com/go/firestore"
)

const (
	testUserID   = "1c53d4e4-6c2a-4fb6-8ee2-be22b5f395ce"
	testUserName = "test-user-name"
)

func newTestFirebaseApp(ctx context.Context) (*firebase.App, error) {
	return firebase.NewApp(ctx, nil)
}

func newTestFirestore(ctx context.Context) (*firestore.Client, error) {
	app, err := newTestFirebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}

func newTestUserRepository(ctx context.Context) (UserRepository, error) {
	store, err := newTestFirestore(ctx)
	if err != nil {
		return nil, err
	}
	return &FirestoreUserRepository{store}, nil
}

func newTestUserService(ctx context.Context) (*UserService, error) {
	repo, err := newTestUserRepository(ctx)
	if err != nil {
		return nil, err
	}
	return &UserService{repo}, nil
}

func newTestFirebaseUser() *auth.UserInfo {
	return &auth.UserInfo{
		UID:         testUserID,
		DisplayName: testUserName,
	}
}

func TestSignUp(t *testing.T) {
	ctx := context.Background()
	sv, err := newTestUserService(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fuser := newTestFirebaseUser()
	u, err := sv.SignUp(ctx, fuser)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fuser.UID, u.id)
	assert.Equal(t, fuser.DisplayName, u.DisplayName)
}
