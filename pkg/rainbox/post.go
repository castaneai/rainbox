package rainbox

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type PostID string

type Post struct {
	ID           PostID    `firestore:"-"`
	AuthorUserID UserID    `firestore:"authorUserId"`
	ThumbnailURL string    `firestore:"thumbnailUrl"`
	ImageURLs    []string  `firestore:"imageUrls"`
	Tags         []string  `firestore:"tags"`
	Likes        int       `firestore:"likes"`
	CreatedAt    time.Time `firestore:"createdAt"`
	UpdatedAt    time.Time `firestore:"updatedAt"`
}

func NewPost(author *User, imageURLs, tags []string) *Post {
	return &Post{
		AuthorUserID: author.ID,
		ImageURLs:    imageURLs,
		Tags:         tags,
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

func (sv *PostService) Create(ctx context.Context, post *Post) error {
	return sv.repo.Create(ctx, post)
}

func (sv *PostService) CountAll(ctx context.Context) (int, error) {
	return sv.repo.CountAll(ctx)
}

type PostRepository interface {
	Create(context.Context, *Post) error
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

func (repo *FirestorePostRepository) Create(ctx context.Context, post *Post) error {
	if _, _, err := repo.store.Collection("posts").Add(ctx, post); err != nil {
		return err
	}
	return nil
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
