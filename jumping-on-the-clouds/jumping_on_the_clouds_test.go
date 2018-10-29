package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

var (
	inPrefix  = "input"
	outPrefix = "output"
)

func TestAddGigasecond(t *testing.T) {
	inputs, err := ioutil.ReadDir("./input")
	checkError(err)
	outputs, err := ioutil.ReadDir("./output")
	checkError(err)

	for k, i := range inputs {
		fi, err := os.Open(filepath.Join(inPrefix, i.Name()))
		checkError(err)
		defer fi.Close()
		ir := bufio.NewReaderSize(fi, 1024*1024)

		nTemp, err := strconv.ParseInt(readLine(ir), 10, 64)
		checkError(err)
		n := int32(nTemp)

		cTemp := strings.Split(readLine(ir), " ")
		var c []int32

		for i := 0; i < int(n); i++ {
			cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
			checkError(err)
			cItem := int32(cItemTemp)
			c = append(c, cItem)
		}

		r := jumpingOnClouds(c)

		fo, err := os.Open(filepath.Join(outPrefix, outputs[k].Name()))
		checkError(err)
		defer fo.Close()
		or := bufio.NewReaderSize(fo, 1024*1024)

		oTemp, err := strconv.ParseInt(readLine(or), 10, 64)
		checkError(err)

		o := int32(oTemp)

		if r != o {
			t.Fatalf(`FAIL: jumpingOnClouds(%v) = %d want %d`, c, r, o)
		}
		t.Logf("PASS:jumpingOnClouds(%v) = %d want %d`", c, r, o)
	}
}

func BenchmarkAddGigasecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0})
	}
}
