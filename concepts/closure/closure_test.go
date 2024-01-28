package closure

import (
	"fmt"
	"testing"
)

/*
Kinds of Tests which are available in Go:
- Unit Test (`testing.T`)
- Benchmark Test (`testing.B`)
- Example Test (see the example below)
- Fuzzy Test (not explored more)
- Main Test (added to have setup and tear down code before other functions are run)

Helpful commands:
```sh
go test -v -run TestOddInt # to run specific test cases
go test packageName -run TestName # to run single test in a package
go test -v -test.short # to only execute tests which are not marked as short; useful in separating unit and integration tests
go test -cover # to check test coverage
```

Ref:
https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file
https://freshman.tech/snippets/go/run-specific-test/
https://pkg.go.dev/testing # Documentation
https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/ # comprehensive


*/

var iocases = []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
var iecases = []int{0, 2, 4, 6, 8, 10, 12, 14, 16}
var iacases = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

func TestMain(m *testing.M) {
	fmt.Println("Main test started")
	m.Run()
	fmt.Println("Main test ended")
}

func TestNextInt(t *testing.T) {
	// Table driven tests
	var tests = []struct {
		name string
		want int
	}{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
		{"ten", 10},
	}
	nxt := nextInt()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel() // this runs the test in parellel
			t.Logf("Want %d\n", tt.want)
			got := nxt()
			if got != tt.want {
				t.Errorf("Got %d; Expected %d\n", got, tt.want)
			}
		})
	}
}

func TestOddInt(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	nxt := OddInt()

	for _, expect := range iocases {
		if out := nxt(); out != expect {
			t.Errorf("Got %d; Expected %d\n", out, expect)
		}
	}
}

func TestEvenInt(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	nxt := EvenInt()

	for _, expect := range iecases {
		if out := nxt(); out != expect {
			t.Errorf("Got %d; Expected %d\n", out, expect)
		}
	}
}

func TestAddNextInt(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	nxt := NextNo()

	for _, expect := range iacases {
		if out := nxt(); out != expect {
			t.Errorf("Got %d; Expected %d\n", out, expect)
		}
	}
}

// Benchmark Test
func BenchmarkNextInt(b *testing.B) {
	nxt := nextInt()

	for i := 0; i <= b.N; i++ {
		nxt()
	}
}

func BenchmarkOddInt(b *testing.B) {
	nxt := OddInt()

	for i := 0; i <= b.N; i++ {
		nxt()
	}
}

func BenchmarkAddNo(b *testing.B) {
	nxt := NextNo()

	for i := 0; i <= b.N; i++ {
		nxt()
	}
}

func ExampleNextNo() {
	fmt.Println(NextNo()())
	// Output: 0
}

func ExampleEvenInt() {
	fmt.Println(EvenInt()())
	// Output: 0
}

func ExampleOddInt() {
	fmt.Println(OddInt()())
	// Output:
	// 1
}

func ExamplenextInt() {
	fmt.Println(nextInt()())
	// Output:
	// 1
}
