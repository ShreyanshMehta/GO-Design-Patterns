package main

import "errors"

type IFactory interface {
	createVehicle(string) (IVehicle, error)
}

func createFactory(factoryName string) (IFactory, error) {
	var factory IFactory
	if factoryName == "tata" {
		factory = &TataFactory{}
	} else if factoryName == "jeep" {
		factory = &JeepFactory{}
	} else {
		return nil, errors.New("invalid factory")
	}
	return factory, nil
}

type TataFactory struct {
	IFactory
}

func (t *TataFactory) createVehicle(vehicle string) (IVehicle, error) {
	if vehicle == "car" {
		return &TataCar{}, nil
	} else if vehicle == "bike" {
		return &TataBike{}, nil
	} else {
		return nil, errors.New("invalid vehicle")
	}
}

type JeepFactory struct {
	IFactory
}

func (j *JeepFactory) createVehicle(vehicle string) (IVehicle, error) {
	if vehicle == "car" {
		return &JeepCar{}, nil
	} else if vehicle == "bike" {
		return &JeepBike{}, nil
	} else {
		return nil, errors.New("invalid vehicle")
	}
}
