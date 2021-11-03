package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	startObfuscation()
}

func startObfuscation() {

	texts, err := getObTexts()

	if err != nil {
		log.Fatal(err)
	}
	obTexts := [][]string{}

	for _, v := range texts {

		en := Obfuscation(v)
		obTexts = append(obTexts, en)

	}
	s := GenerateGoCode(
		TplData{

			Obfuscated: obTexts,
		},
	)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filepath.Join(wd, filepath.Base(fmt.Sprintf("%s.go", "obfuscated"))), []byte(s), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}
