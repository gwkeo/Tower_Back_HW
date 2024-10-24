package uniq

import "testing"

var (
	threeSameLines      = []string{"cake", "cake", "cake"}
	sameAndUniqueLines  = []string{"cake", "cake", "bake"}
	differentFirstChar  = []string{"cake", "bake", "take"}
	differentFirstField = []string{"I met", "We met", "They met"}
	differentCase       = []string{"cAke", "cakE", "Cake"}

	defaultFlags   = Attributes{}
	skipFieldFlag  = Attributes{NumberOfFieldsToSkip: 1}
	skipCharFlag   = Attributes{NumberOfCharsToSkip: 1}
	ignoreCaseFlag = Attributes{IgnoreCase: true}
	countSameFlag  = Attributes{CountSameLines: true}
	uniqueOnlyFlag = Attributes{ReturnOnlyUniqueLines: true}
	sameOnlyFlag   = Attributes{ReturnOnlySameLines: true}
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name       string
		content    []string
		attributes Attributes
		want       string
	}{
		{"Same 3 lines with no flags", threeSameLines, defaultFlags, "cake\n"},
		{"Different 3 lines with no flags", differentFirstChar, defaultFlags, "cake\nbake\ntake\n"},
		{"3 lines with different first chars with differentFirstChar flag", differentFirstChar, skipCharFlag, "cake\n"},
		{"3 lines with different first fields with differentFirstField flag", differentFirstField, skipFieldFlag, "I met\n"},
		{"3 lines with random case with ignoreCase flag", differentCase, ignoreCaseFlag, "cAke\n"},
		{"3 lines with two same and one unique with countSame flag", sameAndUniqueLines, countSameFlag, "2 cake\n1 bake\n"},
		{"3 lines with two same and one unique with uniqueOnly flag", sameAndUniqueLines, uniqueOnlyFlag, "cake\n"},
		{"3 lines with two same and one unique with sameOnly flag", sameAndUniqueLines, sameOnlyFlag, "bake\n"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result, ok := Uniq(test.content, &test.attributes); ok == nil && result != test.want {
				t.Errorf("Uniq() got: %v, want: %v", result, test.want)
			}
		})
	}
}
