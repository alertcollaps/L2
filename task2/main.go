package main

import (
	"fmt"
)

//48 and 59 - count of 0..9
/*
func UnBox(input string) string {
	massRune := []rune(input)
	massRune_change := []rune(input)
	mass := strings.Split(input, "/")
	for i := 0; i < len(massRune); i++{
		current := massRune[i]
		if current == '/'{

		}
	}

	return string(massRune)
}
*/
func isAnumber(rn rune) bool {
	if rn > 59 || rn < 48 {
		return false
	}
	return true
}

func unBox(input string) (string, error) {
	massRune := []rune(input)
	massruneChange := make([]rune, 0, 32)

	for i := 0; i < len(massRune); i++ {
		current := massRune[i]
		if current == '\\' {
			if i+1 >= len(massRune) {
				return "", fmt.Errorf("No enouth symbols after \\")
			}
			next := massRune[i+1]
			nextNext := 'a'
			if i+2 < len(massRune) {
				nextNext = massRune[i+2]
			}
			if isAnumber(next) == (next == 92) {
				return "", fmt.Errorf("%v must be number", string(next))
			}
			if isAnumber(nextNext) {
				for i := 0; i < int(nextNext-'0'); i++ {
					//Add number in slice
					massruneChange = append(massruneChange, next)
				}
				i += 2
				continue
			}
			massruneChange = append(massruneChange, next)
			i++
		} else if !isAnumber(current) {
			next := 'a'
			if i+1 < len(massRune) {
				next = massRune[i+1]
			}
			if isAnumber(next) {
				for i := 0; i < int(next-'0'); i++ {
					massruneChange = append(massruneChange, current)
				}
				i++
				continue
			}
			massruneChange = append(massruneChange, current)
		} else {
			return "", fmt.Errorf("Error in collumn %v. Expected no number value", i)
		}
	}
	return string(massruneChange), nil
}
func main() {

	val, err := unBox("\\\\")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Output:", val)
}
