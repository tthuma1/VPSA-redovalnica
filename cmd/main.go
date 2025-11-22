package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tthuma1/VPSA-redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Redovalnica je aplikacija za hranjenje ocen in računanje končnega uspeha.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			run(stOcen, minOcena, maxOcena)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(stOcen int, minOcena int, maxOcena int) {
	r := redovalnica.UstvariRedovalnico(stOcen, minOcena, maxOcena)

	r.DodajStudenta("63230111", redovalnica.Student{Ime: "Janez", Priimek: "Novak", Ocene: []int{8, 9, 7}})

	r.DodajStudenta("63220222", redovalnica.Student{Ime: "Polona", Priimek: "Polončič", Ocene: []int{}})

	r.DodajStudenta("63240333", redovalnica.Student{Ime: "Marcuss", Priimek: "Favela", Ocene: []int{7}})

	r.DodajStudenta("63240444", redovalnica.Student{Ime: "Odlični", Priimek: "Odličnjakovič", Ocene: []int{10, 9, 10, 10, 9, 9}})

	r.DodajOceno("63230111", 6)
	r.DodajOceno("63230111", 9)
	r.DodajOceno("63230111", 10)

	r.DodajOceno("63230111", 11)
	r.DodajOceno("63230111", 0)
	r.DodajOceno("82934", 7)

	// fmt.Println("Povprečje študenta 63230111:", povprecje("63230111"))
	// fmt.Println("Povprečje študenta 63220222:", povprecje("63220222"))
	// fmt.Println("Povprečje študenta 9304329:", povprecje("9304329"))

	fmt.Println()
	r.IzpisVsehOcen()

	fmt.Println()
	r.IzpisiKoncniUspeh()
}
