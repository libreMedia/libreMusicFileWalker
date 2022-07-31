package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dhowden/tag"
)

func doThing(dbName string, name string, routePath string, path string) {
	file, err := os.Open(path) // For read access.

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

		var title, artist, album, year, givenGenre, votedGenre, comment, composer, lyrics string

		if m == nil {
			fmt.Println("no data")
		} else {

			if m.Title() != "" {
				title = m.Title()
				log.Print(m.Title()) // The title of the track (see Metadata interface for more details).
			} else {
				title = "no title"
				fmt.Println("no title")
			}
			if m.Artist() != "" {
				artist = m.Artist()
				log.Print(m.Artist()) // The artist of the track (see Metadata interface for more details).
			} else {
				artist = "no artist"
				fmt.Println("no artist")
			}
			if m.Album() != "" {
				album = m.Album()
				log.Print(m.Album()) // The album of the track (see Metadata interface for more details).
			} else {
				album = "no album"
				fmt.Println("no album")
			}
			if m.Year() != 0 {
				year = strconv.Itoa(m.Year())
				log.Print(m.Year()) // The year of the track (see Metadata interface for more details).
			} else {
				year = "no year"
				fmt.Println("no year")
			}
			if m.Genre() != "" {
				givenGenre = m.Genre()
				log.Print(m.Genre()) // The genre of the track (see Metadata interface for more details).
			} else {
				givenGenre = "nill"
				fmt.Println("no genre")
			}
			if m.Comment() != "" {
				comment = m.Comment()
				log.Print(m.Comment()) // The comment of the track (see Metadata interface for more details).
			} else {
				comment = "no comment"
				fmt.Println("no comment")
			}
			if m.Composer() != "" {
				composer = m.Composer()
				log.Print(m.Composer()) // The composer of the track (see Metadata interface for more details).
			} else {
				composer = "no composer"
				fmt.Println("no composer")
			}
			if m.Lyrics() != "" {
				lyrics = m.Lyrics()
				log.Print(m.Lyrics()) // The lyrics of the track (see Metadata interface for more details).
			} else {
				lyrics = "no lyrics"
				fmt.Println("no lyrics")
			}
			// s := newStructy(name, routePath, path, m.Title(), m.Artist(), m.Album(), m.Year(), m.Genre(), "GenreTBD", m.Comment(), m.Composer(), m.Lyrics())
			// log.Print(m.Format())      // The detected format.
			// log.Print(m.AlbumArtist()) // The title of the track (see Metadata interface for more details).
			// log.Print(m.Composer())    // The title of the track (see Metadata interface for more details).
			// r := s
			// fmt.Print(r)
		}

		dbInsert(dbName, path, routePath, path, title, artist, album, year, givenGenre, votedGenre, comment, composer, lyrics)
	} else {
		println("no file")
		// r := newStructy("", "", "", "", "", "", 0, "", "", "", "", "")
		// return r
	}
}

func walk(dirToWalk string, dbName string) {
	err := filepath.Walk(getParentDir()+dirToWalk,
		func(path string, info os.FileInfo, err error) error {
			filename := filepath.Join(path, info.Name())
			if err != nil {
				return err
			}
			if suffixChecker(path, ".mp3") ||
				suffixChecker(path, ".ogg") ||
				suffixChecker(path, ".wav") ||
				suffixChecker(path, ".flac") {
				routePath := slashReplacer(path)[3:]
				doThing(dbName, filename, routePath, path)
			} else {
				fmt.Println("things")
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
