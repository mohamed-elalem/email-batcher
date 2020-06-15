package main

import (
	"flag"
	"fmt"

	"github.com/mohamed-elalem/email-batcher/cli"
)

type varFlag []string

func (vf *varFlag) String() string {
	return fmt.Sprintf("%+v", []string(*vf))
}

func (vf *varFlag) Set(value string) error {
	*vf = append(*vf, value)
	return nil
}

var templatePath = flag.String("email-template-path", "./email.txt", "The email template path")
var recipientsPath = flag.String("recipients-path", "./recipients.txt", "The recipients list path")
var extraVariables varFlag

func main() {
	flag.Var(&extraVariables, "e", "Add new custom variable")
	flag.Parse()

	cli.Run(*templatePath, *recipientsPath, extraVariables)
}
