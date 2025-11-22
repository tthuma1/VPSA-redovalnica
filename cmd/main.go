package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "weatherStation",
		Usage: "Weather station displays Temperature, Humidity, and Pressure",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "timeout",
				Usage: "Timeout for reading sensor in seconds",
				Value: 5,
			},
			&cli.IntFlag{
				Name:  "pollTime",
				Usage: "Sensor refresh interval in ms",
				Value: 500,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			timeoutSec := cmd.Int("timeout")
			pollTime := cmd.Int("pollTime")
			return runTUI(pollTime, timeoutSec)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	studenti := make(map[string]Student)

	student1 := Student{"Janez", "Novak", []int{8, 9, 7}}
	studenti["63230111"] = student1

	student2 := Student{"Polona", "Polončič", []int{}}
	studenti["63220222"] = student2

	student3 := Student{"Marcuss", "Favela", []int{7}}
	studenti["63240333"] = student3

	student4 := Student{"Odlični", "Odličnjakovič", []int{10, 9, 10, 10, 9, 9}}
	studenti["63240444"] = student4

	dodajOceno(studenti, "63230111", 6)
	dodajOceno(studenti, "63230111", 9)
	dodajOceno(studenti, "63230111", 10)

	dodajOceno(studenti, "63230111", 11)
	dodajOceno(studenti, "63230111", 0)
	dodajOceno(studenti, "82934", 7)

	fmt.Println("Povprečje študenta 63230111:", povprecje(studenti, "63230111"))
	fmt.Println("Povprečje študenta 63220222:", povprecje(studenti, "63220222"))
	fmt.Println("Povprečje študenta 9304329:", povprecje(studenti, "9304329"))

	fmt.Println()
	izpisRedovalnice(studenti)

	fmt.Println()
	izpisiKoncniUspeh(studenti)
}
