package rainbox

import (
	"context"
	"fmt"
)

const (
	testUserID = "test-user-id"
)

type MockAuthVerifier struct{}

func (mav *MockAuthVerifier) Verify(ctx context.Context, idToken string) (UserID, error) {
	if idToken == validIDToken {
		return testUserID, nil
	}
	return InvalidUserID, fmt.Errorf("failed to verify id token")
}
