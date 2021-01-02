package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

const (
	bin            = "go-git"
	receivePackBin = "git-receive-pack"
	uploadPackBin  = "git-upload-pack"
)

func main() {
	if os.Args[1] == "-c" {
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
		writeToLogFile("-c found")
	}
	writeToLogFile(fmt.Sprintf("%+v\n", os.Args))

	splitArgs := strings.Split(os.Args[1], " ")
	splitArgs[1] = splitArgs[1][1 : len(splitArgs[1])-1]
	switch splitArgs[0] {
	case receivePackBin:
		os.Args = append([]string{"git", "receive-pack"}, splitArgs[1])
		writeToLogFile("receivehappening")
	case uploadPackBin:
		os.Args = append([]string{"git", "upload-pack"}, splitArgs[1])
		writeToLogFile("uploadhappening")
	default:
		writeToLogFile("default happening" + os.Args[1])
	}

	writeToLogFile(fmt.Sprintf("NEW ARGS: %+v\n", os.Args))

	theseflags := flags.HelpFlag | flags.PrintErrors | flags.PassDoubleDash | flags.IgnoreUnknown
	parser := flags.NewNamedParser(bin, flags.Options(theseflags))
	parser.AddCommand("receive-pack", "", "", &CmdReceivePack{})
	parser.AddCommand("upload-pack", "", "", &CmdUploadPack{})
	parser.AddCommand("version", "Show the version information.", "", &CmdVersion{})

	_, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrCommandRequired {
			parser.WriteHelp(os.Stdout)
		}

		os.Exit(1)
	}
}

type cmd struct {
	Verbose bool `short:"v" description:"Activates the verbose mode"`
}
