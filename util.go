package util

/*
 * Modul enthält einige allgemeine Funktionen
 */

import (
	"fmt"
	"log"
)

// allgemeine Fehlerkontrolle
func CheckErr(err error) {
//	fmt.Println("in checkErr", err)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("error %v\n", err)
		// panic(err)
	}
}


// ein Slit mit Namen wird in einen String mit Trennern ", " umgewandelt
func Array2string(in []string) string {
	var s string
	for i, h := range in {
		if i == 0 {
			s += h
		} else {
			s += ", " + h
		}
	}
	return s
}


// Erstellen eines Strings "?, ?"
func QuestionMark(anz int) string {
	var s string
	for i := 0; i < anz; i++ {
		if i == 0 {
			s += "?"
		} else {
			s += ", ?"
		}
	}
	return s
}
