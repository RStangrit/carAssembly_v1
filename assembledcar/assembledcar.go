package AssembledCar

import (
	"main/body"
	"main/engine"
	"main/propsgenerator"
	"main/transmission"
	"main/wheels"
)

type AssembledCar struct {
	Engine       engine.Engine
	Transmission transmission.Transmission
	Wheels       wheels.Wheels
	Body         body.Body
}

func NewCar() AssembledCar {

	horsePowers, volume, torqueMoment, gearboxType, clutchType, rimSize, tireSize, bodyMaterial, bodyColor, bodyTrunkVolume, bodyRoof := propsgenerator.GenerateRandomCarProperties()
	return AssembledCar{
		Engine:       engine.New(horsePowers, volume, torqueMoment),
		Transmission: transmission.NewTransmission(gearboxType, clutchType),
		Wheels:       wheels.NewWheels(rimSize, tireSize),
		Body:         body.NewBody(bodyMaterial, bodyColor, bodyTrunkVolume, bodyRoof),
	}
}
