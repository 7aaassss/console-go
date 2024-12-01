package crs

import (
	"fmt"
	"math/rand"
	"time"
)

func Handle_error(err error) {
	if err != nil {
		panic(err)
	}
}

func Pass(lenn int) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var s string
	for i := 0; i < lenn; i++ {
		a := string(rune(random.Intn(89) + 33))
		s += a
	}

	fmt.Println(s)
}
