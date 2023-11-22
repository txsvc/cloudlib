package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitObserver(t *testing.T) {
	assert.Nil(t, googleStorageProvider)
}
