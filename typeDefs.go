package main

type structy struct {
	name       string
	routePath  string
	path       string
	title      string
	artist     string
	album      string
	year       string
	givenGenre string
	votedGenre string
	comment    string
	composer   string
	lyrics     string
}

func newStructy(name string, routePath string, path string, title string, artist string, album string, year string, givenGenre string, votedGenre string, comment string, composer string, lyrics string) *structy {
	s := structy{
		name:       name,
		routePath:  routePath,
		path:       path,
		title:      title,
		artist:     artist,
		album:      album,
		year:       year,
		givenGenre: givenGenre,
		votedGenre: votedGenre,
		comment:    comment,
		composer:   composer,
		lyrics:     lyrics,
	}

	return &s
}
