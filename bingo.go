package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("welcome to bingo")
	fmt.Println("Enter upper limit :")
	var limit int
	var anykey = make([]byte, 1)
	fmt.Scan(&limit)

	var n [100]int
	var i, count int

	for i = 0; i < limit; i++ {
		n[i] = 0
	}

	count = limit

	for true {

		if count == 0 {
			break
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		b := r.Int() % limit
		if n[b] == 0 {
			fmt.Println(b)
			n[b] = 1
			count--
			fmt.Println("Press any key to continue")
			os.Stdin.Read(anykey)
		}

	}
}
