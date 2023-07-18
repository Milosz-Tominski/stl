package vector 


import (
	"fmt"
	"errors"
)




type Signed interface { int | int8 | int16 | int32 | int64 }

type Unsigned interface { uint | uint8 | uint16 | uint32 | uint64 }

type FloatingPoint interface { float32 | float64 }

type Text interface { byte | string | rune }

type vectorType interface { Signed | Unsigned | FloatingPoint | Text | bool }



type Vector[Type vectorType] struct {
	Length int 
	Values []Type
}



func (vector *Vector[Type]) At(index int) Type { return vector.Values[index] }

func (vector *Vector[Type]) Front() Type { return vector.Values[0] }

func (vector *Vector[Type]) Back() Type { return vector.Values[len(vector.Values)-1] }

func (vector *Vector[Type]) Size() int { return vector.Length }

func (vector *Vector[Type]) Display() { fmt.Printf("%#v\n", vector) }



func (vector *Vector[Type]) Empty() bool {
	if len(vector.Values) != 0 {
		return false
	}
	return true
}


func (vector *Vector[Type]) Clear() {
	vector.Values = []Type{}
	vector.Length = 0
}


func Swap[Type vectorType](vector1, vector2 *Vector[Type]) {
	vector1.Values, vector2.Values = vector2.Values, vector1.Values
	vector1.Length, vector2.Length = vector2.Length, vector1.Length
}


func (vector *Vector[Type]) Push(element Type) {
	vector.Values = append(vector.Values, element)
	vector.Length = len(vector.Values)
}


func (vector *Vector[Type]) Erase(index int) {
	vector.Values = append(vector.Values[:index], vector.Values[index+1:]...)
	vector.Length = len(vector.Values)
}


func (vector *Vector[Type]) Pop(element Type) {
	index := func(arr []Type, element Type) int {
		for index, value := range arr {
			if value == element {
				return index
			}
		}
		panic(errors.New("Element not found"))
	}(vector.Values, element)
	vector.Values = append(vector.Values[:index], vector.Values[index+1:]...)
	vector.Length = len(vector.Values)
}


func (vector *Vector[Type]) Insert(element Type, index int) {
	if index >= vector.Length+1 {
		panic(errors.New("Element cannot be inserted at this index"))
	}
	if index == vector.Length {
		vector.Push(element)
		return
	}
	vector.Values = append(vector.Values, element)
	copy(vector.Values[index+1:], vector.Values[index:])
	vector.Values[index] = element
	vector.Length = len(vector.Values)
}


func vector[Type vectorType](arr []Type) *Vector[Type] {
	return &Vector[Type]{
		Length: len(arr),
		Values: arr,
	}
}

func Add[Type vectorType](vec1 *Vector[Type], vec2 *Vector[Type]) *Vector[Type] {
	values := append(vector1.Values, vector2.Values...)
	return vector[Type](values)
}