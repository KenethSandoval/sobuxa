package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//GenerateNewFileSQL genera el sql basado en un scriptsql base
func GenerateNewFileSQL(nameFile string, description string, idcobertura string) {
	s := `INSERT INTO seguros.cobertura 
(idcobertura, tipo, codigo, descripcion, suma, prima, tasa, deducible_mto, 
deducible_porcentaje, deducible_minimo, deducible_base, orden, 
especificacion, indestandar, idaseguradora, indacumulasuma, 
categoria, subGrupo, id_tarifa, deducible_monto, indeliminable, indcobdecl, estado)
VALUES (` + idcobertura + `, 'exclusion', '', '` + description + `', 0.00, 0.00, 0.000000, NULL, 0.000000, 0.00, 'null', 0, '', 'S', 1, 'N', 
'DAÑOS', NULL, 0, 0, 'S', NULL, NULL);

INSERT INTO seguros.ramo_cobertura
(idcobertura, idramo, tipo, codigo, descripcion, suma, prima, tasa, deducible_mto, deducible_porcentaje, 
deducible_minimo, deducible_base, orden, especificacion, indacumulasuma, indtarificar, indcopiar, indcobraprima, 
indimprimir, indeliminable, ind_caratula, ind_imprimir_caratula, agrupador, codplanrea, codramorea, 
indcalculabomberos, indsuscripcion, inddeclaracion, indprimadeposito) 
VALUES (` + idcobertura + `, 4, 'exclusion', NULL, NULL, 0.00, 0.00, 0.000000, 0.00, 0.000000, 0.00, '0', 0, 0, 'N', 'N', 'S', 'N', 'S', 'S', 
NULL, 'NO', 'BIENESINCE', NULL, NULL, 'N', 'S', 'N', 'N');

INSERT INTO seguros.plantilla_ramo_cobertura
(idplantilla, version, idramo, idcobertura, descripcion) VALUES (47, 1, 4, ` + idcobertura + `, NULL);
`

	b := []byte(s)

	err := ioutil.WriteFile(nameFile+".sql", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//JoinFile une los archivos
func JoinFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
