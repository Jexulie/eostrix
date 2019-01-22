package conversions

// Speed p
type Speed struct {
	Value float64
	Type  string
}

// ToMPH s
func (s *Speed) ToMPH() {
	if s.Type == "KPH" {
		s.Value = s.Value * 0.621371
		s.Type = "MPH"
	} else if s.Type == "FPS" {
		s.Value = s.Value / 1.467
		s.Type = "MPH"
	} else if s.Type == "Knot" {
		s.Value = s.Value * 1.151
		s.Type = "MPH"
	} else {
		panic("Speed Type Is Wrong")
	}
}

// ToKPH s
func (s *Speed) ToKPH() {
	if s.Type == "MPH" {
		s.Value = s.Value * 1.6093440
		s.Type = "KPH"
	} else if s.Type == "FPS" {
		s.Value = s.Value * 1.097
		s.Type = "KPH"
	} else if s.Type == "Knot" {
		s.Value = s.Value * 1.852
		s.Type = "KPH"
	} else {
		panic("Speed Type Is Wrong")
	}
}

// ToFPS s
func (s *Speed) ToFPS() {
	if s.Type == "MPH" {
		s.Value = s.Value * 1.467
		s.Type = "FPS"
	} else if s.Type == "KPS" {
		s.Value = s.Value / 1.097
		s.Type = "FPS"
	} else if s.Type == "Knot" {
		s.Value = s.Value * 1.688
		s.Type = "FPS"
	} else {
		panic("Speed Type Is Wrong")
	}
}

// ToKnot s
func (s *Speed) ToKnot() {
	if s.Type == "MPH" {
		s.Value = s.Value / 1.151
		s.Type = "Knot"
	} else if s.Type == "FPS" {
		s.Value = s.Value / 1.688
		s.Type = "Knot"
	} else if s.Type == "KPH" {
		s.Value = s.Value / 1.852
		s.Type = "Knot"
	} else {
		panic("Speed Type Is Wrong")
	}
}
