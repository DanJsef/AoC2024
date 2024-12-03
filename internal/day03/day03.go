package day03

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func checkNextRune(reader *bufio.Reader, expectedRune rune) bool {
	input, _, err := reader.ReadRune()
	if err != nil {
		return false
	}
	return input == expectedRune
}

func checkNextRuneNumeric(reader *bufio.Reader) (rune, bool) {
	input, _, err := reader.ReadRune()
	if err != nil {
		return 0, false
	}

	return input, unicode.IsDigit(input)
}

func isInstructionSwitch(reader *bufio.Reader) (bool, bool) {
	if !checkNextRune(reader, 'o') {
		reader.UnreadRune()
		return false, false
	}

	if !checkNextRune(reader, 'n') {
		reader.UnreadRune()

		if !checkNextRune(reader, '(') {
			reader.UnreadRune()
			return false, false
		}

		if !checkNextRune(reader, ')') {
			reader.UnreadRune()
			return false, false
		}

		return true, true
	}

	if !checkNextRune(reader, '\'') {
		reader.UnreadRune()
		return false, false
	}

	if !checkNextRune(reader, 't') {
		reader.UnreadRune()
		return false, false
	}

	if !checkNextRune(reader, '(') {
		reader.UnreadRune()
		return false, false
	}

	if !checkNextRune(reader, ')') {
		reader.UnreadRune()
		return false, false
	}
	return true, false
}

func Run() {
	file, err := os.Open("./inputs/day03.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	enabled := true

	puzzleResult := 0
  secondPuzzleResult := 0

	for input, _, err := reader.ReadRune(); err == nil; input, _, err = reader.ReadRune() {

		if input != 'm' {
			if input == 'd' {
				if isInstruction, switchTo := isInstructionSwitch(reader); isInstruction {
					enabled = switchTo
				}
			}

			continue
		}

		if !checkNextRune(reader, 'u') {
			reader.UnreadRune()
			continue
		}
		if !checkNextRune(reader, 'l') {
			reader.UnreadRune()
			continue
		}
		if !checkNextRune(reader, '(') {
			reader.UnreadRune()
			continue
		}

		firstArgSlice := []rune{}

		inputRune, isNumeric := checkNextRuneNumeric(reader)

		if !isNumeric {
			reader.UnreadRune()
			continue
		}

		for isNumeric {
			firstArgSlice = append(firstArgSlice, inputRune)
			inputRune, isNumeric = checkNextRuneNumeric(reader)
		}

		if inputRune != ',' {
			reader.UnreadRune()
			continue
		}

		secondArgSlice := []rune{}

		inputRune, isNumeric = checkNextRuneNumeric(reader)

		if !isNumeric {
			reader.UnreadRune()
			continue
		}

		for isNumeric {
			secondArgSlice = append(secondArgSlice, inputRune)
			inputRune, isNumeric = checkNextRuneNumeric(reader)
		}

		if inputRune != ')' {
			reader.UnreadRune()
			continue
		}

		firstArg, _ := strconv.Atoi(string(firstArgSlice))
		secondArg, _ := strconv.Atoi(string(secondArgSlice))

		puzzleResult += firstArg * secondArg
		if enabled {
			secondPuzzleResult += firstArg * secondArg
		}
	}

	fmt.Println("First puzzle result:", puzzleResult)
	fmt.Println("Second puzzle result:", secondPuzzleResult)
}
