package cli

import (
	"log"
	"strings"
)

type ExtraVariables map[string]string

var extraVariables ExtraVariables

func init() {
	extraVariables = make(ExtraVariables)
}

type RenderingData struct {
	Recipient Recipient
	Extra     ExtraVariables
}

func Run(emailTemplatePath, recipientsPath string, variablesStr []string) {
	for _, variable := range variablesStr {
		parts := strings.Split(variable, "=")
		extraVariables[parts[0]] = parts[1]
	}

	for _, recipient := range LoadRecipients(recipientsPath) {
		renderer := NewRenderer(emailTemplatePath)
		out, err := renderer.render(RenderingData{recipient, extraVariables})
		if err != nil {
			log.Fatal(err)
		}

		log.Println(out)
	}
}
