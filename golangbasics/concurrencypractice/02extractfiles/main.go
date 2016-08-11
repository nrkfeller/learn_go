package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bigMat := extractDirectory("8x8 Samples")
	fmt.Println(bigMat)
}

func extractDirectory(dirName string) [][]int {
	var mat [][]int
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		img := readInputFile(f.Name(), dirName)
		fmt.Println(img)
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
