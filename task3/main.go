package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Sortirovka(input []string, k int, n bool, r bool, u bool) []string {
	masStr := make([][]string, len(input))
	mp := make(map[string]int)
	slice := make([]string, 0, len(input))
	mpCompare := make(map[string]int)
	output := make([]string, 0, len(input))
	for i, val := range input {
		if u {
			if mpCompare[val] > 0 {
				continue
			}
			mpCompare[val]++
		}
		masStr[i] = strings.Split(val, " ")
		if len(masStr[i]) < k+1 {
			continue
		}
		temp := masStr[i][k]
		mp[temp] = i
		slice = append(slice, temp)
	}
	if r {
		sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	} else {
		sort.Strings(slice)
	}
	for _, val := range slice {
		output = append(output, input[mp[val]])
	}
	return output
}

func ReadFile(name string) ([]string, error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	output := make([]string, 0)
	sc := bufio.NewScanner(file)
	i := 0
	for sc.Scan() {
		output = append(output, sc.Text())
		i++
	}
	return output, nil
}

func WriteFile(input []string, name string) error {
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		return err
	}
	for _, val := range input {
		file.WriteString(val + "\n")
	}
	return nil
}

func main() {
	flagK := flag.Int("k", 0, "Start column")
	flagN := flag.Bool("n", false, "Start column")
	flagR := flag.Bool("r", false, "Start column")
	flagU := flag.Bool("u", false, "Start column")
	flag.Parse()
	input, err := ReadFile("TextFile")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	result := Sortirovka(input, *flagK, *flagN, *flagR, *flagU)
	err = WriteFile(result, "OutTextFile")
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
