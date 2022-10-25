package main

type VeggieMania struct {
	pizza IPizza
}

func (p *VeggieMania) getPrice() int {
	return 15
}
