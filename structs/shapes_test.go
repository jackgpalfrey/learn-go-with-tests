package shapes

import "testing"

func TestPermimeter(t *testing.T){
	rect := Rectangle{10.0,10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("expected %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T){
	checkArea := func(t testing.TB, shape Shape, want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("expected %.2f, but got %.2f", want, got)
		}
	}
	
	t.Run("rectangles", func(t *testing.T){
		rect := Rectangle{12,6}
		checkArea(t, rect, 72.0)
	})

	t.Run("circles", func(t *testing.T){
		circ := Circle{10}
		want := 314.1592653589793
		checkArea(t, circ, want)

	})
}

func TestAreaTable(t *testing.T){
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12,6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12,6}, 36.0 },
	}

	for _, tt := range(areaTests) {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("expected %.2f, but got %.2f", tt.want, got)
		}
	}
}
