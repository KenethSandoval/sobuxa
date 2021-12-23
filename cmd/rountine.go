package cmd

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

var mode = flag.String("mode", "normal", "Modo de ejecución de la herramienta")

type data struct {
	nameFile    string
	description string
	idcobertura string
}

func InitCMD() {
	flag.Parse()
	switch *mode {
	case "join":
		var linesNew []string
		for i := 657; i <= 687; i++ {
			lines, err := joinFile("querys/inserCoberturas" + strconv.Itoa(i) + ".sql")
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			linesNew = append(linesNew, lines...)

		}
		if err := writeLines(linesNew, "out/foo.out.sql"); err != nil {
			log.Fatalf("writeLines: %s", err)
		}
		break
	case "create":
		var init int = 657
		var d data

		for i := init; i <= (init + 30); i++ {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			line := scanner.Text()

			d = data{
				"inserCoberturas" + strconv.Itoa(i),
				line,
				strconv.Itoa(i),
			}

			generateNewFileSQL(d.nameFile, d.description, d.idcobertura)
		}
		break
	case "clean":
		os.RemoveAll("querys")
		os.RemoveAll("out")
		os.Mkdir("querys", 0755)
		os.Mkdir("out", 0755)
		break
	default:
		log.Fatal("El modo de ejecución no existe")

	}
}
