package transmission

type Transmission struct {
	GearboxType GearboxType
	Clutch      Clutch
}

type GearboxType struct {
	Type string
}

type Clutch struct {
	Type string
}

func NewTransmission(gearboxType, clutchType string) Transmission {
	gearBox := NewGearbox(gearboxType)
	clutch := NewClutch(clutchType)
	return Transmission{
		GearboxType: gearBox,
		Clutch:      clutch,
	}
}

func NewGearbox(gearboxType string) GearboxType {
	return GearboxType{
		Type: gearboxType,
	}
}

func NewClutch(clutchType string) Clutch {
	return Clutch{
		Type: clutchType,
	}
}
