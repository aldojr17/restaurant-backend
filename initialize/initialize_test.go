package initialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitApp(t *testing.T) {
	app := App()
	assert.NotNil(t, app.DB)
}
