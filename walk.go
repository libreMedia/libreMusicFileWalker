package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
)

func doThing(fff string) {
	file, err := os.Open(fff) // For read access.
	if err != nil {
		// log.Fatal(err)
		fmt.Println("no artfdsfsdist")
	}

	if file != nil {

		m, err := tag.ReadFrom(file)
		if err != nil {
			// log.Fatal(err)
			fmt.Println("no artist")
		}

		if m == nil {
			fmt.Println("no data")
		} else {

			if m.Title() != "" {
				log.Print(m.Title()) // The title of the track (see Metadata interface for more details).
			} else {
				fmt.Println("no title")
			}
			if m.Artist() != "" {
				log.Print(m.Artist()) // The artist of the track (see Metadata interface for more details).
			} else {
				fmt.Println("no artist")
			}
			if m.Album() != "" {
				log.Print(m.Album()) // The album of the track (see Metadata interface for more details).
			} else {
				fmt.Println("no album")
			}
			if m.Year() != 0 {
				log.Print(m.Year()) // The year of the track (see Metadata interface for more details).
			} else {
				fmt.Println("no year")
			}
			// if m.Track() != 0 {
			// 	log.Print(m.Track()) // The track number of the track (see Metadata interface for more details).
			// } else {
			// 	fmt.Println("no track")
			// }
			if m.Genre() != "" {
				log.Print(m.Genre()) // The genre of the track (see Metadata interface for more details).
			} else {
				fmt.Println("no genre")
			}

			// log.Print(m.Format())      // The detected format.
			// log.Print(m.AlbumArtist()) // The title of the track (see Metadata interface for more details).
			// log.Print(m.Composer())    // The title of the track (see Metadata interface for more details).

		}
	} else {
		println("no file")
	}

}

func walk(dirToWalk string, dbName string) {
	err := filepath.Walk(getParentDir()+dirToWalk,
		func(path string, info os.FileInfo, err error) error {
			// filename := filepath.Join(path, info.Name())
			if err != nil {
				return err
			}
			if suffixChecker(path, ".mp3") ||
				suffixChecker(path, ".ogg") ||
				suffixChecker(path, ".wav") ||
				suffixChecker(path, ".flac") {
				// routePath := slashReplacer(path)[3:]
				doThing(path)
				// dbInsert(dbName, path, routePath, path)
			} else {
				fmt.Println("things")
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
