package basicconcurrencyworkerpool

import (
	"fmt"
	"math"
)

// 1. นิยาม Interface ชื่อ Shape
type Shape interface {
	Area() float64
}

// 2. สร้าง Struct สำหรับ Rectangle และ Circle
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// 3. Implement method Area สำหรับ Rectangle
// สูตร: กว้าง * ยาว
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement method Area สำหรับ Circle
// สูตร: pi * r^2
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// 4. ฟังก์ชัน PrintArea ที่รับ Interface Shape
func PrintArea(s Shape) {
	fmt.Printf("The area of this shape is: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	// เรียกใช้ PrintArea โดยส่งทั้ง Rectangle และ Circle เข้าไปได้เลย
	// เพราะทั้งคู่มี Method Area() ตามที่ Interface Shape กำหนดไว้
	fmt.Print("Rectangle: ")
	PrintArea(rect)

	fmt.Print("Circle: ")
	PrintArea(circ)
}