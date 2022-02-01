package wordle

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type WordleService interface {
	New() error
	Guess(word string) (string, error)
	IsGuessed() bool
	GetTries() int
}

type service struct {
	sync.RWMutex
	secrete        string
	dictionaryPath string
	wordsCount     int
	tries          int
	guessed        bool
}

func New(path string) (*service, error) {
	serv := &service{
		dictionaryPath: path,
	}

	err := serv.CountWords()
	if err != nil {
		return nil, err
	}

	return serv, nil
}

func (ws *service) New() error {
	ws.RLock()
	defer ws.RUnlock()

	randsource := rand.NewSource(time.Now().UnixNano())
	randgenerator := rand.New(randsource)
	pick := randgenerator.Intn(ws.wordsCount)

	file, err := os.Open(ws.dictionaryPath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("dictionary file closed with error: %v", err.Error())
		}
	}()

	scanner := bufio.NewScanner(file)
	var lineNum int
	for scanner.Scan() {
		if lineNum == pick {
			ws.secrete = scanner.Text()
			break
		}
		lineNum++
	}
	if err = scanner.Err(); err != nil {
		log.Println(err)
	}

	ws.tries = 0
	ws.guessed = false

	return nil

}

func (ws *service) GetWord() string {
	ws.RLock()
	defer ws.RUnlock()

	return ws.secrete

}

func (ws *service) CountWords() error {
	ws.RLock()
	defer ws.RUnlock()

	file, err := os.Open(ws.dictionaryPath)
	if err != nil {
		log.Println(err)
		return err
	}

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	err = file.Close()
	if err != nil {
		log.Printf("dictionary file closed with error: %v", err.Error())
		return err
	}

	ws.wordsCount = count

	return nil
}

func (ws *service) IsGuessed() bool {
	ws.RLock()
	defer ws.RUnlock()
	return ws.guessed

}

func (ws *service) GetTries() int {
	ws.RLock()
	defer ws.RUnlock()
	return ws.tries

}
