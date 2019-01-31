package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var chars = []byte("abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ0123456789!@#$%&*()[]{}?;:/*-+.,\\|<>áéíóúãẽĩõũàèìòùäëïöüÁÉÍÓÚÃẼĨÕŨÀÈÌÒÙÄËÏÖÜ")

func main() {
	var size int
	if len(os.Args) <= 1 {
		size = 16
	} else {
		sizeArg, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			panic(err)
		}

		size = int(sizeArg)
	}

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	randomizer := rand.New(source)

	securePass := ""
	lastIndex := -1
	for i := 0; i < int(size); i++ {
		var index int
		for {
			index = randomizer.Int() % len(chars)
			if index != lastIndex {
				lastIndex = index
				break
			}
		}

		securePass += string(chars[index])
	}

	fmt.Println(securePass)
}
