package main

import (
	"os"
	"fmt"
	"bufio"
	"math"
)

func main() {
	f,err := os.Open(os.Args[1])

	if err != nil {
		os.Exit(1)
	}

	br := bufio.NewReader(f)
	count := 0.0

	for {
		line, err := br.ReadString('\n')
        if err != nil{
            break
        } else {
        	count++
        	if math.Mod(count,4) ==2 {
        		fmt.Print(line)
        	}
        }
	}
}