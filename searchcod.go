package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var resposta string

func main() {
	// Open the file.
	f, _ := os.Open("access_log")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		Searchproduto(line)
	}
}

func Searchproduto(a string) {

	//re := regexp.MustCompile(`&codproduto=(.*)&`)
	re := regexp.MustCompile(`&codproduto=[a-zA-Z]*(.*)&o`)
	//re := regexp.MustCompile(`carrinhocompra` + `&codcliente=(.*)&c`) //carinho de compras

	submatchall := re.FindAllStringSubmatch(a, -1)
	for _, element := range submatchall {
		fmt.Println(element[1])
	}

}
