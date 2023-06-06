/*Créer un fichier Act à partir d'un fichier xsia A.Villanueva*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error
	alms := 0 //Nombre d'alarmes créées
	cmds := 1 //Nombre de cmd créés

	space := 10 //Espace entre les alarmes et les commandes

	comentario := "/**/" //Récupère les commentaires du fichier de configuration cfg

	lignes := 0 //Num de ligne

	//Analysez s'il y a des arguments
	if len(os.Args) == 1 {
		fmt.Println("Erreur args \n Example : makeAct nom_fichier.cfg ")
		os.Exit(0)
	}

	// Ouvrir le premier fichier en lecture
	//entradaFile, err := os.Open("xsia.cfg")

	entradaFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier d'entrée : ", err)
		return
	}
	defer entradaFile.Close()

	// Créer le deuxième fichier pour l'écriture

	//supprimer la partie après le point
	parts := strings.Split(os.Args[1], ".")
	result := strings.TrimSpace(parts[0])

	salidaFile, err := os.Create("act_" + result + ".txt") //fichier avec act_
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier de sortie :", err)
		return
	}
	defer salidaFile.Close()

	// Lire le contenu du fichier d'entrée ligne par ligne
	scanner := bufio.NewScanner(entradaFile)

	//writeNumAlms(salidaFile, alms) //Write Num Alms Virtuelle

	for scanner.Scan() { //Analyser le fichier d'entrée pour écrire chaque ALM
		linea := scanner.Text()
		lignes++ //numéro de la ligne actuelle

		if strings.Contains(linea, "/*") { /* copier le commentaire */
			comentario = linea
		}

		if strings.Contains(linea, "NBR_ALM:") { //Copier  NBR_ALM:
			fmt.Fprintln(salidaFile, linea)
			//writeNumAlms(salidaFile, alms)
		}

		// Vérifiez si la ligne contient l'entrée "ALM_VALUE"
		if strings.Contains(linea, "ALM_VALUE:") {

			// Obtenir la valeur de ALM_VALUE sous la forme d'un entier INT
			valorStr := strings.TrimSpace(strings.TrimPrefix(linea, "ALM_VALUE:"))
			valor, err := strconv.Atoi(valorStr)
			if err != nil {
				fmt.Println("Erreur num ligne :", lignes)
				//fmt.Println("Erreur lors de la conversion de ALM_VALUE en nombre entier:", err)
				return
			}

			// Écrire les entrées proportionnelles à la valeur de ALM_VALUE dans le fichier de sortie
			writeAlm(salidaFile, comentario, valor)
			alms++
		}
	}

	//Write num. cmds
	writeNumCmds(salidaFile, cmds) //Write Num Alms Virtuelle

	//écrire des commandes
	for cmd := space; cmd <= cmds*space; cmd += space {
		writeCmd(salidaFile, "/**/", cmd)
	}

	//écrire la dernière partie act
	writeEnd(salidaFile)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo de entrada:", err)
		return
	}

	fmt.Println("Fichier de sortie créé avec succès.")

}

func writeNumAlms(salidaFile *os.File, valor int) {
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
	fmt.Fprintln(salidaFile, "/* COMMANDES */\n")
	fmt.Fprintln(salidaFile, "NBR_CMD:", valor)
	fmt.Fprintln(salidaFile, "")
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
