package cli

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gwkeo/hw3/uniq"
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

func GetContent(args []string) ([]string, string, error) {
	var content []string
	var exportPath string
	var err error
	switch len(args) {
	case 0:
		content, err = GetInput()
		if err != nil {
			return nil, "", err
		}
	case 1:
		content, err = GetFileContent(args[0])
		if err != nil {
			return nil, "", err
		}
	case 2:
		content, err = GetFileContent(args[0])
		if err != nil {
			return nil, "", err
		}
		exportPath = args[1]
	}
	return content, exportPath, nil
}

func GetAttributes() ([]uniq.AttributesFunc, error) {
	countSameLines := flag.Bool("c", false, "вывести число повторений строки в ее начале")
	returnOnlySameLines := flag.Bool("d", false, "вывести только повторяющиеся строки")
	returnOnlyUniqueLines := flag.Bool("u", false, "вывести только уникальные строки")
	numberOfFieldsToSkip := flag.Int("f", 1, "не учитывать n полей в строке")
	numberOfCharsToSkip := flag.Int("s", 1, "не учитывать n строк")
	ignoreCase := flag.Bool("i", false, "не учитывать регистр строк")

	args := flag.Args()

	flag.Parse()

	if *countSameLines && (*returnOnlySameLines || *returnOnlyUniqueLines) {
		err := fmt.Errorf("можно передать лишь один из трех параметров: [-c | -d | -u]")
		flag.Usage()
		return nil, err
	}

	content, exportPath, err := GetContent(args)
	if err != nil {
		return nil, err
	}

	type optionConditions struct {
		isFlagUsed bool
		option     uniq.AttributesFunc
	}

	optionsList := []optionConditions{
		{true, uniq.PasteContent(content)},
		{exportPath == "", uniq.WithExportPath(exportPath)},
		{*countSameLines, uniq.WithCountSameLines()},
		{*returnOnlySameLines, uniq.WithReturnOnlySameLines()},
		{*returnOnlyUniqueLines, uniq.WithReturnOnlyUniqueLines()},
		{*numberOfFieldsToSkip > 0, uniq.WithNumberOfFieldsToSkip(*numberOfFieldsToSkip)},
		{*numberOfCharsToSkip > 0, uniq.WithNumberOfCharsToSkip(*numberOfCharsToSkip)},
		{*ignoreCase, uniq.WithIgnoreCase()},
	}

	var options []uniq.AttributesFunc

	for _, option := range optionsList {
		if option.isFlagUsed {
			options = append(options, option.option)
		}
	}

	return options, nil
}

func Run() error {
	options, err := GetAttributes()
	if err != nil {
		return err
	}

	uniqResult, uniqErr := uniq.Uniq(options...)

	if uniqErr != nil {
		return uniqErr
	}

	fmt.Println(uniqResult)
	return nil
}
