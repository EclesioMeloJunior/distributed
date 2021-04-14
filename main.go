package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/EclesioMeloJunior/distributed/storage"
)

var (
	datadirFlag string
)

func init() {
	flag.StringVar(&datadirFlag, "datadir", getHomeDir(), "--datdir will setup where in the user filesystem the database will setup")
}

func main() {
	flag.Parse()

	err := storage.CreateDataDirIfNotExists(datadirFlag)

	if err != nil {
		log.Fatal(err)
	}

	s, err := storage.NewStorage(datadirFlag)

	if err != nil {
		log.Fatal(err)
	}

	err = s.StoreFile([]byte("1"), []byte("eclesio"))

	if err != nil {
		log.Fatal(err)
	}

	data, err := s.GetFile([]byte("1"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func getHomeDir() string {
	dir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("could not find homedir, you must pass a datadir")
	}

	return dir
}
