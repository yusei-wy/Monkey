package evaluator

import (
	"testing"

	"monkey/object"
)

func Test_Quote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`quote(5)`,
			`5`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		quote, ok := evaluated.(*object.Quote)
		if !ok {
			t.Errorf("expected *object.Quote. got=%T (%+v)", evaluated, evaluated)
		}

		if quote.Node == nil {
			t.Fatalf("quote.Node is nil")
		}

		if quote.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), tt.expected)
		}
	}
}
