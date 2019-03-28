package rainbox

import (
	"context"

	"firebase.google.com/go/auth"
)

type Verifier interface {
	Verify(ctx context.Context, idToken string) (UserID, error)
}

type FirebaseAuthVerifier struct {
	client *auth.Client
}

func NewFirebaseAuthVerifier(client *auth.Client) Verifier {
	return &FirebaseAuthVerifier{client: client}
}

func (fav *FirebaseAuthVerifier) Verify(ctx context.Context, idToken string) (UserID, error) {
	token, err := fav.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return InvalidUserID, err
	}
	return UserID(token.UID), nil
}
