package rainbox

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

const (
	validIDToken = "valid-id-token"
)

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

func newTestUser() *User {
	uid := uuid.Must(uuid.NewRandom()).String()
	return &User{
		id:          UserID(uid),
		DisplayName: fmt.Sprintf("name-of-%s", uid),
	}
}

func TestRegister(t *testing.T) {
	ctx := context.Background()
	sv, err := newTestUserService(ctx)
	if err != nil {
		t.Fatal(err)
	}

	user := newTestUser()
	if err := sv.Register(ctx, user); err != nil {
		t.Fatal(err)
	}

	suser, err := sv.Get(ctx, user.id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user.id, suser.id)
	assert.Equal(t, user.DisplayName, suser.DisplayName)
}
