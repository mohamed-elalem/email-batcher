package cli

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	templatePath := "../tmp/recipients.txt"
	recipients := LoadRecipients(templatePath)

	if len(recipients) != 2 {
		t.Errorf("LoadRecipients(%s) got %d recipients expected %d", templatePath, len(recipients), 2)
	}

	expectedRecipients := []Recipient{
		Recipient{Name: "Mohamed Elalem 1", Email: "mohamed.elalem94+1@gmail.com"},
		Recipient{Name: "Mohamed Elalem 2", Email: "mohamed.elalem94+2@gmail.com"},
	}

	for i, recipient := range recipients {
		if !reflect.DeepEqual(recipient, expectedRecipients[i]) {
			t.Errorf("Recipient %d: expected %+v got %+v", i+1, expectedRecipients[i], recipient)
		}
	}
}

func TestReadRecipient(t *testing.T) {
	tests := []struct {
		input    string
		expected Recipient
	}{
		{
			input:    "Mohamed Elalem 1, mohamed.elalem94+1@gmail.com",
			expected: Recipient{Name: "Mohamed Elalem 1", Email: "mohamed.elalem94+1@gmail.com"},
		},
	}

	for _, test := range tests {
		recipient := readRecipient(test.input)
		if !reflect.DeepEqual(recipient, test.expected) {
			t.Errorf("Expected %+v got %+v", test.expected, recipient)
		}
	}
}
