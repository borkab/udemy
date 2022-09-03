package saying

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	s := "Welcome home"
	got := Greet(s)
	want := "Welcome home"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func ExampleGreet() {
	s := "Welcome home"
	fmt.Println(Greet(s))
	//Output:
	//Welcome home
}

func BenchmarkGreet(b *testing.B) {
	s := "Welcome my dear"
	for i := 0; i < b.N; i++ {
		Greet(s)
	}
}
