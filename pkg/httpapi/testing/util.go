package testing

import (
	"context"
	"fmt"

	"github.com/castaneai/rainbox/pkg/rainbox"
	"github.com/google/uuid"
)

func newRandomUserID() rainbox.UserID {
	return rainbox.UserID(uuid.Must(uuid.NewRandom()).String())
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

var testClientUser1 = &clientUser{ID: newRandomUserID(), IDToken: "token1"}
var testClientUser2 = &clientUser{ID: newRandomUserID(), IDToken: "token2"}
var testClientUser3 = &clientUser{ID: newRandomUserID(), IDToken: "token3"}

var testClientUsers = []*clientUser{
	testClientUser1, testClientUser2, testClientUser3,
}
