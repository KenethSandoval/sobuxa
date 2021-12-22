package main

import (
	"github.com/KenethSandoval/sobuxa/cmd"
)

type data struct {
	nameFile    string
	description string
	idcobertura string
}

func main() {
	//cmd.InformationData()
	d := data{"inserCoberturas", "SINIESTROS ANTERIORES AL INICIO DE VIGENCIA.", "532"}
	cmd.GenerateNewFileSQL(d.nameFile, d.description, d.idcobertura)
}
