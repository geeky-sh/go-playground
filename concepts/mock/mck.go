package mck

import (
	"fmt"
	"time"
)

/*
Goal:
To use mock package to test non-deterministic parts of the code.

Details:
Non deterministic part refers to the code whose values we can't control.
Like random number generation, get current time, division by zero, etc

Do:
The first step to do is to extract out non-deterministic parts of the code into a separate interface.
Next step is to mock the interface.
Next step is to pass the mocked object instead of the actual one while initlialising the object
Next step is to assume and pass values which will be used with mocked functions.
Last step is the actually call the function to be tested.

Example:
1. Test current time
2. Test getting uuid

Ref:
https://dev.to/salesforceeng/mocks-in-go-tests-with-testify-mock-6pd
https://github.com/stretchr/testify#installation
*/

type currentTime interface {
	Get() time.Time
}

type MyUUID interface {
	Get() string
}

type Siri struct {
	c currentTime
	u MyUUID
}

func (r Siri) SayCurrentTime() string {
	return fmt.Sprintf("The answer is %s\n", r.c.Get())
}

func (r Siri) GiveUUID() string {
	return fmt.Sprintf("The answer is %s\n", r.u.Get())
}
