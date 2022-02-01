package wordle

import "errors"

var (
	WordNotExistsErr = errors.New("word doesn't exists")
	WordWasNotSet    = errors.New("word wasn't set")
)
