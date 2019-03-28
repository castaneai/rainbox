package testing

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/castaneai/rainbox/pkg/httpapi"

	"github.com/stretchr/testify/suite"

	"go.uber.org/dig"

	"github.com/castaneai/rainbox/pkg/rainbox"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

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
		suite.Handler = httpapi.NewHandler(v, &services)
		suite.server = httptest.NewServer(suite.Handler)
	}); err != nil {
		log.Fatal(err)
	}
}

func (suite *TestSuite) TearDownSuite() {
	defer suite.server.Close()
}
