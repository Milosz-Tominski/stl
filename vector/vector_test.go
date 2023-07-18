package vector

import "testing"

func TestOperators(t *testing.T) {
	vec := vector[float64]([]float64{1.0, 2.0, 3.0})
	if vec.At(0) != 1.0 {
		t.Errorf("Expected 1.0, got %.2f", vec.At(0))
	}
	if vec.Front() != 1.0 {
		t.Errorf("Expected 1.0, got %.2f", vec.Front())
	}
	if vec.Back() != 3.0 {
		t.Errorf("Expected 3.0, got %.2f", vec.Back())
	}
	if vec.Size() != 3 {
		t.Errorf("Expected 3, got %d", vec.Size())
	}
	vec.Clear()
	if !vec.Empty() {
		t.Errorf("Expected empty, got full")
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
	}
	vec.Pop(2.0)
	if vec.At(0) != 3.0 {
		t.Errorf("Expected 3.0, got %.2f", vec.At(0))
	}
	vec.Insert(2.0, 0)
	if vec.At(0) != 2.0 {
		t.Errorf("Expected 2.0, got %.2f", vec.At(0))
	}
}

func TestSwap(t *testing.T) {
	vec1 := vector[string]([]string{"Hi", "Hello", "World"})
	vec2 := vector[string]([]string{"Go", "Lang"})
	t.Log(vec1, vec2)
	Swap(vec1, vec2)
	if vec1.Size() != 2 {
		t.Errorf("Expected 2, got %d", vec1.Size())
	}
	if vec2.Size() != 3 {
		t.Errorf("Expected 3, got %d", vec2.Size())
	}
}
