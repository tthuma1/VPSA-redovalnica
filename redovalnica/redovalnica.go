package redovalnica

import "fmt"

type Redovalnica struct {
	stOcen   int
	minOcena int
	maxOcena int
	studenti map[string]Student
}

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func UstvariRedovalnico(stOcen int, minOcena int, maxOcena int) *Redovalnica {
	return &Redovalnica{stOcen, minOcena, maxOcena, make(map[string]Student)}
}

func (r *Redovalnica) DodajStudenta(vpisnaStevilka string, student Student) {
	r.studenti[vpisnaStevilka] = student
}

func (r *Redovalnica) DodajOceno(vpisnaStevilka string, ocena int) {
	if ocena < r.minOcena {
		fmt.Println("Ocena", ocena, "je premajhna! Najmanjša možna ocena je", r.minOcena)
		return
	}

	if ocena > r.maxOcena {
		fmt.Println("Ocena", ocena, "je prevelika! Največja možna ocena je", r.maxOcena)
		return
	}

	student, ok := r.studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent z vpisno številko", vpisnaStevilka, "ne obstaja!")
		return
	}

	student.Ocene = append(student.Ocene, ocena)
	r.studenti[vpisnaStevilka] = student
}

func (r *Redovalnica) povprecje(vpisnaStevilka string) float64 {
	student, ok := r.studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}

	if len(student.Ocene) < r.stOcen {
		return 0.0
	}

	sum := 0

	for _, ocena := range student.Ocene {
		sum += ocena
	}

	return float64(sum) / float64(len(student.Ocene))
}

func (r *Redovalnica) IzpisRedovalnice() {
	fmt.Println("REDOVALNICA:")

	for vpisnaStevilka, student := range r.studenti {
		fmt.Println(vpisnaStevilka, "-", student.Ime, student.Priimek+":", student.Ocene)
	}
}

func (r *Redovalnica) IzpisiKoncniUspeh() {
	for vpisnaStevilka, student := range r.studenti {
		povp := r.povprecje(vpisnaStevilka)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.Ime, student.Priimek, povp)

		if povp >= 9 {
			fmt.Println("Odličen študent!")
		} else if povp >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}
