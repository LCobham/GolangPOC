// Perform Celsius and Farenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Farenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0.0
	Boiling       Celsius = 100.0
)

// String parsing.
func (c Celsius) String() string   { return fmt.Sprintf("%.2f°C", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (k Kelvin) String() string    { return fmt.Sprintf("%.2fK", k) }

// Conversions.
func CtoF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }
func FtoC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KtoC(k Kelvin) Celsius    { return Celsius(k) - AbsoluteZeroC }
func CtoK(c Celsius) Kelvin    { return Kelvin(c + AbsoluteZeroC) }
func KtoF(k Kelvin) Farenheit  { return CtoF(KtoC(k)) }
func FtoK(f Farenheit) Kelvin  { return CtoK(FtoC(f)) }
