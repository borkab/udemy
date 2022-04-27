package main

import "time"

type MyAwesomeCookbook struct {
	Refrigerator Refrigerator
	Oven         Oven
	Mixer        Mixer
	//	Stove        Stove
	Cupboard Cupboard
	//Tools    Tools
}

type Recipes struct {
	//TomatoSoup struct{}
	//	Pancake    struct{}
	CupCake struct {
	}
	//SemolinaPudding struct{}
}
type Ingrerdients struct {
	Flour  struct{}
	Yogurt struct{}
	Egg    struct{}
	Butter struct{}
	Sugar  struct{}
	Milk   struct{}
}

/*
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
*/
/*
type Kitchentools struct {
	//	Knife struct{} // a kes egy viselkedes, barmilyen kesed van, ossze tudod vele vagni a zoldseget,
	//tehat interface, tartalmazza a Cutting() methodot

	//Scissors  struct{} kell az egy szakacskonyvbe?
	//Fork      struct{}
	Spoon struct{}
	//	Ladle struct{}

	//	Pan struct{} // ez is inkabb viselkedes, mindegy milyen serpenyo v wok, megsul benne a rantott hus,
	// tehat interface, tartalmazhatna  Frying() methodot

	//	Spatula struct{}
	Sheet struct{}
}
*/

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
	butter := cookBook.Refrigerator.GetButter()
	milk := cookBook.Refrigerator.GetMilk()
	yogurt := cookBook.Refrigerator.GetYogurt()
	egg := cookBook.Refrigerator.GetEggs()
	flour := cookBook.Cupboard.GetFlour()
	sugar := cookBook.Cupboard.GetSugar()
	cp := &CupCake{}
	//tool := Kitchentools.Sheet{}

	if err := cookBook.Mixer.Mix(cp, butter, milk, yogurt, egg, flour, sugar); err != nil {
		return nil, err
	}

	if err := cookBook.Oven.Baking(45*time.Minute, 180, cp); err != nil {
		return nil, err
	}

	return cp, nil
}

/*
func (cookbook MyAwesomeCookbook) MakePancake() (*Pancake, error) {
	oil := cookBook.Cupboard.GetOil()
	milk := cookBook.Refrigerator.GetMilk()
	pk := &Pancake{}
}
*/
type Refrigerator interface {
	GetMilk() Milk     //Milk: undeclared name
	GetButter() Butter //Butter: undeclared name
	GetEggs() Egg      //Egg: undeclared name
	//GetMeet()   Meet
	GetYogurt() Yogurt //Yogurt: undeclared name
}

type Oven interface {
	//	Frying(d time.Duration, temp int, stuffToFry any, tool Kitchentools) error
	Baking(d time.Duration, temp int, stuffToBake any) error
}

/*
type Stove interface {
	Cooking(d time.Duration, temp int, stuffToCook, tool Kitchentools) error
	Stewing(d time.Duration, temp int, stuffToStew, tool Kitchentools) error
}
*/

type Cupboard interface {
	//	GetOil()
	GetFlour() Flour //Flour: undeclared name
	GetSugar() Sugar //Sugar: undeclared name
}
type Mixer interface {
	Mix(outPtr any, ingredients ...any) error
}

/*
type Tools interface {
	Cutting(tool Kitchentools, ingrerdients ...any) error
	Pouring(tool Kitchentools, ingrerdients ...any) error
	Stirring(tool Kitchentools, ingrerdients ...any) error
	GetSheet() Sheet
}
*/
