package pkg

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

type User struct {
	id          string
	DisplayName string `firestore:"displayName"`
}

func NewUser(fuser *auth.UserInfo) *User {
	return &User{
		id:          fuser.UID,
		DisplayName: fuser.DisplayName,
	}
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (sv *UserService) SignUp(ctx context.Context, fuser *auth.UserInfo) (*User, error) {
	u := NewUser(fuser)
	if err := sv.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (sv *UserService) SignIn(ctx context.Context, userID string) (*User, error) {
	return sv.repo.Get(ctx, userID)
}

type UserRepository interface {
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, u *User) error
}

type FirestoreUserRepository struct {
	store *firestore.Client
}

func NewFirestoreUserRepository(store *firestore.Client) *FirestoreUserRepository {
	return &FirestoreUserRepository{store: store}
}

func (repo *FirestoreUserRepository) Get(ctx context.Context, id string) (*User, error) {
	ref := repo.store.Doc("users/" + id)
	ds, err := ref.Get(ctx)
	if err != nil && status.Code(err) == codes.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var user *User
	if err := ds.DataTo(&user); err != nil {
		return nil, err
	}
	user.id = ds.Ref.ID
	return user, nil
}

func (repo *FirestoreUserRepository) Create(ctx context.Context, u *User) error {
	ref := repo.store.Doc("users/" + u.id)
	if _, err := ref.Create(ctx, u); err != nil {
		return err
	}
	return nil
}
