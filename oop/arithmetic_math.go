package arithmeticmath

import (
	"fmt"
)

type arithmeticmath struct {
	Firstname   string
	Lastname    string
	Totalleaves int
	Leavestaken int
}

func (error arithmeticmath) LeaveRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", error.Firstname, error.Lastname, (error.Totalleaves, error.Leavestaken))
}
