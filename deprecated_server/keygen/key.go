package keygen

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const generationDelta = time.Hour
const wordLength = 6
const path = ".key"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetKey() string {
	if value, err := ioutil.ReadFile(path); err != nil {
		fmt.Printf("[%s] Error reading .key file!\n", time.Now().Format("15:04:05"))
		panic(err)
	} else {
		return string(value)
	}
}

func IsValid(candidate string) bool {
	if value, err := ioutil.ReadFile(path); err != nil {
		fmt.Printf("[%s] Error reading .key file!\n", time.Now().Format("15:04:05"))
		panic(err)
	} else {
		return candidate == string(value)
	}
}

func Generate() {
	for {
		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("[%s] Error opening .key file!\n", time.Now().Format("15:04:05"))
			panic(err)
		}
		defer file.Close()

		if _, err := file.WriteString(randomWord()); err != nil {
			fmt.Printf("[%s] Error writing to .key file!\n", time.Now().Format("15:04:05"))
			panic(err)
		}

		time.Sleep(generationDelta)
	}
}

func randomWord() string {
	charset := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, wordLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
