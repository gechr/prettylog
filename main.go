package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/rs/zerolog"
)

func init() {
	log.SetFlags(0)
}

func main() {
	w := zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			w.FormatLevel = formatLevel(w.NoColor)
			w.PartsExclude = []string{zerolog.CallerFieldName}
			w.TimeFormat = "15:04:05.000"
		},
	)

	var query string
	switch len(os.Args[1:]) {
	case 0:
		break
	case 1:
		query = os.Args[1]
	default:
		log.Fatalln("usage: prettylog [query]")
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		r := bytes.NewReader(s.Bytes())
		if err := jq(r, w, query); err != nil {
			if errors.Is(err, errInvalidJSON) {
				fmt.Println(s.Text())
			} else {
				log.Fatalln(err)
			}
		}
	}
}
