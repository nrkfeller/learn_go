package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/sunfmin/fanout"
)

func main() {

	var workers = runtime.NumCPU()

	bigMat := extractDirectory("33x33 Samples")

	inputs := []interface{}{}
	for _, img := range bigMat {
		inputs = append(inputs, img)
	}

	fmt.Printf("Running program of %d cores \n", workers)

	results, err2 := fanout.ParallelRun(workers, func(input interface{}) (interface{}, error) {
		img := input.([]byte)
		outr := domainAvailable(img)
		return outr, nil
	}, inputs)

	fmt.Println("Finished ", results, ", Error:", err2)

}
func domainAvailable(word []byte) (ch int) {
	for _, v := range word {
		ch += int(v)
	}
	return
}

func extractDirectory(dirName string) [][]byte {
	var mat [][]byte
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		img := readInputFile(f.Name(), dirName)
		mat = append(mat, img)
	}
	return mat
}

func readInputFile(name string, ReadDir string) []byte {
	var imgMat []byte
	inFile, _ := os.Open(ReadDir + "/" + name)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		stringSlice := strings.Split(line, " ")

		si, _ := sliceAtoi(stringSlice)
		if len(si) > 10 {
			imgMat = si
		}
	}
	return imgMat
}

func sliceAtoi(sa []string) ([]byte, error) {
	si := make([]byte, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, byte(i))
	}
	return si, nil
}
