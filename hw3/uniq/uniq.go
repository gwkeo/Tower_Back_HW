package uniq

import (
	"fmt"
	"strings"
)

type StringCountPair struct {
	str   string
	count int
}

type Attributes struct {
	ExportPath            string
	CountSameLines        bool // -c
	ReturnOnlySameLines   bool // -d
	ReturnOnlyUniqueLines bool // -u
	NumberOfFieldsToSkip  int  // -f num_fields
	NumberOfCharsToSkip   int  // -s num_chars
	IgnoreCase            bool // -i
}

func SkipFieldsOfLine(line string, numFields int) string {
	result := strings.Split(line, " ")
	if numFields > len(result) {
		return ""
	}
	result = result[numFields:]
	return strings.Join(result, " ")
}

func SkipCharsOfLine(line string, numChars int) string {
	if numChars > len(line) {
		return ""
	}
	return line[numChars:]
}

func Uniq(content []string, attributes *Attributes) (string, error) {

	type StrCountPair struct {
		Count int
		Str   string
	}

	var strCountPairs []StrCountPair
	var index int
	k := 1
	for i := 0; i < len(content)-1; i++ {
		currentLine := content[i]
		nextLine := content[i+1]

		if attributes.NumberOfCharsToSkip > 0 &&
			attributes.NumberOfCharsToSkip < len(currentLine) &&
			attributes.NumberOfCharsToSkip < len(nextLine) {
			currentLine = currentLine[attributes.NumberOfCharsToSkip:]
			nextLine = currentLine[attributes.NumberOfCharsToSkip:]
		}

		if currentLine == nextLine {
			if k == 1 {
				index = i
			}
			k++
		} else {
			strCountPairs = append(strCountPairs, StrCountPair{
				Count: k,
				Str:   content[index],
			})
			k = 1
		}
	}

	result := ""
	for _, str := range strCountPairs {
		result += fmt.Sprintln(str.Count, str.Str)
	}

	return result, nil
}

/*

attr = 0 =>

*/
