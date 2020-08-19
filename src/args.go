package main
import (
	"fmt"
	"os"
	"flag"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	carMakePtr := flag.String("make", "no car", "Vehicle")
	carRegPtr := flag.String("reg", "Not road legal", "Registration plate")
	carYearPtr := flag.Int("year", 0, "Year of manufacture")

	fmt.Println("Car args:")
	fmt.Println("Car make: ", *carMakePtr)
	fmt.Println("Registration number: ", *carRegPtr)
	fmt.Println("Car age: ", *carYearPtr)

	fmt.Println("Args with Program")
	fmt.Println(argsWithProg)

	fmt.Println("Args without Program")
	fmt.Println(argsWithoutProg)
}