package rainbox

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type Post struct {
	AuthorUserID UserID    `firestore:"authorUserId"`
	ThumbnailURL string    `firestore:"thumbnailUrl"`
	ImageURLs    []string  `firestore:"imageUrls"`
	Tags         []string  `firestore:"tags"`
	Likes        int       `firestore:"likes"`
	CreatedAt    time.Time `firestore:"createdAt"`
	UpdatedAt    time.Time `firestore:"updatedAt"`
}

func NewPost(author *User, imageURLs []string) *Post {
	return &Post{
		AuthorUserID: author.ID,
		ImageURLs:    imageURLs,
	}
}

type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (sv *PostService) CountAll(ctx context.Context) (int, error) {
	return sv.repo.CountAll(ctx)
}

type PostRepository interface {
	CountAll(context.Context) (int, error)
}

type FirestorePostRepository struct {
	store *firestore.Client
}

func NewFirestorePostRepository(store *firestore.Client) PostRepository {
	return &FirestorePostRepository{
		store: store,
	}
}

func (repo *FirestorePostRepository) CountAll(ctx context.Context) (int, error) {
	// 😥
	iter := repo.store.Collection("posts").DocumentRefs(ctx)
	all, err := iter.GetAll()
	if err != nil {
		return 0, err
	}
	return len(all), nil
}
