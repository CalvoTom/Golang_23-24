package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

type Door struct {
	state int
}

const (
	CLOSE = 0
	OPEN  = 1
)

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	z01.PrintRune('\n')
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	z01.PrintRune('\n')
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	z01.PrintRune('\n')
	return ptrDoor.state == 1
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	z01.PrintRune('\n')
	return ptrDoor.state == 0
}

func main() {
	door := &Door{
		state: CLOSE,
	}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
