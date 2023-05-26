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
	comentario := "/**/" //Recupera los comentarios del fichero de configuracion cfg
	alm_num := 0         //Numero de alarmas creadas
	lignes := 0          //Num de ligne

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
			escribirEntradas(salidaFile, comentario, valor)
			alm_num++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo de entrada:", err)
		return
	}

	fmt.Println("Archivo de salida creado exitosamente.")

}

// Función para escribir las entradas proporcionales al valor de ALM_VALUE en el alnrchivo de salida
func escribirEntradas(salidaFile *os.File, comentario string, valor int) {
	fmt.Fprintln(salidaFile, comentario)
	fmt.Fprintln(salidaFile, "ALM_NUM:", valor)
	fmt.Fprintln(salidaFile, "ALM_OUT:", "0xF")
	fmt.Fprintln(salidaFile, "ALM_MASK:", "0")
	fmt.Fprintln(salidaFile, "ALM_CHG:", "0")
	fmt.Fprintln(salidaFile, "ALM_STATUS:", "0x11")
	fmt.Fprintln(salidaFile, "ALM_SCNEW:\tbegin")
	fmt.Fprintln(salidaFile, "\tinfo")
	fmt.Fprintln(salidaFile, "end")
	fmt.Fprintln(salidaFile, "ALM_SCOLD:\tbegin")
	fmt.Fprintln(salidaFile, "\tinfo")
	fmt.Fprintln(salidaFile, "end")
	fmt.Fprintln(salidaFile, "")
}
