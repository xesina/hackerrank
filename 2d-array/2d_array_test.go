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

		var arr [][]int32
		for i := 0; i < 6; i++ {
			arrRowTemp := strings.Split(readLine(ir), " ")

			var arrRow []int32
			for _, arrRowItem := range arrRowTemp {
				arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
				checkError(err)
				arrItem := int32(arrItemTemp)
				arrRow = append(arrRow, arrItem)
			}

			if len(arrRow) != int(6) {
				panic("Bad input")
			}

			arr = append(arr, arrRow)
		}

		r := hourglassSum(arr)

		fo, err := os.Open(filepath.Join(outPrefix, outputs[k].Name()))
		checkError(err)
		defer fo.Close()
		or := bufio.NewReaderSize(fo, 1024*1024)

		oTemp, err := strconv.ParseInt(readLine(or), 10, 64)
		checkError(err)
		o := int32(oTemp)

		if r != o {
			t.Fatalf(`FAIL: hourglassSum(%v) = %d want %d`, arr, r, o)
		}
		t.Logf("PASS:hourglassSum(%v) = %d want %d`", arr, r, o)
	}
}

func BenchmarkAddGigasecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hourglassSum([][]int32{
			[]int32{1, 1, 1, 0, 0, 0},
			[]int32{0, 1, 0, 0, 0, 0},
			[]int32{1, 1, 1, 0, 0, 0},
			[]int32{0, 0, 2, 4, 4, 0},
			[]int32{0, 0, 0, 2, 0, 0},
			[]int32{0, 0, 1, 2, 4, 0},
		})
	}
}
