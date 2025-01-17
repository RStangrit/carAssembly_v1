package wheels

type Wheels struct {
	Tires Tire
	Rims  Rim
}

type Tire struct {
	Size int
}

type Rim struct {
	Size int
}

func NewWheels(rimSize, tireSize int) Wheels {
	Rim := NewRim(rimSize)
	Tire := NewTire(tireSize)
	return Wheels{
		Tires: Tire,
		Rims:  Rim,
	}
}

func NewTire(tireSize int) Tire {
	return Tire{
		Size: tireSize,
	}
}

func NewRim(rimSize int) Rim {
	return Rim{
		Size: rimSize,
	}
}
