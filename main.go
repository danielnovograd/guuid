package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
)

var number = flag.Int("n", 1, "how many uuids")

func main() {
	flag.Parse()

	for i := 0; i < *number; i++ {
		fmt.Println(uuid.NewString())
	}
}
