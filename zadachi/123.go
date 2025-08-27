package main

import "fmt"

type Animals interface {
	Voice()
}

type Dog struct {
	Name string
}

func (d *Dog) Voice() {
	fmt.Printf("%s: Woof\n", d.Name)
}

type Cat struct {
	Name string
}

func (c *Cat) Voice() {
	fmt.Printf("%s: Meow\n", c.Name)
}

func MakeVoice(a Animals) {
	a.Voice()
}

func main() {
	animals := []Animals{
		&Dog{Name: "Sharik"},
		&Cat{Name: "Murka"},
	}

	for _, animal := range animals {
		MakeVoice(animal)
	}
}

/*
import "fmt"

func ValueInfo(obj interface{}) {
	switch val := obj.(type) {
	case string:
		fmt.Printf("Длина строки: %d\n", len(val))
	case []int:
		fmt.Println("Емкость слайса:", cap(val))
	default:
		fmt.Printf("Тип %T", val)
	}
}

func main() {
	ValueInfo("str")      // длина строки: 3
	ValueInfo([]int{1, 2}) // емкость слайса: 2
	ValueInfo(true)        // тип bool
}

/*
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// реализация метода Area интерфейса Shape
func (r Rectangle) Area() float64 {
	if r.Height > 0 && r.Width > 0 {
		return r.Height * r.Width
	}
	return -1
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rectangle := Rectangle{5, 6}
	fmt.Println(GetArea(rectangle)) // 30
}

/*
type User struct {
	ID   int
	Name string
	Age  int
}

func (u *User) GetUserInfo() {
	fmt.Printf("%d-%s-%d\n", u.ID, u.Name, u.Age)
}

func (u *User) ChangeUserInfo() {
	u.Name = "Вася"
	u.Age = 20
	fmt.Printf("Изменение: %d-%s-%d\n", u.ID, u.Name, u.Age)
}

func main() {
	user1 := User{1, "Гоша", 18}
	user1.GetUserInfo()
	user1.ChangeUserInfo()
	user1.GetUserInfo()
}

*/
