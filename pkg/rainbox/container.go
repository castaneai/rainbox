package rainbox

import (
	"cloud.google.com/go/firestore"
	"go.uber.org/dig"
)

type Services struct {
	dig.In

	User *UserService
	Post *PostService
}

func SetupDIContainer(c *dig.Container) error {
	if err := c.Provide(func(store *firestore.Client) UserRepository {
		return NewFirestoreUserRepository(store)
	}); err != nil {
		return err
	}

	if err := c.Provide(func(repo UserRepository) *UserService {
		return NewUserService(repo)
	}); err != nil {
		return err
	}

	if err := c.Provide(func(store *firestore.Client) PostRepository {
		return NewFirestorePostRepository(store)
	}); err != nil {
		return err
	}

	if err := c.Provide(func(repo PostRepository) *PostService {
		return NewPostService(repo)
	}); err != nil {
		return err
	}
	return nil
}
