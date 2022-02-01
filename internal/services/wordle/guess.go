package wordle

import (
	"log"
	"strings"
)

func (ws *service) Guess(word string) (string, error) {
	ws.RLock()
	defer ws.RUnlock()

	if ws.guessed || ws.secrete == "" {
		return "", WordWasNotSet
	}
	yes, err := ws.wordExists(word)
	if err != nil {
		log.Printf("error ocured while checking word %v", err)
		return "", err
	}
	if !yes {
		return "", WordNotExistsErr
	}

	ws.tries += 1

	// word guessed
	if ws.secrete == word {
		ws.guessed = true
	}

	return compareWords(word, ws.secrete), nil
}

func compareWords(word, secrete string) string {
	checkedChars := map[int]bool{
		0: false,
		1: false,
		2: false,
		3: false,
		4: false,
	}
	answerRunes := []rune("uuuuu") // status of all letters unknown
	wordRunes := []rune(word)
	secreteRunes := []rune(secrete)

	// finding all correctly placed letters
	for pos, char := range wordRunes {
		if secreteRunes[pos] == char {
			answerRunes[pos] = 'g' // status green means green square
			checkedChars[pos] = true
			continue
		}
	}

	// finding incorrectly placed and wrong letters
	for pos, char := range wordRunes {
		// if status already known - skip
		if answerRunes[pos] != 'u' {
			continue
		}

		// if no letter in secret - set as white square
		if !strings.Contains(secrete, string(char)) {
			answerRunes[pos] = 'w'
			continue
		}

		wrongPlacedLetterIndex := -1
		// looking for wrong placed letters
		for secPos, secChar := range secreteRunes {
			// if this letter doesn't have status but in the secrete
			if secChar == char && !checkedChars[secPos] {
				wrongPlacedLetterIndex = secPos
				break
			}
		}
		// if we found such letter set it to yellow square else - wrong letter
		if wrongPlacedLetterIndex > -1 {
			answerRunes[pos] = 'y'
			checkedChars[wrongPlacedLetterIndex] = true
		} else {
			answerRunes[pos] = 'w'
		}

	}
	return buildAnswer(answerRunes)
}

func buildAnswer(answerRunes []rune) string {
	answer := ""

	for _, r := range answerRunes {
		switch r {
		case 'g':
			answer = answer + CorrectLetter
		case 'y':
			answer = answer + WrongPlacedLetter
		default:
			answer = answer + IncorrectLetter

		}
	}
	return answer
}
