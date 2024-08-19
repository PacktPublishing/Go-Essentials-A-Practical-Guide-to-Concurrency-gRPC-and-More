package adult

import (
	"errors"
)

type Adult struct {
	Age  int
	Name string
}

func (a *Adult) SetAge(newAge int) error {
	if newAge < 1 {
		return errors.New("invalid age")
	}
	a.Age = newAge
	return nil
}

func (a *Adult) GetAge() int {
	return a.Age
}