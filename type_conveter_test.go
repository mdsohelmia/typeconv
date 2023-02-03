package gotypconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValueToPointer(t *testing.T) {
	strValue := "Hello Gotipath"
	prtStr := String(strValue)
	assert.Equal(t, &strValue, prtStr)
}

func TestStringPointerToValueString(t *testing.T) {
	expValue := "Hello Gotipath"

	value := StringValue(&expValue)

	assert.Equal(t, expValue, value)

}

func TestBooleanToPointerBoolean(t *testing.T) {
	t.Run("it boolean on true", func(t *testing.T) {
		value := true
		output := Bool(value)
		assert.Equal(t, &value, output)
	})

	t.Run("it boolean on false", func(t *testing.T) {
		value := false
		output := Bool(value)
		assert.Equal(t, &value, output)
	})
}
