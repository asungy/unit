package length

import (
	"math"
)

/// Returns true if `l` and `other` have equivalent values.
func (l Length) Equal(other Length) bool {
	return l.meters == other.meters
}

/// Returns true if `l` and `other` have equivalent values within a delta.
func (l Length) EqualWithDelta(other Length, delta float64) bool {
	return math.Abs(l.meters - other.meters) < delta
}

/// Return true if value of `l` is greater than value of `other`.
func (l Length) GreaterThan(other Length) bool {
	return l.meters > other.meters
}
