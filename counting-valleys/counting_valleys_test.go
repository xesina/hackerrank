package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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
		
		s := readLine(ir)
		r := countingValleys(n, s)

		fo, err := os.Open(filepath.Join(outPrefix, outputs[k].Name()))
		checkError(err)
		defer fo.Close()
		or := bufio.NewReaderSize(fo, 1024*1024)

		oTemp, err := strconv.ParseInt(readLine(or), 10, 64)
		checkError(err)

		o := int32(oTemp)

		if r != o {
			t.Fatalf(`FAIL: countingValleys(%d, %s) = %d want %d`, n, s, r, o)
		}
		t.Logf("PASS:sockMerchant(%d, %s) = %d want %d`", n, s, r, o)
	}
}

func BenchmarkAddGigasecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countingValleys(8, "UDDDUDUU")
	}
}
