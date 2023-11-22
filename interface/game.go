package main

import "fmt"

func main() {
	var loc Location
	fmt.Println(loc)
	fmt.Printf("+: %+v\n", loc)
	fmt.Printf("#: %#v\n", loc) // Use this in logging/debugging

	loc.X = 7 // no setters/getters. You can use unexported fields
	loc.Y = 19
	fmt.Println(loc)

	loc2 := Location{22, 33} // fields by position
	fmt.Println(loc2)

	loc3 := Location{ // fields by name
		Y: 9,
		//		X: 17,
	}
	fmt.Println(loc3)

	if loc, err := NewLocation(-1, 200); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("loc:", loc)
	}

	if loc, err := NewLocation(1, 200); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("loc:", loc)
	}

	// Use NewLocation is not mandatory
	loc4 := Location{-100, 10038}
	fmt.Println("loc4:", loc4)
	// Go proverb: Make the zero value useful

	loc4.Move(0, 0)
	fmt.Println("loc4 (move):", loc4)

	p1 := Player{
		Name:     "Parzival",
		Location: Location{500, 500},
	}
	fmt.Println("p1", p1)
	fmt.Printf("%s at <%v,%v>\n", p1.Name, p1.X, p1.Y)
	p1.Move(900, 200)
	fmt.Println("p1 (move)", p1)
	// There's an implicit Location field in Player
	fmt.Println("p1 location:", p1.Location)

	items := []Mover{
		&p1,
		&Car{"OUTATIME", Location{0, 0}},
	}
	moveAll(items, 20, 30)
	fmt.Printf("items: %#v\n", items)

	k := Jade
	fmt.Println("k:", k)

	k1 := Key(77) // invalid key
	fmt.Println("k1:", k1)
}

func (p *Player) FoundKey(key Key) error {
	/* TODO: Implement
	   - Check that key is a valid key
	   - Add a key only once
	*/
	if key.isValidKey() && !p.ContainsKey(key){
		p.Keys = append(p.Keys, key)
	}

}

func (p *Player) ContainsKey(key Key) bool {
	for _, k := range p.Keys {
		if k == key {
			return true
		}
	}
	return false
}

func (k Key) isValidKey() (bool, error) {
	switch k {
		case Copper:
			return true
		case Jade:
			return true
		case Crystal:
			return true
	}
	return fmt.Errorf("invalid key: %d", k)
}

// implement the fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key: %d>", k)
	// return fmt.Sprintf("<Key: %v>", k) // ∞ recursion
}

// Enum like
type Key uint8

const (
	Copper Key = iota + 1 // iota grows by 1 for each line in a const group
	Jade
	Crystal
)

// You can also do bitmasks with: Copper Key = 1 << iota

/* Interface "best practices":
   - Interfaces are defined at usage point
   - Interfaces are small
   - Accept interfaces, return concrete types (in functions)
*/
type Mover interface {
	Move(int, int)
}

func moveAll(items []Mover, x, y int) {
	for _, item := range items {
		item.Move(x, y)
	}
}

type Car struct {
	Plate string
	Location
}

// Google for "embedding vs extending"

type Player struct {
	Name string
	Keys []Key
	// X        float64 // X will "shadow" Location.X
	Location // Player embeds Location (not inherit from Location)
}

// l is called the method receiver
// When changing fields, use a pointer receiver
func (l *Location) Move(x, y int) {
	l.X = x
	l.Y = y
}

func NewLocation(x, y int) (*Location, error) {
	if x < 0 || x > MaxX || y < 0 || y > MaxY {
		return nil, fmt.Errorf("%d, %d out of bounds", x, y)
	}

	loc := Location{x, y}
	return &loc, nil // Go does escape analysis, and will allocate loc on the heap
}

const (
	MaxX = 1000
	MaxY = 1000
)

type Location struct {
	X int
	Y int
}

/* Getter/Setter
type Location struct {
	x int
	y int
}

func (l Location) X() int {
	return l.x
}

func (l *Location) SetX(x int) int {
	l.x = x
}
*/
