package sizeconv

// FtToM converte pés para metros
func FtToM(f Feet) Metre { return Metre(f / 3.2808) }

// MtoFt converte metros para pés
func MtoFt(m Metre) Feet { return Feet(m * 3.2808) }
