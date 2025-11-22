// Package redovalnica provides a simple gradebook system for managing students,
// storing their grades, enforcing grade boundaries, and computing final results
// based on configurable grading rules.
//
// Example usage:
//
// r := redovalnica.UstvariRedovalnico(3, 1, 10) // need 3 grades, valid range 1..10
// r.DodajStudenta("1001", redovalnica.Student{Ime: "Ana", Priimek: "Horvat"})
// r.DodajOceno("1001", 9)
// r.DodajOceno("1001", 8)
// r.DodajOceno("1001", 10)
// fmt.Println("=== Izpis redovalnice ===")
// r.IzpisVsehOcen()
// fmt.Println("\n=== Končni uspeh ===")
// r.IzpisiKoncniUspeh()
package redovalnica

import "fmt"

// Redovalnica represents a gradebook with specific grading rules and a collection of students.
type Redovalnica struct {
	stOcen   int
	minOcena int
	maxOcena int
	studenti map[string]Student
}

// Student represents a student with a name, surname, and a list of grades.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// UstvariRedovalnico creates a new Redovalnica with specified grading rules.
func UstvariRedovalnico(stOcen int, minOcena int, maxOcena int) *Redovalnica {
	return &Redovalnica{stOcen, minOcena, maxOcena, make(map[string]Student)}
}

// DodajStudenta adds a new student to the gradebook.
func (r *Redovalnica) DodajStudenta(vpisnaStevilka string, student Student) {
	r.studenti[vpisnaStevilka] = student
}

// DodajOceno adds a grade to a student's record, enforcing grade boundaries.
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

// povprecje calculates the average grade for a student if they have enough grades.
// If the student doesn't have enough grades, it returns 0.0. If the student does not exist, it returns -1.0.
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

// IzpisVsehOcen prints the list of students along with their grades.
func (r *Redovalnica) IzpisVsehOcen() {
	fmt.Println("REDOVALNICA:")

	for vpisnaStevilka, student := range r.studenti {
		fmt.Println(vpisnaStevilka, "-", student.Ime, student.Priimek+":", student.Ocene)
	}
}

// IzpisiKoncniUspeh prints the final success evaluation for each student based on their average grade.
func (r *Redovalnica) IzpisiKoncniUspeh() {
	for vpisnaStevilka, student := range r.studenti {
		povp := r.povprecje(vpisnaStevilka)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.Ime, student.Priimek, povp)

		razponOcen := r.maxOcena - r.minOcena + 1

		if povp >= 0.9*float64(razponOcen) {
			fmt.Println("Odličen študent!")
		} else if povp >= 0.6*float64(razponOcen) {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}
