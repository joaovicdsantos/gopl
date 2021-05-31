package weightconv

// LbToKg converte um valor em libras para kilo
func LbToKg(p Pound) Kilo { return Kilo(p / 2.205) }

// KgToLb converte um valor em kilo para libras
func KgToLb(k Kilo) Pound { return Pound(k * 2.205) }
