package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/nishanths/go-hgconfig"
	"github.com/tcnksm/go-gitconfig"
)

const (
	nameEnv       = "LICENSE_FULL_NAME"
	versionString = "v5"
)

var (
	stdout = log.New(os.Stdout, "", 0)
	stderr = log.New(os.Stderr, "", 0)
)

var (
	fName    string = getName()
	fYear    string = strconv.Itoa(time.Now().Year())
	fOutput  string
	fVersion bool
	fHelp    bool
	fList    bool
)

func main() {
	flag.StringVar(&fName, "name", fName, "name on license")
	flag.StringVar(&fName, "n", fName, "name on license")
	flag.StringVar(&fYear, "year", fYear, "year on license")
	flag.StringVar(&fYear, "y", fYear, "year on license")
	flag.StringVar(&fOutput, "output", "", "path to output file")
	flag.StringVar(&fOutput, "o", "", "path to output file")
	flag.BoolVar(&fVersion, "version", false, "print version")
	flag.BoolVar(&fVersion, "v", false, "print version")
	flag.BoolVar(&fHelp, "help", false, "print help")
	flag.BoolVar(&fList, "list", false, "print available licenses")

	flag.Parse()

	run()
}

func run() {
	if flag.NArg() != 1 && !(fVersion || fHelp || fList) {
		userInputError.Abort(errors.New("Unsuitable number of arguments with respect to use version, help, or list switches"))
	}

	switch {
	case fVersion:
		printVersion()
		noError.Abort(nil)

	case fHelp:
		flag.Usage()
		noError.Abort(nil)

	case fList:
		printList()
		noError.Abort(nil)

	default:
		license := strings.ToLower(flag.Arg(0))
		printLicense(license, fOutput, fName, fYear) // internally calls os.Exit() on failure
	}
}

func printVersion() {
	stdout.Printf("%s", versionString)
}

func getName() string {
	n := os.Getenv(nameEnv)
	if n != "" {
		return n
	}
	n, err := gitconfig.Username()
	if err == nil {
		return n
	}
	n, err = gitconfig.Global("user.name")
	if err == nil {
		return n
	}
	n, err = hgconfig.Username()
	if err == nil {
		return n
	}
	usr, err := user.Current()
	if err == nil {
		return usr.Name
	}
	return ""
}
