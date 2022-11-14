package main

import (
	"LAB_01/cell"
	"fmt"
	"os"
	"strconv"
)

var (
	N            int     = 16
	K            int     = 32
	p            float64 = 0.5
	timeToRunSec int64   = 10
	isSafe       bool    = true
)

var sharedArray []cell.Cell
var savedTotal int

func prepare() {
	fmt.Println("###################### PREPARATION ###########################")

	savedTotal = 0

	for i := 0; i < N; i++ {
		var c cell.Cell
		if isSafe {
			c = &cell.SafeCell{}
		} else {
			c = &cell.UnsafeCell{}
		}
		val := K / N

		c.SetValue(val)
		c.SetIndex(i)
		savedTotal += val
		sharedArray = append(sharedArray, c)
	}

	fmt.Printf("values: %v\n", sharedArray)
	fmt.Printf("total:  %d\n", savedTotal)
}

func run() {
	fmt.Println("###################### RUN ###################################")
	ch := make(chan int)
	for _, c := range sharedArray {
		go c.Run(timeToRunSec, sharedArray, p, ch)
	}
	for i := 0; i < N; i++ {
		i := <-ch
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func aftermath() {
	fmt.Println("###################### AFTERMATH #############################")

	total := 0
	for _, v := range sharedArray {
		total += v.GetValue()
	}

	fmt.Printf("new values: %v\n", sharedArray)
	fmt.Printf("new total:  %d\n", total)
}

func printUsage() {
	println("Wrong arguments!")
	println("usage:")
	println("    app.exe N K p s")
	println()
	println("        N - number of cells         | int > 0")
	println("        K - sum of atoms            | int > 0, >= N")
	println("        p - probability of movement | float 0 <= p <= 1")
	println("        s - use safe function       | int 0 or 1")
}

func main() {
	args := os.Args[1:]

	if len(args) != 4 {
		printUsage()
		return
	}

	var err error

	N, err = strconv.Atoi(args[0])
	if err != nil {
		printUsage()
		return
	}

	K, err = strconv.Atoi(args[1])
	if err != nil {
		printUsage()
		return
	}

	p, err = strconv.ParseFloat(args[2], 64)
	if err != nil {
		printUsage()
		return
	}

	if K < N {
		fmt.Println("K must be >= N")
		return
	}

	s, err := strconv.Atoi(args[3])
	if err != nil {
		printUsage()
		return
	}

	if s != 0 && s != 1 {
		printUsage()
		return
	}

	if s == 0 {
		isSafe = false
	} else {
		isSafe = true
	}

	prepare()
	run()
	aftermath()

}
