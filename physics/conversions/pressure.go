package conversions

// Pressure p
type Pressure struct {
	Value float64
	Type  string
}

// ToBar s
func (p *Pressure) ToBar() {
	if p.Type == "ATM" {
		p.Value = p.Value * 1.013
		p.Type = "Bar"
	} else if p.Type == "Pascal" {
		p.Value = p.Value / 100000
		p.Type = "Bar"
	} else if p.Type == "Torr" {
		p.Value = p.Value / 750.062
		p.Type = "Bar"
	} else {
		panic("Speed Type Is Wrong")
	}
}

// ToATM s
func (p *Pressure) ToATM() {
	if p.Type == "Bar" {
		p.Value = p.Value / 1.013
		p.Type = "ATM"
	} else if p.Type == "Pascal" {
		p.Value = p.Value / 101325
		p.Type = "ATM"
	} else if p.Type == "Torr" {
		p.Value = p.Value / 760
		p.Type = "ATM"
	} else {
		panic("Speed Type Is Wrong")
	}
}

// ToPascal s
func (p *Pressure) ToPascal() {
	if p.Type == "Bar" {
		p.Value = p.Value * 100000
		p.Type = "Pascal"
	} else if p.Type == "ATM" {
		p.Value = p.Value * 101325
		p.Type = "Pascal"
	} else if p.Type == "Torr" {
		p.Value = p.Value * 133.322
		p.Type = "Pascal"
	} else {
		panic("Speed Type Is Wrong")
	}
}

// ToTorr s
func (p *Pressure) ToTorr() {
	if p.Type == "ATM" {
		p.Value = p.Value / 760
		p.Type = "Torr"
	} else if p.Type == "Bar" {
		p.Value = p.Value / 750.062
		p.Type = "Torr"
	} else if p.Type == "Pascal" {
		p.Value = p.Value / 133.322
		p.Type = "Torr"
	} else {
		panic("Speed Type Is Wrong")
	}
}
