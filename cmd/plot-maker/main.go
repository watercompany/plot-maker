package main

import (
	"flag"
	"fmt"
	"plot-maker/pkg/jsonarg"
	"time"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Debug parameters")
	jsonFlag := flag.String("json", "", "Input JSON")

	flag.Parse()
	jsonarg.Parse(jsonFlag)

	if *debugFlag {
		for _, a := range jsonarg.PosArgs {
			fmt.Println(a)
		}
	} else {
		startTime := time.Now()
		// cmd := exec.Command("ProofOfSpace", posArgs...)
		// var out bytes.Buffer
		// cmd.Stdout = &out
		// err := cmd.Run()
		// if err != nil {
		// 	panic(err)
		// }
		elapsedTime := time.Since(startTime)
		fmt.Println("Execution took:", elapsedTime)
	}
}
