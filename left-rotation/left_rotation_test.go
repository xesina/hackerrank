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

const (
	InPrefix  = "input"
	OutPrefix = "output"
)

func TestLeftRot(t *testing.T) {
	inputs, err := ioutil.ReadDir("./input")
	checkErr(err)
	outputs, err := ioutil.ReadDir("./output")
	checkErr(err)

	for i, file := range inputs {
		fd, err := os.Open(filepath.Join(InPrefix, file.Name()))
		checkErr(err)
		defer fd.Close()

		scanner := bufio.NewScanner(fd)
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
		for j, item := range aTemp {
			itemTemp, err := strconv.ParseInt(item, 10, 64)
			checkErr(err)
			a[j] = int32(itemTemp)
		}

		result := leftRot(a, d)

		ofd, err := os.Open(filepath.Join(OutPrefix, outputs[i].Name()))
		checkErr(err)
		defer ofd.Close()

		scanner2 := bufio.NewScanner(ofd)
		scanner2.Buffer(buf, bufSize)
		scanner2.Scan()
		oTemp := strings.Split(scanner2.Text(), " ")
		if len(oTemp) != len(result) {
			t.Fatal(`FAIL: len(result) != len(expected)`)
		}
		out := make([]int32, len(oTemp))
		for j, item := range oTemp {
			itemTemp, err := strconv.ParseInt(item, 10, 64)
			checkErr(err)
			out[j] = int32(itemTemp)
		}

		for x, y := range result {
			if y != out[x] {
				t.Fatalf(`FAIL: result != output  %v != %v`, result, out)
			}
		}

	}

}

func TestLeftRotSlicing(t *testing.T) {
	inputs, err := ioutil.ReadDir("./input")
	checkErr(err)
	outputs, err := ioutil.ReadDir("./output")
	checkErr(err)

	for i, file := range inputs {
		fd, err := os.Open(filepath.Join(InPrefix, file.Name()))
		checkErr(err)
		defer fd.Close()

		scanner := bufio.NewScanner(fd)
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
		for j, item := range aTemp {
			itemTemp, err := strconv.ParseInt(item, 10, 64)
			checkErr(err)
			a[j] = int32(itemTemp)
		}

		result := leftRotSlicing(a, d)

		ofd, err := os.Open(filepath.Join(OutPrefix, outputs[i].Name()))
		checkErr(err)
		defer ofd.Close()

		scanner2 := bufio.NewScanner(ofd)
		scanner2.Buffer(buf, bufSize)
		scanner2.Scan()
		oTemp := strings.Split(scanner2.Text(), " ")
		if len(oTemp) != len(result) {
			t.Fatal(`FAIL: len(result) != len(expected)`)
		}
		out := make([]int32, len(oTemp))
		for j, item := range oTemp {
			itemTemp, err := strconv.ParseInt(item, 10, 64)
			checkErr(err)
			out[j] = int32(itemTemp)
		}

		for x, y := range result {
			if y != out[x] {
				t.Fatalf(`FAIL: result != output  %v != %v`, result, out)
			}
		}

	}

}

func BenchmarkLeftRot(b *testing.B) {
	items := 100
	d := int32(10)
	a := make([]int32, items)
	for i := 1; i < items; i++ {
		a = append(a, int32(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		leftRot(a, d)
	}
}

func BenchmarkLeftRotSlicing(b *testing.B) {
	items := 100
	d := int32(10)
	a := make([]int32, items)
	for i := 1; i < items; i++ {
		a = append(a, int32(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		leftRotSlicing(a, d)
	}
}
