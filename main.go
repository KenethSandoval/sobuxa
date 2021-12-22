package main

import (
	"fmt"
	"strconv"

	"github.com/KenethSandoval/sobuxa/cmd"
)

type data struct {
	nameFile    string
	description string
	idcobertura string
}

func main() {
	//cmd.InformationData()
	var init int = 538
	var d data
	var desc string

	for i := 538; i <= (init + 2); i++ {
		d = data{
			"inserCoberturas" + strconv.Itoa(i),
			desc,
			strconv.Itoa(i),
		}
		fmt.Println(d)
		cmd.GenerateNewFileSQL(d.nameFile, d.description, d.idcobertura)
	}
}
