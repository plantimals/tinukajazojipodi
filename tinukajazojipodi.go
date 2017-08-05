package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

var (
	consonants = []rune("bcdfghjklmnpqrstvwxz")
	vowels     = []rune("aeiouy")
	voices     = []string{
		"Alex",
		"Alice",
		"Alva",
		"Amelie",
		"Anna",
		"Carmit",
		"Damayanti",
		"Daniel",
		"Diego",
		"Ellen",
		"Fiona",
		"Fred",
		"Ioana",
		"Joana",
		"Jorsge",
		"Juan",
		"Kanya",
		"Karen",
		"Kyoko",
		"Laura",
		"Lekha",
		"Luca",
		"Luciana",
		"Maged",
		"Mariska",
		"Mei-Jia",
		"Melina",
		"Milena",
		"Moira",
		"Monica",
		"Nora",
		"Paulina",
		"Samantha",
		"Sara",
		"Satu",
		"Sin-ji",
		"Tessa",
		"Thomas",
		"Ting-Ting",
		"Veena",
		"Victoria",
		"Xander",
		"Yelda",
		"Yuna",
		"Yuri",
		"Zosia",
		"Zuzana",
	}
)

func main() {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	sentence := randomSentence(r)
	voice := randomVoice(r)
	fmt.Printf("say -v %s \"%s\"\n", voice, sentence)
	cmd := exec.Command("say", "-v", voice, sentence)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func randomSentence(r *rand.Rand) string {
	length := r.Int()%5 + 1
	answer := randomWord(r)
	for i := 0; i < length-1; i++ {
		answer = fmt.Sprintf("%s %s", answer, randomWord(r))
	}
	return answer
}

func randomWord(r *rand.Rand) string {
	length := r.Int()%15 + 5
	answer := ""
	for i := 0; i < length; i++ {
		letter := ""
		if i%2 == 1 {
			letter = randomConsonant(r)
		} else {
			letter = randomVowel(r)
		}
		answer = fmt.Sprintf("%s%s", answer, letter)
	}

	return answer
}

func randomConsonant(r *rand.Rand) string {
	return string(consonants[r.Intn(len(consonants))])
}

func randomVowel(r *rand.Rand) string {
	return string(vowels[r.Intn(len(vowels))])
}

func randomVoice(r *rand.Rand) string {
	return voices[r.Intn(len(voices))]
}
