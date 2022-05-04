package billkennedy

import (
	"time"
)

type MyAwesomeCookbook struct {
	Refrigerator Refrigerator
	Oven         Oven
	Mixer        Mixer
	//	Stove        Stove
	Cupboard Cupboard
	//Tools    Tools
}

type Cupcake struct {
	HasDressing bool
	IsRaw       bool
}

//	TomatoSoup Recipe
//	Pancake Recipe
//	SemolinaPudding Recipe

type Ingredient int

const ( // enum
	Flour Ingredient = iota
	Yogurt
	Egg
	Butter
	Sugar
	Milk

//	TomatoPaste Ingredient
//	Potatoe     Ingredient
//	Onion       Ingredient
// 	Egg 		Ingredient
// 	Meet    	Ingredient
//	Oil			Ingredient
//	Salt		Ingredient
//	Basilikum 	Ingredient
)

type Recipe struct{} //csak h meg tudjam kulonboztetni hozzavalot es receptet

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

type unit string

const (
	Spoon unit = "spoon"
	Mug   unit = "mug"
	Cup   unit = "cup"
	Etc   unit = "etc"
)

type TODOGiveThisANormalNameAndUseRefactoringToRenameIt struct {
	Ingredient Ingredient
	Quantity   int
	Unit       unit
}

func (cookBook MyAwesomeCookbook) MakeCupcake() (*Cupcake, error) { //why is CupCake undeclared? I didn't change it
	butter := cookBook.Refrigerator.GetButter()
	milk := cookBook.Refrigerator.GetMilk(3, Cup)
	yogurt := cookBook.Refrigerator.GetYogurt()
	egg := cookBook.Refrigerator.GetEggs()
	flour := cookBook.Cupboard.GetFlour()
	sugar := cookBook.Cupboard.GetSugar()

	type X struct {
		someField string
	}

	x := X{}
	_ = x

	cp := &Cupcake{}
	//tool := Kitchentools.Sheet{}

	cp1 := cp // pass by value -> cp is a pointer -> cp of the pointer -> both pointer points to the same variable
	_ = cp1

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
	GetMilk(quantity int, unit unit) TODOGiveThisANormalNameAndUseRefactoringToRenameIt
	GetButter() Ingredient
	GetEggs() Ingredient
	//GetMeet()   Meet
	GetYogurt() Ingredient
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
	GetFlour() Ingredient //Flour: undeclared name
	GetSugar() Ingredient //Sugar: undeclared name
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
