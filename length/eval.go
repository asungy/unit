package length

func (l Length) Meters() float64 {
	return l.meters
}

func (l Length) Feet() float64 {
	return metersToFeet(l.meters)
}

