package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFake(t *testing.T) {

	t.Run("Phone", func(t *testing.T) {
		phone := Phone()
		assert.Len(t, phone, 11)
	})

	t.Run("Email", func(t *testing.T) {
		email := Email()
		assert.Contains(t, email, "@gmail.com")
	})
}
