package tempconv

import (
	"fmt"
)

type Fahrenheit float64
type Celsius float64

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°", f) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°", c) }
