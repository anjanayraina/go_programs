package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	var result = Perimeter(10, 20)
	value := float32(60)

	if result != value {
		t.Errorf("The expected value is : %f but the recieved value is : %f", result, value)
	}
}

func TestArea(t *testing.T) {
	var result = Area(10, 20)
	value := float32(200)
	if result != value {
		t.Errorf("The expected value is : %f but the recieved value is : %f", result, value)
	}
}
