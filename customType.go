package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}


type Box struct {
	width, height, depth float64
	color Color
}
func(box Box) Volume() float64 {
	return box.width * box.height * box.depth
}

// 这里的receive是一个指针，因为目的是为了改变当前box的颜色
// 而不是通过副本改变副本的颜色
func(box * Box) SetColor(_color Color) {
	// 这里go会知道要访问的是box的指针
	// 所以*box.coloer  === box.color
	box.color = _color
}


type BoxList []Box
func(bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}


func main(){
	boxies := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Println(boxies)
}