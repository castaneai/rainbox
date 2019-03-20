package rainbox

import (
	"context"
	"testing"
)

func TestNewPost(t *testing.T) {
	ctx := context.Background()
	store, err := newTestFirestore(ctx)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewPostRepository(store)

	user, err := newTestUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
	post := NewPost(user, []string{"image01", "image02"})
	if err := repo.Save(ctx, post); err != nil {
		t.Fatal(err)
	}
}
