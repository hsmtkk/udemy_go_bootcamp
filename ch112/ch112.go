package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hsmtkk/udemy_go_bootcamp/ch112/feet"
	"github.com/spf13/cobra"
)

func main() {
	var (
		meter bool
		yard  bool
	)
	command := &cobra.Command{
		Use:  "ch112 feet",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			f, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			converter := feet.New()
			if meter {
				meter := converter.FeetMeter(f)
				fmt.Printf("%f feet is %f meter\n", f, meter)
			} else if yard {
				yard := converter.FeetYard(f)
				fmt.Printf("%f feet is %f yard\n", f, yard)
			} else {
				meter := converter.FeetMeter(f)
				fmt.Printf("%f feet is %f meter\n", f, meter)
			}
		},
	}
	command.Flags().BoolVar(&meter, "meter", false, "convert feet to meter")
	command.Flags().BoolVar(&yard, "yard", false, "convert feet to yard")
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
