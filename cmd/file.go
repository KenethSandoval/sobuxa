package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func InformationData() {
	var init int
	var finish int
	var desc string

	fmt.Print("idcobertura para inicar: ")
	fmt.Scanf("%d\n", &init)

	fmt.Print("Datos restantes: ")
	fmt.Scanf("%d\n", &finish)

	for i := init; i <= (init + finish); i++ {
		fmt.Print("Ingres la descripcion: ")
		fmt.Scanf("%s", &desc)
		test(strconv.Itoa(i))
		//d := data{"inserCoberturas", "SINIESTROS ANTERIORES AL INICIO DE VIGENCIA.", strconv.Itoa(i)}
		//cmd.GenerateNewFileSQL(d.nameFile, d.description, d.idcobertura)
	}
}

func test(lines string) {
	b := []byte(lines)

	err := ioutil.WriteFile("test.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

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
