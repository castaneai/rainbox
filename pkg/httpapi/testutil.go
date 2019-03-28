package httpapi

import (
	"context"
	"fmt"
	"github.com/google/uuid"

	"github.com/castaneai/rainbox/pkg/rainbox"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func newRandomUserID() rainbox.UserID {
	return rainbox.UserID(uuid.Must(uuid.NewRandom()).String())
}

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

type clientUser struct {
	ID      rainbox.UserID
	IDToken string
}

type testClientVerifier struct {
	users map[string]*clientUser // idToken -> *clientUser
}

func NewTestClientVerifier(users []*clientUser) rainbox.Verifier {
	mp := make(map[string]*clientUser)
	for _, u := range users {
		mp[u.IDToken] = u
	}
	return &testClientVerifier{
		users: mp,
	}
}

func (tcv *testClientVerifier) Verify(ctx context.Context, idToken string) (rainbox.UserID, error) {
	cuser, ok := tcv.users[idToken]
	if ok && cuser != nil {
		return cuser.ID, nil
	}
	return rainbox.InvalidUserID, fmt.Errorf("failed to verify id token")
}
