package main

import (
	"os"
	"fmt"
	"bufio"
	"math"
	"bytes"
	"strings"
)

func main() {
	f,err := os.Open(os.Args[1])

	if err != nil {
		os.Exit(1)
	}

	br := bufio.NewReader(f)
	byte_array := bytes.Buffer{}
    keep := true
	count := 0.0

	for {
		line, err := br.ReadString('\n')
        if err != nil{
            break
        } else {
        	count++
        	if math.Mod(count,4) ==1 {
        		// fmt.Println("reset")
        		byte_array.Reset()
        		byte_array.Write([]byte(line))
        	} else if math.Mod(count,4) ==2 {
        		byte_array.Write([]byte(line))
        		if strings.Contains(line, "N") {
        			keep = false
        		} else {
        			keep = true
        		}
        	} else if math.Mod(count,4) ==3{
            	byte_array.Write([]byte(line))
        	} else {
            	byte_array.Write([]byte(line))
            	if keep {
            		fmt.Print(byte_array.String())
            	}
        	}
        }
	}
}