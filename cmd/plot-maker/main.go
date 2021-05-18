package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plot-maker/pkg/jsonarg"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var version = "0.0.1"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	debugFlag := flag.Bool("debug", false, "Debug parameters")
	jsonFlag := flag.String("json", "", "Input JSON")
	versionFlag := flag.Bool("version", false, "Prints versin and exit")
	binaryFlag := flag.String("bin", "ProofOfSpace", "ProofOfSpace binary path")

	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		return
	}

	jsonarg.Parse(jsonFlag)

	if *debugFlag {
		log.Debug().Msg("Arguments:")
		for _, a := range jsonarg.PosArgs {
			log.Debug().Msg(a)
		}
	} else {
		startTime := time.Now()
		cmd := exec.Command(*binaryFlag, jsonarg.PosArgs...)

		cmd.Dir, _ = exec.LookPath(*binaryFlag)
		cmd.Dir = filepath.Dir(cmd.Dir)

		log.Debug().Msg(fmt.Sprintf("Working directory: %s", cmd.Dir))
		log.Debug().Msg(fmt.Sprintf("Command: %q", cmd.String()))
		log.Debug().Msg(fmt.Sprintf("Args: %q", cmd.Args))

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		cmd.Stderr = cmd.Stdout

		log.Debug().Msg(cmd.String())

		err = cmd.Start()
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		log.Info().Msg(fmt.Sprintf("PID: %d", cmd.Process.Pid))

		r := bufio.NewReader(stdout)
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			log.Info().Msg(scanner.Text())
		}
		cmd.Wait()

		elapsedTime := time.Since(startTime)
		log.Info().Msg(fmt.Sprintf("Execution took: %v", elapsedTime))
	}
}
