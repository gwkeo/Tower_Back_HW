package uniq

import (
	"strconv"
	"strings"
)

type StrCount struct {
	count       int
	str         string
	modifiedStr string
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

func SkipFields(line *string, num int) {
	fields := strings.Split(*line, " ")
	if len(fields) < num {
		*line = ""
	} else {
		fields = fields[num:]
		*line = strings.Join(fields, " ")
	}
}

func SkipChars(line *string, num int) {
	if len(*line) < num {
		*line = ""
	} else {
		*line = (*line)[num:]
	}
}

func ModifyLine(line string, attributes *Attributes) string {
	if attributes.NumberOfFieldsToSkip > 0 {
		SkipFields(&line, attributes.NumberOfFieldsToSkip)
	}
	if attributes.NumberOfCharsToSkip > 0 {
		SkipChars(&line, attributes.NumberOfCharsToSkip)
	}
	if attributes.IgnoreCase {
		line = strings.ToLower(line)
	}
	return line
}

func Uniq(content []string, attributes *Attributes) (string, error) {

	var strCount []StrCount
	strCount = append(strCount, StrCount{
		1,
		content[0],
		ModifyLine(content[0], attributes),
	})

	for i, j := 1, 0; i < len(content); i++ {
		line := ModifyLine(content[i], attributes)
		if strCount[j].modifiedStr == line {
			strCount[j].count++
		} else {
			strCount = append(strCount, StrCount{
				1,
				content[i],
				line,
			})
			j++
		}
	}

	result := ""
	for _, elem := range strCount {
		if attributes.ReturnOnlySameLines && elem.count > 1 {
			result += elem.str + "\n"
		} else if attributes.ReturnOnlyUniqueLines && elem.count == 1 {
			result += elem.str + "\n"
		} else if attributes.CountSameLines {
			result += strconv.Itoa(elem.count) + " " + elem.str + "\n"
		} else {
			result += elem.str + "\n"
		}
	}

	return result, nil
}
