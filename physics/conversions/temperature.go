package conversions

// Temperature t
type Temperature struct {
	Value float64
	Type  string
}

// ToCelsius t
func (t *Temperature) ToCelsius() {
	if t.Type == "K" {
		t.Value = t.Value - 273.15
		t.Type = "C"
	} else if t.Type == "F" {
		t.Value = (t.Value - 32.0) * (5.0 / 9.0)
		t.Type = "C"
	} else if t.Type == "C" {
		return
	} else {
		panic("Temperature Type Is Wrong")
	}
}

// ToFahrenheit t
func (t *Temperature) ToFahrenheit() {
	if t.Type == "C" {
		t.Value = (t.Value * (9.0 / 5.0)) + 32.0
		t.Type = "F"
	} else if t.Type == "K" {
		t.Value = ((t.Value - 273.15) * (9.0 / 5.0)) + 32.0
		t.Type = "F"
	} else if t.Type == "F" {
		return
	} else {
		panic("Temperature Type Is Wrong")
	}
}

// ToKelvin t
func (t *Temperature) ToKelvin() {
	if t.Type == "C" {
		t.Value = t.Value + 273.15
		t.Type = "K"
	} else if t.Type == "F" {
		t.Value = ((t.Value - 32.0) * (5.0 / 9.0)) + 273.15
		t.Type = "K"
	} else if t.Type == "K" {
		return
	} else {
		panic("Temperature Type Is Wrong")
	}
}
