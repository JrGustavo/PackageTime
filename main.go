package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println("Año:", now.Year())
	fmt.Println("Mes:", now.Month())
	fmt.Println("Dia:", now.Day())
	fmt.Println("Hora:", now.Hour())
	fmt.Println("Minutos:", now.Minute())
	fmt.Println("Segundos:", now.Second())

}
