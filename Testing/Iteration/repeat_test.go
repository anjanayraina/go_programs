package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("A", 5)
	expected := "AAAAA"
	if repeated != expected {
		t.Errorf("The expected value is %q but we got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A", 5)
	}
}
