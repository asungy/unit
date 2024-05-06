package length

/// Returns the sum of two Length structs as a new Length struct.
func (l Length) Add(other Length) Length {
	return Length {
		meters: l.meters + other.meters,
	}
}

/// Returns the difference between two Length structs as a new Length stuct.
func (l Length) Sub(other Length) Length {
	return Length {
		meters: l.meters - other.meters,
	}
}

/// Returns the product of two Length structs as a new Length struct.
func (l Length) Mul(other Length) Length {
	return Length {
		meters: l.meters * l.meters,
	}
}

/// Returns the quotient of two Length structs as a new Length struct.
func (l Length) Div(other Length) Length {
	return Length {
		meters: l.meters / l.meters,
	}
}
