package wordle

import (
	"testing"
)

func Test_buildAnswer(t *testing.T) {
	tests := []struct {
		name        string
		answerRunes []rune
		want        string
	}{
		{
			name:        "all green",
			answerRunes: []rune("ggggg"),
			want:        ":green_square::green_square::green_square::green_square::green_square:",
		},
		{
			name:        "all yellow",
			answerRunes: []rune("yyyyy"),
			want:        ":orange_square::orange_square::orange_square::orange_square::orange_square:",
		},
		{
			name:        "all white",
			answerRunes: []rune("wwwww"),
			want:        "️:white_large_square:️:white_large_square:️:white_large_square:️:white_large_square:️:white_large_square:",
		},
		{
			name:        "mixed",
			answerRunes: []rune("gywyg"),
			want:        ":green_square::orange_square:️:white_large_square::orange_square::green_square:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildAnswer(tt.answerRunes); got != tt.want {
				t.Errorf("buildAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareWords(t *testing.T) {
	tests := []struct {
		name    string
		word    string
		secrete string
		want    string
	}{
		{
			name:    "all missed",
			word:    "топор",
			secrete: "банан",
			want:    "️:white_large_square:️:white_large_square:️:white_large_square:️:white_large_square:️:white_large_square:",
		},
		{
			name:    "mixed answer",
			word:    "закон",
			secrete: "банан",
			want:    "️:white_large_square::green_square:️:white_large_square:️:white_large_square::green_square:",
		},
		{
			name:    "mixed answer",
			word:    "карта",
			secrete: "банан",
			want:    "️:white_large_square::green_square:️:white_large_square:️:white_large_square::orange_square:",
		},
		{
			name:    "mixed answer",
			word:    "карта",
			secrete: "банан",
			want:    "️:white_large_square::green_square:️:white_large_square:️:white_large_square::orange_square:",
		},
		{
			name:    "right guess",
			word:    "закон",
			secrete: "закон",
			want:    ":green_square::green_square::green_square::green_square::green_square:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareWords(tt.word, tt.secrete); got != tt.want {
				t.Errorf("compareWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
