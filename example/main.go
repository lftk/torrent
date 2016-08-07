package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/4396/torrent"
)

func dump(m *torrent.MetaInfo) {
	fmt.Println("Announce:", m.Announce)
	fmt.Println("Announce List:")
	for _, val := range m.AnnounceList {
		for _, elem := range val {
			fmt.Println("\t", elem)
		}
	}
	tm := time.Unix(m.CreationDate, 0)
	fmt.Println("Creation Date:", tm)
	fmt.Println("Comment:", m.Comment)
	fmt.Println("Created By:", m.CreatedBy)
	fmt.Println("Encoding:", m.Encoding)
	fmt.Printf("InfoHash: %X\n", m.Hash)
	fmt.Println("Info:")
	fmt.Println("\tPiece Length:", m.Data.PieceLength)
	fmt.Println("\tPrivate:", m.Data.Private)
	fmt.Println("\tName:", m.Data.Name)
	fmt.Println("\tLength:", m.Data.Length)
	fmt.Println("\tMd5sum:", m.Data.MD5)
	fmt.Println("\tFiles:")
	for _, file := range m.Data.Files {
		fmt.Println("\t\tLength:", file.Length)
		fmt.Println("\t\tPath:", strings.Join(file.Path, "/"))
		fmt.Println("\t\tMd5sum:", file.MD5)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please input torrent path")
	}

	for _, path := range os.Args[1:] {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		m, err := torrent.Decode(b)
		if err != nil {
			log.Fatal(err)
		}
		dump(m)
	}
}
