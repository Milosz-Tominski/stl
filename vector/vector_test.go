package vector 

import "testing"


func TestOperators(t *testing) {
	vec := vector[float64]([]float64{1.0, 2.0, 3.0})
	if vec.At(0) != 1.0 {
		t.Errorf("Expected 1.0, got %.2f", vec.At(0))
	}
	if vec.Front() != 1.0 {
		t.Errorf("Expected 1.0, got %.2f", vec.At(0))
	}
	if vec.Back() != 3.0 {
		t.Errorf("Expected 3.0, got %.2f", vec.At(0))
	}
	if vec.Size() != 3 {
		t.Errorf("Expected 3, got %d", vec.Size())
	}
	vec.Clear()
	if !vec.Empty() {
		t.Errorf("Expected clear, got full")
	}
	vec.Push(1.0)
	vec.Push(2.0)
	vec.Push(3.0)
	if vec.Size() != 3 {
		t.Errorf("Expected 3, got %d", vec.Size())
	}
	vec.Erase(0)
	if vec.At(0) != 2.0 {
		t.Errorf("Expected 2.0, got %.2f", vec.At(0))
	vec.Pop(2.0)
	if vec.At(0) != 3.0 {
		t.Errorf("Expected 2.0, got %.2f", vec.At(0))
	vec.Insert(2.0, 0)
	if vec.At(0) != 2.0 {
		t.Errorf("Expected 2.0, got %.2f", vec.At(0))
	}
}

func TestSwap(t *testing.T) {
	vec1 := vector[int]([]int{1,2,3})
	vec2 := vector[int]([]int{4,5})
	Swap(vec1, vec2)
	if vec1.Size() != 2 {
		t.Errorf("Expected 2, got %d", vec.Size())
	}
	if vec2.Size() != 3 {
		t.Errorf("Expected 3, got %d", vec.Size())
	}
}