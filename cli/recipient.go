package cli

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Recipient struct {
	Name  string
	Email string
}

type Recipients []Recipient

func LoadRecipients(path string) Recipients {
	file, err := openRecipientsFile(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	recipients := readRecipientsFromFile(file)
	return recipients
}

func openRecipientsFile(path string) (*os.File, error) {
	return os.Open(path)
}

func readRecipientsFromFile(file *os.File) Recipients {
	scanner := bufio.NewScanner(file)

	recipients := make(Recipients, 0)

	for scanner.Scan() {
		line := scanner.Text()

		recipient := readRecipient(line)

		recipients = append(recipients, recipient)
	}

	return recipients
}

func readRecipient(line string) Recipient {
	parts := strings.Split(line, ",")
	return Recipient{Name: strings.TrimSpace(parts[0]), Email: strings.TrimSpace(parts[1])}
}
