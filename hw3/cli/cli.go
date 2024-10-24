package cli

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gwkeo/Tower_Back_HW/hw3/uniq"
	"os"
	"strings"
)

func GetInput() ([]string, error) {
	var result []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetFileContent(path string) ([]string, error) {
	result, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(result), "\n"), nil
}

func GetContent(args []string) ([]string, error) {
	var content []string
	var err error

	if len(args) < 1 {
		content, err = GetInput()
		if err != nil {
			return nil, err
		}
	} else {
		content, err = GetFileContent(args[0])
		if err != nil {
			return nil, err
		}
	}
	return content, nil
}

func GetPath(args []string) string {
	if len(args) < 2 {
		return ""
	} else {
		return args[1]
	}
}

func ParseFlags() ([]string, *uniq.Attributes, error) {
	countSameLines := flag.Bool("c", false, "вывести число повторений строки в ее начале")
	returnOnlySameLines := flag.Bool("d", false, "вывести только повторяющиеся строки")
	returnOnlyUniqueLines := flag.Bool("u", false, "вывести только уникальные строки")
	numberOfFieldsToSkip := flag.Int("f", 0, "не учитывать n полей в строке")
	numberOfCharsToSkip := flag.Int("s", 0, "не учитывать n символов")
	ignoreCase := flag.Bool("i", false, "не учитывать регистр строк")

	flag.Parse()

	args := flag.Args()

	content, contentErr := GetContent(args)
	if contentErr != nil {
		return nil, nil, contentErr
	}

	exportPath := GetPath(args)

	attributes := &uniq.Attributes{
		ExportPath:            exportPath,
		CountSameLines:        *countSameLines,
		ReturnOnlySameLines:   *returnOnlySameLines,
		ReturnOnlyUniqueLines: *returnOnlyUniqueLines,
		NumberOfFieldsToSkip:  *numberOfFieldsToSkip,
		NumberOfCharsToSkip:   *numberOfCharsToSkip,
		IgnoreCase:            *ignoreCase,
	}

	return content, attributes, nil
}

func CheckFlagConflicts(attributes *uniq.Attributes) error {
	if attributes.CountSameLines && (attributes.ReturnOnlySameLines || attributes.ReturnOnlyUniqueLines) {
		err := fmt.Errorf("можно передать лишь один из трех параметров: [-c | -d | -u]")
		flag.Usage()
		return err
	}
	return nil
}

func GetAttributes() ([]string, *uniq.Attributes, error) {
	content, attributes, err := ParseFlags()
	if err != nil {
		return nil, nil, err
	}

	if flagConflictsErr := CheckFlagConflicts(attributes); flagConflictsErr != nil {
		return nil, nil, flagConflictsErr
	}

	return content, attributes, nil
}

func WriteToFile(content string, path string) error {
	ok := os.WriteFile(path, []byte(content), 0644)
	if ok != nil {
		return fmt.Errorf("unable to write to file: %s", ok)
	}
	return nil
}
