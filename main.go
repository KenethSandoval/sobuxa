package main

import (
	//"bufio"
	//"log"
	//"os"
	//"strconv"

	"log"
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
	/*var init int = 553
	var d data

	for i := init; i <= (init + 4); i++ {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		d = data{
			"inserCoberturas" + strconv.Itoa(i),
			line,
			strconv.Itoa(i),
		}

		log.Default(d)
		cmd.GenerateNewFileSQL(d.nameFile, d.description, d.idcobertura)
	}*/

	var linesNew []string
	for i := 541; i <= 555; i++ {
		lines, err := cmd.JoinFile("./test/inserCoberturas" + strconv.Itoa(i) + ".sql")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		linesNew = append(linesNew, lines...)

	}
	if err := cmd.WriteLines(linesNew, "out/foo.out.sql"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}

}
