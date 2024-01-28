package mck

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// define mocks
type mockCurrentTime struct {
	mock.Mock
}

func (r *mockCurrentTime) Get() time.Time {
	args := r.Called()
	return args.Get(0).(time.Time)
}

type mockMyUUID struct {
	mock.Mock
}

func (r *mockMyUUID) Get() string {
	args := r.Called()
	return args.String(0)
}

func TestCurrentTime(t *testing.T) {
	n := time.Now()
	u := uuid.NewString()

	mc := mockCurrentTime{}
	mc.On("Get").Return(n)

	mu := mockMyUUID{}
	mu.On("Get").Return(u)

	s := Siri{c: &mc, u: &mu}

	want := fmt.Sprintf("The answer is %s\n", n)
	got := s.SayCurrentTime()
	if got != want {
		t.Errorf("Got %s, Wanted %s", got, want)
	}

	want = fmt.Sprintf("The answer is %s\n", u)
	got = s.GiveUUID()
	if got != want {
		t.Errorf("Got %s, Wanted %s", got, want)
	}

	want = fmt.Sprintf("The answer is 3\n")
}
