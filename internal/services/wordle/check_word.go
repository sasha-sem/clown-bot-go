package wordle

import (
	"bufio"
	"log"
	"os"
)

func (ws *service) wordExists(word string) (bool, error) {
	ws.RLock()
	defer ws.RUnlock()

	file, err := os.Open(ws.dictionaryPath)
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("dictionary file closed with error: %v", err.Error())
		}
	}()

	if err != nil {
		log.Println(err)
		return false, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if word == scanner.Text() {
			return true, nil
		}

	}
	return false, nil
}
