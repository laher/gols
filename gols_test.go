package gols

import (
	"reflect"
	"testing"
)

func TestSplitQuotedString(t *testing.T) {

	var fixtures = map[string][]string{
		"a b c":     []string{"a", "b", "c"},
		"a 'b c'":   []string{"a", "'b c'"},
		"a 'b c' d": []string{"a", "'b c'", "d"},
		"a 'b c":    []string{"a 'b", "c"},
		`"a b c"`:   []string{`"a b c"`},
	}

	for input, expected := range fixtures {
		actual := splitQuotedString(input)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("splitQuotedString expected: [%#v] vs actual [%#v]", expected, actual)
		}
	}

}
