package main

import "time"

type MyAwesomeCookbook struct { //hogyan formazom meg, h minden szepen egy oszlopban legyen automatikusan?
	Refrigerator Refrigerator
	Oven         Oven
	Mixer        Mixer
	Stove        Stove
	Cupboard     Cupboard
	Tools        Tools
}

type Recipes struct {
	TomatoSoup      struct{}
	Pancake         struct{}
	CupCake         struct{}
	SemolinaPudding struct{}
}

type Ingrerdients struct {
	Dairy
	Grains
	Proteins
	Vegetables
	Fruits
}
type Vegetables struct {
	TomatoPaste struct{}
	Potatoe     struct{}
	Onion       struct{}
}
type Grains struct {
	Grits   struct{}
	Flour   struct{}
	Noodles struct{}
}
type Proteins struct {
	Egg  struct{}
	Meet struct{}
}
type Dairy struct {
	Butter struct{}
	Milk   struct{}
	Yogurt struct{}
}
type Spices struct {
	Oil       struct{}
	Salt      struct{}
	Sugar     struct{}
	Basilikum struct{}
}
type Kitchentools struct {
	Knife     struct{}
	Scissors  struct{}
	Fork      struct{}
	Spoon     struct{}
	Ladle     struct{}
	FryingPan struct{}
	Spatula   struct{}
	Sheet     struct{}
}

/*
type (
	TomatoSoup struct{}
	Pancake struct{}
	CupCake struct{}
	Butter  struct{}
	Milk    struct{}

)
*/

func (cookBook MyAwesomeCookbook) MakeCupcake() (*CupCake, error) {
	oil := cookBook.Cupboard.GetOil()
	milk := cookBook.Refrigerator.GetMilk()
	yogurt := cookBook.Refrigerator.GetYogurt()
	egg := cookBook.Refrigerator.GetEggs()
	flour := cookBook.Cupboard.GetFlour()
	sugar := cookBook.Cupboard.GetSugar()
	cp := &CupCake{}
	tool := Kitchentools.Sheet{}

	if err := cookBook.Mixer.Mix(cp, oil, milk, yogurt, egg, flour, sugar); err != nil {
		return nil, err
	}

	if err := cookBook.Oven.Baking(45*time.Minute, 180, cp); err != nil {
		return nil, err
	}

	return cp, nil
}

func (cookbook MyAwesomeCookbook) MakePancake() (*Pancake, error) {
	oil := cookBook.Cupboard.GetOil()
	milk := cookBook.Refrigerator.GetMilk()
	pk := &Pancake{}
}

type Refrigerator interface {
	GetMilk() Dairy    //with type of Milk it was invalid type
	GetButter() Dairy  //with type of Butter it was invalid type
	GetEggs() Proteins //with type of Egg it was invalid type
	GetMeet() Proteins //with type of Meet it was invalid type
	GetYogurt() Dairy  //with type of Yogurt it was invalid type

}

type Oven interface {
	Frying(d time.Duration, temp int, stuffToFry Recipes, tool Kitchentools) error  //stuffToFry any invalid type
	Baking(d time.Duration, temp int, stuffToBake Recipes, tool Kitchentools) error //stuffToBake any invalid type

}
type Stove interface {
	//mixed named and unnamed parameterssyntax by Cooking(d )
	Cooking(d time.Duration, temp int, stuffToCook Recipes, tool Kitchentools) error //stuffToCook invalid type
	Stewing(d time.Duration, temp int, stuffToStew Recipes, tool Kitchentools) error //stuffToStew invalid type
	//mixed named and unnamed parameters (compile)go-staticcheck by Stewing()
}

type Cupboard interface {
	GetOil() Spices
	GetDairies() Dairy //is it ok or I should declare GetFlour() and GetNoodles() separately?
}
type Mixer interface {
	Mix(outPtr any, ingredients ...any) error // any is undeclared and invalid type
}
type Tools interface {
	Cutting(tool Kitchentools, ingrerdients ...any) error
	Pouring(tool Kitchentools, ingrerdients ...any) error
	Stirring(tool Kitchentools, ingrerdients ...any) error
}
