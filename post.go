package rainbox

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Post struct {
	authorUserID string
	AuthorRef    *firestore.DocumentRef `firestore:"authorRef"`
	ThumbnailURL string                 `firestore:"thumbnailUrl"`
	ImageURLs    []string               `firestore:"imageUrls"`
	Tags         []string               `firestore:"tags"`
	Likes        int                    `firestore:"likes"`
}

func NewPost(author *User, imageURLs []string) *Post {
	return &Post{
		authorUserID: author.UserID,
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
	post.AuthorRef = pr.store.Doc("users/" + post.authorUserID)
	if _, _, err := pr.store.Collection("posts").Add(ctx, post); err != nil {
		return err
	}
	return nil
}
