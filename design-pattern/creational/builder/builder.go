package builder

type iBuilder interface {
	setWindowType(windowType string)
	setDoorType(doorType string)
	setNumFloor(numFloor int)
	getHouse() house
}
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}
type house struct {
	windowType string
	doorType   string
	floor      int
}

func getBuilder() iBuilder {
	return &normalBuilder{}
}
func (b *normalBuilder) setWindowType(windowType string) {
	b.windowType = windowType
}

func (b *normalBuilder) setDoorType(doorType string) {
	b.doorType = doorType
}

func (b *normalBuilder) setNumFloor(numFloor int) {
	b.floor = numFloor
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
