package redovalnica

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 1 {
		fmt.Println("Ocena", ocena, "je premajhna!")
		return
	}

	if ocena > 10 {
		fmt.Println("Ocena", ocena, "je prevelika!")
		return
	}

	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent z vpisno številko", vpisnaStevilka, "ne obstaja!")
		return
	}

	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}

	if len(student.ocene) < 6 {
		return 0.0
	}

	sum := 0

	for _, ocena := range student.ocene {
		sum += ocena
	}

	return float64(sum) / float64(len(student.ocene))
}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")

	for vpisnaStevilka, student := range studenti {
		fmt.Println(vpisnaStevilka, "-", student.ime, student.priimek+":", student.ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisnaStevilka, student := range studenti {
		povp := povprecje(studenti, vpisnaStevilka)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.ime, student.priimek, povp)

		if povp >= 9 {
			fmt.Println("Odličen študent!")
		} else if povp >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}
