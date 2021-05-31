package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/joaovicdsantos/gopl/cap02/sizeconv"
	"github.com/joaovicdsantos/gopl/cap02/tempconv"
	"github.com/joaovicdsantos/gopl/cap02/weightconv"
)

var tipo string

func main() {
	flag.StringVar(&tipo, "t", "temperatura", "Tipo dos valores que será convertido: temperatura, peso e tamanho")

	flag.Parse()

	for _, v := range flag.Args() {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "O valor %v é inválido\n", v)
		}
		switch tipo {
		case "temperatura":
			convertTemp(value)
		case "peso":
			convertWeight(value)
		case "tamanho":
			convertSize(value)
		default:
			fmt.Fprint(os.Stderr, "Tipo inválido!")
		}
	}

}

func convertTemp(value float64) {
	valueInCelsius := tempconv.Celsius(value)
	valueInFahrenheit := tempconv.Fahrenheit(value)
	valueInKelvin := tempconv.Kelvin(value)
	fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n\n", valueInCelsius, tempconv.CToF(valueInCelsius), valueInCelsius, tempconv.CToK(valueInCelsius), valueInFahrenheit, tempconv.FToC(valueInFahrenheit), valueInCelsius, tempconv.FToK(valueInFahrenheit), valueInKelvin, tempconv.KToC(valueInKelvin), valueInKelvin, tempconv.KToF(valueInKelvin))
}

func convertWeight(value float64) {
	valueInKilo := weightconv.Kilo(value)
	valueInPound := weightconv.Pound(value)
	fmt.Printf("%s = %s\n%s = %s\n\n", valueInKilo, weightconv.KgToLb(valueInKilo), valueInPound, weightconv.LbToKg(valueInPound))
}

func convertSize(value float64) {
	valueInMetre := sizeconv.Metre(value)
	valueInFeet := sizeconv.Feet(value)
	fmt.Printf("%s = %s\n%s = %s\n\n", valueInMetre, sizeconv.MtoFt(valueInMetre), valueInFeet, sizeconv.FtToM(valueInFeet))
}
