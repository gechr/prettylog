package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/itchyny/gojq"
)

var (
	errInvalidJSON = errors.New("invalid json")
	errInvalidJQ   = errors.New("invalid jq")
)

func jq(r io.Reader, w io.Writer, query string) error {
	if query == "" {
		_, err := io.Copy(w, r)
		if err != nil {
			return fmt.Errorf("%w: %w", errInvalidJSON, err)
		}
		return nil
	}

	query = fmt.Sprintf("select(%s)", query)
	q, err := gojq.Parse(query)
	if err != nil {
		return fmt.Errorf("%w: %w", errInvalidJQ, err)
	}
	var input interface{}
	err = json.NewDecoder(r).Decode(&input)
	if err != nil {
		return fmt.Errorf("%w: %w", errInvalidJSON, err)
	}
	iter := q.Run(input)
	for {
		v, ok := iter.Next()
		if !ok {
			return nil
		}
		if err, ok := v.(error); ok {
			return err
		}
		result, err := gojq.Marshal(v)
		if err != nil {
			return err
		}
		fmt.Fprintln(w, string(result))
	}
}
