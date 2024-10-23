package uniq

import (
	"fmt"
)

type Attributes struct {
	Content               []string
	ExportPath            string
	CountSameLines        bool // -c
	ReturnOnlySameLines   bool // -d
	ReturnOnlyUniqueLines bool // -u
	NumberOfFieldsToSkip  int  // -f num_fields
	NumberOfCharsToSkip   int  // -s num_chars
	IgnoreCase            bool // -i
}

type AttributesFunc func(attributes *Attributes)

func PasteContent(content []string) AttributesFunc {
	return func(attributes *Attributes) {
		attributes.Content = content
	}
}

func WithExportPath(exportPath string) AttributesFunc {
	return func(attributes *Attributes) {
		attributes.ExportPath = exportPath
	}
}

func WithCountSameLines() AttributesFunc {
	return func(attributes *Attributes) {
		attributes.CountSameLines = true
	}
}

func WithReturnOnlySameLines() AttributesFunc {
	return func(attributes *Attributes) {
		attributes.ReturnOnlySameLines = true
	}
}

func WithReturnOnlyUniqueLines() AttributesFunc {
	return func(attributes *Attributes) {
		attributes.ReturnOnlyUniqueLines = true
	}
}

func WithNumberOfFieldsToSkip(numberOfFieldToSkip int) AttributesFunc {
	return func(attributes *Attributes) {
		attributes.NumberOfFieldsToSkip = numberOfFieldToSkip
	}
}

func WithNumberOfCharsToSkip(numberOfFieldToSkip int) AttributesFunc {
	return func(attributes *Attributes) {
		attributes.NumberOfCharsToSkip = numberOfFieldToSkip
	}
}

func WithIgnoreCase() AttributesFunc {
	return func(attributes *Attributes) {
		attributes.IgnoreCase = true
	}
}

func Uniq(attributesFunc ...AttributesFunc) (string, error) {
	attributes := &Attributes{}
	for _, fn := range attributesFunc {
		fn(attributes)
	}

	result, ok := GetUniq(attributes)
	if ok != nil {
		return "", fmt.Errorf("GetUniq failed with %v", ok)
	}
	return result, nil
}

func GetUniq(attributes *Attributes) (string, error) {
	return "", nil
}
