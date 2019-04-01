package httpapi

import (
	"context"
	"fmt"

	"github.com/castaneai/rainbox/pkg/rainbox"
)

type verifier struct {
	users map[string]*User // idToken -> *clientUser
}

func newStaticVerifier(users []*User) rainbox.Verifier {
	mp := make(map[string]*User)
	for _, u := range users {
		mp[u.IDToken] = u
	}
	return &verifier{
		users: mp,
	}
}

func (v *verifier) Verify(ctx context.Context, idToken string) (rainbox.UserID, error) {
	cuser, ok := v.users[idToken]
	if ok && cuser != nil {
		return cuser.ID, nil
	}
	return rainbox.InvalidUserID, fmt.Errorf("failed to verify id token")
}
