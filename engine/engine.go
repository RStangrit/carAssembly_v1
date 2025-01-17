package engine

type Engine struct {
	HorsePowers  int
	Volume       int
	TorqueMoment int
}

func New(horsePowers, volume, torqueMoment int) Engine {
	return Engine{
		HorsePowers:  horsePowers,
		Volume:       volume,
		TorqueMoment: torqueMoment,
	}
}
