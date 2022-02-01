package wordle

import (
	"os"
	"testing"
)

func Test_service_wordExists(t *testing.T) {
	dictPath := "./test_dict"
	dict := []byte("банан\nборец\nзабор\n")
	err := os.WriteFile(dictPath, dict, 0644)
	if err != nil {
		t.Errorf("coudn't create test file %v", err)
	}
	defer func() {
		err = os.Remove(dictPath)
		if err != nil {
			t.Errorf("couldn't delete test file %v", err)
		}
	}()

	ws := service{
		secrete:        "банан",
		dictionaryPath: dictPath,
		wordsCount:     3,
		tries:          0,
		guessed:        false,
	}
	tests := []struct {
		name    string
		word    string
		want    bool
		wantErr bool
	}{
		{
			name:    "existing word",
			word:    "банан",
			want:    true,
			wantErr: false,
		},
		{
			name:    "not existing word",
			word:    "арара",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ws.wordExists(tt.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("wordExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("wordExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
