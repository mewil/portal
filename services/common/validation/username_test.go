package validation_test

import (
	"testing"

	. "github.com/mewil/portal/common/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateUsername(t *testing.T) {
	assert.True(t, ValidUsername("mewil"))
	assert.True(t, ValidUsername("_mewil_"))
	assert.True(t, ValidUsername("9mewil0"))
	assert.True(t, ValidUsername("10000"))
	assert.False(t, ValidUsername("mew"))
	assert.False(t, ValidUsername("&mewil"))
	assert.False(t, ValidUsername("mew il"))
	assert.False(t, ValidUsername("_mewil__mewil__mewil__mewil__mewil__mewil__mewil__mewil__mewil__mewil_"))
	assert.False(t, ValidUsername(""))
}
