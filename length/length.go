package length

type Length struct {
	meters float64
}

/// Creates a Length struct from yoctometers.
func FromYoctometer(yoctometers float64) Length {
	return Length { meters: yoctometersToMeters(yoctometers) }
}

/// Creates a Length struct from meters.
func FromMeters(meters float64) Length {
	return Length { meters }
}

/// Creates a Length struct from feet.
func FromFeet(feet float64) Length {
	return Length { meters: feetToMeters(feet) }
}
