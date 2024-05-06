package length_test

import (
	"testing"

	"github.com/asungy/unit/length"
	"github.com/stretchr/testify/assert"
)

func TestConversion(t *testing.T) {
	l := length.FromMeters(42)
	assert.Equal(t, 42., l.Meters())
	assert.InDelta(t, 137.795, l.Feet(), 0.001)
}

func TestAdd(t *testing.T) {
	a := length.FromMeters(42)
	b := length.FromMeters(13)
	actual := a.Add(b)
	expected := length.FromMeters(55)
	assert.Equal(t, actual.Equal(expected), true)
}

func TestEqualWithDelta(t *testing.T) {

}
