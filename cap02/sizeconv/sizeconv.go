// sizeconv conversor entre metros e pés
package sizeconv

import "fmt"

type Feet float64
type Metre float64

func (f Feet) String() string  { return fmt.Sprintf("%gft\n", f) }
func (m Metre) String() string { return fmt.Sprintf("%gm\n", m) }
