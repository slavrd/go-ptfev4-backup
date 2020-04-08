package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/slavrd/go-ptfev4-backup/helpers"
)

var fpass = flag.String("pass", "", "Encryption password for the backup data. Can also be set via TFE_BACKUP_PASSWORD environment variable.")
var fhost = flag.String("host", "", "Hostname of the tfe instance. E.g. tfe.mydomain.com. Can also be set via TFE_HOSTNAME environment variable.")
var ftoken = flag.String("token", "", "TFE backup authorization token. Can also be set via TFE_BACKUP_TOKEN environment variable.")
var ffile = flag.String("file", "", "File to read/write TFE backup. Can also be set via TFE_BACKUP_FILE environment variable.")

func main() {

	flag.Parse()
	var host, token, pass, file, err = validateInput()
	if err != nil {
		fmt.Printf("%v\n", err)
		printUsage()
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "backup":
		log.Printf("saving backup to %q", file)
		var f *os.File
		f, err = os.Create(file)

		if err != nil {
			log.Fatalf("error saving backup: %v", err)
		}

		defer f.Close()

		err = helpers.TfeBackup(host, token, pass, f)
		if err != nil {
			log.Fatalf("error making backup: %v", err)
		}

		log.Println("backup successful!")

	case "restore":
		log.Printf("restoring backup from %q", file)
		var f *os.File
		f, err = os.Open(file)
		if err != nil {
			log.Fatalf("error opening backup: %v", err)
		}
		defer f.Close()
		err = helpers.TfeRestore(host, token, pass, f)
		if err != nil {
			log.Fatalf("error restoring backup: %v", err)
		}
	}
}

func validateInput() (host, token, pass, file string, err error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	if len(flag.Args()) != 1 {
		err = fmt.Errorf("missing operation - backup or restore")
		return
	}

	if flag.Args()[0] != "backup" && flag.Args()[0] != "restore" {
		err = fmt.Errorf("accepted operations are 'backup' or 'restore'")
		return
	}

	var isErr = false

	if *fpass != "" {
		pass = *fpass
	} else if os.Getenv("TFE_BACKUP_PASSWORD") != "" {
		pass = os.Getenv("TFE_BACKUP_PASSWORD")
	} else {
		log.Printf("encryption password was not provided.")
		isErr = true
	}

	if *fhost != "" {
		host = *fhost
	} else if os.Getenv("TFE_HOSTNAME") != "" {
		host = os.Getenv("TFE_HOSTNAME")
	} else {
		log.Printf("TFE hostname was not provided.")
		isErr = true
	}

	if *ftoken != "" {
		token = *ftoken
	} else if os.Getenv("TFE_BACKUP_TOKEN") != "" {
		token = os.Getenv("TFE_BACKUP_TOKEN")
	} else {
		log.Printf("TFE backup authorization token was not provided.")
		isErr = true
	}

	if *ffile != "" {
		file = *ffile
	} else if os.Getenv("TFE_BACKUP_FILE") != "" {
		file = os.Getenv("TFE_BACKUP_FILE")
	} else {
		log.Printf("file to read/write TFE backup not provided.")
		isErr = true
	}

	if isErr {
		err = fmt.Errorf("missing required input")
	}

	return
}

func printUsage() {
	fmt.Printf("usage %s [args] <backup|restore>\n", os.Args[0])
	flag.Usage()
}
