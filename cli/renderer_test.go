package cli

import "testing"

func TestRender(t *testing.T) {
	renderer := NewRenderer("../tmp/email.txt")

	expected := `Dear John Doe,

This is a test email

Best Regards,

Mohamed Elalem

mohamed.elalem94@gmail.com`

	output, err := renderer.render(struct {
		Recipient Recipient
		Extra     map[string]string
	}{
		Recipient: Recipient{Name: "John Doe"},
		Extra:     map[string]string{"name": "Mohamed Elalem", "email": "mohamed.elalem94@gmail.com"},
	})

	if err != nil {
		t.Error(err)
	}

	if output != expected {
		t.Errorf("Expected\n%s\ngot\n%s", expected, output)
	}

}
