package main 

import "fmt"
import "time"
import "math/rand"

//go:noinline
func Greet(name string) {
	fmt.Println("Hello, " + name)
}

func main() {
	names := []string{"Mauro", "Lucas", "Kere123456"}
	tick := time.Tick(1 * time.Second)

	for range tick {
		Greet(names[rand.Intn(len(names))])
	}
}