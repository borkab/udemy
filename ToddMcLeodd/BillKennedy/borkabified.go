package main

import "time"

type MyAwesomeCookbook struct {
	Refrigerator Refrigerator
	Oven         Oven
	Mixer        Mixer
}

type (
	CupCake struct{}
	Butter  struct{}
	Milk    struct{}
)

func (cookBook MyAwesomeCookbook) MakeCupcake() (*CupCake, error) {
	butter := cookBook.Refrigerator.GetButter()
	milk := cookBook.Refrigerator.GetMilk()
	cp := &CupCake{}

	if err := cookBook.Mixer.Mix(cp, butter, milk); err != nil {
		return nil, err
	}

	if err := cookBook.Oven.Frying(30*time.Minute, 190, cp); err != nil {
		return nil, err
	}

	return cp, nil
}

type Refrigerator interface {
	GetMilk() Milk
	GetButter() Butter
}

type Oven interface {
	Frying(d time.Duration, temp int, stuffToFry any) error
}

type Mixer interface {
	Mix(outPtr any, ingredients ...any) error
}
