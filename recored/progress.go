package recored

import "fmt"

type progress struct {
	Total int
	Done  int
}

var Progresss = progress{0, 0}

func GetProgress() string {
	return fmt.Sprintf("%v/%v", Progresss.Done, Progresss.Total)
}
