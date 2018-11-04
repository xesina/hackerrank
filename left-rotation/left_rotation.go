package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leftRot(a []int32, d int32) []int32 {
	l := len(a)
	var r int32
	for r = 1; r <= d; r++ {
		for i := 0; i < l-1; i++ {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
	return a
}

func leftRotSlicing(a []int32, d int32) []int32 {
	var r int32
	for r = 1; r <= d; r++ {
		a = append(a[1:], a[0])
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bufSize := 1024 * 1024
	buf := make([]byte, bufSize)
	scanner.Buffer(buf, bufSize)

	scanner.Scan()
	l1 := strings.Split(scanner.Text(), " ")

	dTemp, err := strconv.ParseInt(l1[1], 10, 64)
	checkErr(err)
	d := int32(dTemp)

	scanner.Scan()
	aTemp := strings.Split(scanner.Text(), " ")

	a := make([]int32, len(aTemp))

	for k, v := range aTemp {
		itemTemp, err := strconv.ParseInt(v, 10, 64)
		checkErr(err)
		a[k] = int32(itemTemp)
	}

	result := leftRotSlicing(a, d)

	of, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkErr(err)
	defer of.Close()

	w := bufio.NewWriterSize(of, 1024*1024)

	for i, r := range result {
		fmt.Fprintf(w, "%d", r)

		if i < len(result)-1 {
			fmt.Fprintf(w, " ")
		}
	}

	fmt.Fprintf(w, "\n")
	w.Flush()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
