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

type PostRepository struct {
	store *firestore.Client
}

func NewPostRepository(store *firestore.Client) *PostRepository {
	return &PostRepository{
		store: store,
	}
}

func (pr *PostRepository) Save(ctx context.Context, post *Post) error {
	if _, _, err := pr.store.Collection("posts").Add(ctx, post); err != nil {
		return err
	}
	return nil
}
