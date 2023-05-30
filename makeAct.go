package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Analizar si hay argumentos

	alms := 0 //Numero de alarmas creadas
	space := 10

	cmds := 1 //Numero de cmds creados

	comentario := "/**/" //Recupera los comentarios del fichero de configuracion cfg

	lignes := 0 //Num de ligne

	if len(os.Args) == 1 {
		fmt.Println("Erreur args , example : makeAct nom_fichier ")
		os.Exit(0)
	}

	// Abrir el primer archivo para lectura
	entradaFile, err := os.Open("xsia.cfg")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier d'entrée : ", err)
		return
	}
	defer entradaFile.Close()

	// Crear el segundo archivo para escritura
	salidaFile, err := os.Create("act_xsia.txt")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier de sortie :", err)
		return
	}
	defer salidaFile.Close()

	// Leer el contenido del archivo de entrada línea por línea
	scanner := bufio.NewScanner(entradaFile)

	//writeNumAlms(salidaFile, alms) //Write Num Alms Virtuelle

	for scanner.Scan() {
		linea := scanner.Text()
		lignes++ //numéro de la ligne actuelle

		if strings.Contains(linea, "/*") { //Copia comentario /**/
			comentario = linea
		}

		if strings.Contains(linea, "NBR_ALM:") { //Copia  NBR_ALM:
			fmt.Fprintln(salidaFile, linea)
		}

		// Verificar si la línea contiene la entrada "ALM_VALUE"
		if strings.Contains(linea, "ALM_VALUE:") {
			// Obtener el valor de ALM_VALUE como entero
			valorStr := strings.TrimSpace(strings.TrimPrefix(linea, "ALM_VALUE:"))
			valor, err := strconv.Atoi(valorStr)
			if err != nil {
				fmt.Println("Erreur num ligne :", lignes)
				//fmt.Println("Erreur lors de la conversion de ALM_VALUE en nombre entier:", err)
				return
			}

			// Escribir las entradas proporcionales al valor de ALM_VALUE en el archivo de salida
			writeAlm(salidaFile, comentario, valor)
			alms++
		}
	}

	//Write cmd
	writeNumCmds(salidaFile, cmds) //Write Num Alms Virtuelle

	for cmd := space; cmd <= cmds*space; cmd += space {
		writeCmd(salidaFile, "", cmd)

	}

	writeEnd(salidaFile)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo de entrada:", err)
		return
	}

	fmt.Println("Archivo de salida creado exitosamente.")

}

func writeNumAlms(salidaFile *os.File, valor int) {
	fmt.Fprintln(salidaFile, "/* COMMANDES */\n")
	fmt.Fprintln(salidaFile, "NBR_ALM:", valor)
	fmt.Fprintln(salidaFile, "\n")
}

func writeAlm(salidaFile *os.File, comentario string, valor int) {
	fmt.Fprintln(salidaFile, comentario)
	fmt.Fprintln(salidaFile, "ALM_NUM:", valor)
	fmt.Fprintln(salidaFile, "ALM_OUT:", "0xF")
	fmt.Fprintln(salidaFile, "ALM_MASK:", "0")
	fmt.Fprintln(salidaFile, "ALM_CHG:", "0")
	fmt.Fprintln(salidaFile, "ALM_STATUS:", "0x11")
	fmt.Fprintln(salidaFile, "ALM_SCNEW:\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\t\tinfo")
	fmt.Fprintln(salidaFile, "\t\t\tend")
	fmt.Fprintln(salidaFile, "ALM_SCOLD:\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\t\tinfo")
	fmt.Fprintln(salidaFile, "\t\t\tend")
	fmt.Fprintln(salidaFile, "")
}

func writeNumCmds(salidaFile *os.File, valor int) {
	fmt.Fprintln(salidaFile, "/* COMMANDES */")
	fmt.Fprintln(salidaFile, "NBR_CMD:", valor)
}

func writeCmd(salidaFile *os.File, comentario string, valor int) {
	fmt.Fprintln(salidaFile, comentario)
	fmt.Fprintln(salidaFile, "CMD_NUM:", valor)
	fmt.Fprintln(salidaFile, "CMD_OUT:", "0xF")
	fmt.Fprintln(salidaFile, "CMD_STATUS:", "1")
	fmt.Fprintln(salidaFile, "CMD_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\t\tinfo")
	fmt.Fprintln(salidaFile, "\t\t\tend")
	fmt.Fprintln(salidaFile, "")
}

func writeEnd(salidaFile *os.File) {

	fmt.Fprintln(salidaFile, "DATE_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\tend\n")

	fmt.Fprintln(salidaFile, "COUNTER_SC:", "begin")
	fmt.Fprintln(salidaFile, "\t\t\tend\n")

	fmt.Fprintln(salidaFile, "LOGIC_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\tend\n")

	fmt.Fprintln(salidaFile, "CNTADD_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\tend\n")

	fmt.Fprintln(salidaFile, "CNTADDF_SC:", "begin")
	fmt.Fprintln(salidaFile, "\t\t\tend\n")

	fmt.Fprintln(salidaFile, "MASKCMD0:", "\t{")
	fmt.Fprintln(salidaFile, "\t\t\t}\n")

	fmt.Fprintln(salidaFile, "MASKCMD1:", "\t{")
	fmt.Fprintln(salidaFile, "\t\t\t}\n")

	fmt.Fprintln(salidaFile, "MASKCMD2:", "\t{")
	fmt.Fprintln(salidaFile, "\t\t\t}\n")

}
