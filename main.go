package main

import (
	"bufio"
	"bytes"
	ctr "crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
	"unsafe"
)

var (
	Tmp = `
package main
import (
	"unsafe"
)
func eax() uint8{
	return uint8(unsafe.Sizeof(true))
}

{{range .Obfuscated}}

//  [ {{index . 2}} ] ====>  {{index . 0}} 


func {{index . 2}}() string {
            {{index . 1}}
}




{{end}}

`
)

type TplData struct {
	Obfuscated [][]string
}

const (
	EAX = uint8(unsafe.Sizeof(true))
	ONE = "eax()"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getObTexts() ([]string, error) {

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if f.Name() == ".ob" {
			file, err := os.Open("./.ob")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			texts := []string{}
			for scanner.Scan() {
				texts = append(texts, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				return nil, fmt.Errorf("cound not read the .ob file")
			}
			if len(texts) <= 0 {
				return nil, fmt.Errorf("make sure .ob file contains atleast 1 line of string ")
			}
			return texts, nil
		}
	}

	return nil, fmt.Errorf(".ob file cound not be found in the current directory")

}

func Obfuscation(txt string) []string {
	rand.Seed(time.Now().UnixNano())
	lines := []string{}

	for _, item := range []byte(txt) {
		lines = append(
			lines, GetNumber(item),
		)
	}
	n := 3
	b := make([]byte, n)
	if _, err := ctr.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return []string{txt, fmt.Sprintf(
		"return string(\n[]byte{\n%s,\n},\n)",
		strings.Join(lines, ",\n"),
	), RandStringRunes(3) + s}
}

func GetNumber(n byte) (buf string) {
	var arr []byte
	var x uint8

	for n > EAX {
		x = 0

		if n%2 == EAX {
			x = EAX
		}

		arr = append(arr, x)

		n = n >> EAX
	}

	buf = ONE

	rand.Seed(
		time.Now().Unix(),
	)

	for i := len(arr) - 1; i >= 0; i-- {
		buf = fmt.Sprintf(
			"%s<<%s", buf, ONE,
		)

		if arr[i] == EAX {
			op := "(%s|%s)"

			if rand.Intn(2) == 0 {
				op = "(%s^%s)"
			}

			buf = fmt.Sprintf(
				op, buf, ONE,
			)
		}
	}

	return buf
}

func GenerateGoCode(data TplData) string {
	tmpl := template.Must(
		template.New("x").Parse(Tmp),
	)

	w := new(bytes.Buffer)
	err := tmpl.Execute(w, data)

	if err != nil {
		panic(err)
	}

	return w.String()
}

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
