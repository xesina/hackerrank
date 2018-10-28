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

		arTemp := strings.Split(readLine(ir), " ")
		var ar []int32

		for i := 0; i < int(n); i++ {
			arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
			checkError(err)
			arItem := int32(arItemTemp)
			ar = append(ar, arItem)
		}

		r := sockMerchant(n, ar)

		fo, err := os.Open(filepath.Join(outPrefix, outputs[k].Name()))
		checkError(err)
		defer fo.Close()
		or := bufio.NewReaderSize(fo, 1024*1024)

		oTemp, err := strconv.ParseInt(readLine(or), 10, 64)
		checkError(err)

		o := int32(oTemp)

		if r != o {
			t.Fatalf(`FAIL: sockMerchant(%d, %v) = %d want %d`, n, ar, r, o)
		}
		t.Logf("PASS:sockMerchant(%d, %v) = %d want %d`", n, ar, r, o)
	}
}

func BenchmarkAddGigasecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	}
}
