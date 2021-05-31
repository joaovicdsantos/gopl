package weightconv

import "fmt"

type Pound float64
type Kilo float64

func (p Pound) String() string { return fmt.Sprintf("%glb\n", p) }

func (k Kilo) String() string { return fmt.Sprintf("%gkg\n", k) }
