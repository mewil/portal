package validation_test

import (
	"testing"

	. "github.com/mewil/portal/common/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	assert.True(t, ValidEmail("mewil@umich.edu"))
	assert.False(t, ValidEmail("mewil@umichedu"))
	assert.False(t, ValidEmail("@umich.edu"))
	assert.False(t, ValidEmail("mewil @ umich . edu"))
	assert.False(t, ValidEmail("mewilumich.edu"))
	assert.False(t, ValidEmail("mewilumichedu"))
	assert.False(t, ValidEmail("mewil@umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu.umich.edu"))
	assert.False(t, ValidEmail("mewil@u_mich.edu"))
	assert.False(t, ValidEmail("mewil@umich._edu"))
	assert.False(t, ValidEmail(""))
}
