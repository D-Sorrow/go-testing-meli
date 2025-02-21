package application_test

import (
	"app/internal/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain_NewApp(t *testing.T) {
	applicationDefault := application.ConfigApplicationDefault{Addr: ":8080"}
	app := application.NewApplicationDefault(&applicationDefault)

	assert.NotNil(t, app)
}
