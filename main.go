package main

import (
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/micahhausler/utmptail/utmp"
	"io"
	"log"
	"os"
)

const Version = "0.0.1"

var version = flag.Bool("version", false, "print version and exit")

func bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

func main() {
	filename := "/dev/stdin"

	arglen := len(os.Args)
	if arglen == 2 {
		filename = string(os.Args[1])
	} else if arglen > 2 {
		fmt.Fprint(os.Stderr, "Too many arguments!")
		os.Exit(1)
	}

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "utmptail reads in a login record formatted file")
		fmt.Fprint(os.Stderr, "and emits each record as a JSON to\n")
		fmt.Fprint(os.Stderr, "a new line. On most systems these logs are located ")
		fmt.Fprintf(os.Stderr, "at /var/log/{btmp, utmp, wtmp}.\nSee %s ", bold("utmp(5)"))
		fmt.Fprint(os.Stderr, "for more information on login record files.\n\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
		fmt.Fprint(os.Stderr, "  <filename>\n")
		fmt.Fprint(os.Stderr, "    	The file to read in. (default \"/dev/stdin\")\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *version {
		fmt.Printf("utmptail %s\n", Version)
		os.Exit(0)
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		nu := new(utmp.Utmp)
		err := binary.Read(file, binary.LittleEndian, nu)
		if err != nil && err != io.EOF {
			// pass
		}
		if err == io.EOF {
			break
		}
		data, _ := json.Marshal(nu)
		fmt.Println(string(data))
	}
}
