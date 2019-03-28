package httpapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"

	"github.com/google/uuid"
	"go.uber.org/dig"

	"github.com/castaneai/rainbox/pkg/rainbox"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
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

func setupDIContainerForTesting(ctx context.Context, c *dig.Container) error {
	if err := c.Provide(func() (*firebase.App, error) {
		return firebase.NewApp(ctx, nil)
	}); err != nil {
		return err
	}
	if err := c.Provide(func(app *firebase.App) (*firestore.Client, error) {
		return app.Firestore(ctx)
	}); err != nil {
		return err
	}
	if err := c.Provide(func() rainbox.Verifier {
		return NewTestClientVerifier(testClientUsers)
	}); err != nil {
		return err
	}

	if err := rainbox.SetupDIContainer(c); err != nil {
		return err
	}
	return nil
}

type TestSuite struct {
	suite.Suite
	Handler http.Handler
	server  *httptest.Server
}

func (suite *TestSuite) SetupSuite() {
	ctx := context.Background()
	c := dig.New()
	if err := setupDIContainerForTesting(ctx, c); err != nil {
		log.Fatal(err)
	}
	if err := c.Invoke(func(v rainbox.Verifier, services rainbox.Services) {
		suite.Handler = NewHandler(v, &services)
		suite.server = httptest.NewServer(suite.Handler)
	}); err != nil {
		log.Fatal(err)
	}
}

func (suite *TestSuite) TearDownSuite() {
	defer suite.server.Close()
}
