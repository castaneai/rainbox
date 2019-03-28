package rainbox

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cloud.google.com/go/firestore"
)

type UserID string

const (
	InvalidUserID = UserID("")
)

type User struct {
	ID          UserID `json:"id" firestore:"-"`
	DisplayName string `json:"displayName" firestore:"displayName"`
}

func newUser(id UserID, displayName string) *User {
	return &User{
		ID:          id,
		DisplayName: displayName,
	}
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (sv *UserService) Register(ctx context.Context, userID UserID, displayName string) error {
	user := newUser(userID, displayName)
	if err := sv.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (sv *UserService) Get(ctx context.Context, userID UserID) (*User, error) {
	return sv.repo.Get(ctx, userID)
}

type UserRepository interface {
	Get(ctx context.Context, id UserID) (*User, error)
	Create(ctx context.Context, u *User) error
}

type FirestoreUserRepository struct {
	store *firestore.Client
}

func NewFirestoreUserRepository(store *firestore.Client) UserRepository {
	return &FirestoreUserRepository{store: store}
}

func (repo *FirestoreUserRepository) Get(ctx context.Context, id UserID) (*User, error) {
	ref := repo.store.Doc("users/" + string(id))
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
	user.ID = UserID(ds.Ref.ID)
	return user, nil
}

func (repo *FirestoreUserRepository) Create(ctx context.Context, u *User) error {
	ref := repo.store.Doc("users/" + string(u.ID))
	if _, err := ref.Create(ctx, u); err != nil {
		return err
	}
	return nil
}
