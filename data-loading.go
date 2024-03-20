package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	// abrindo o arquivo
	data, err := os.Open("./dataset.csv")

	if err != nil {
		log.Fatal(err)
	}

	// criando um dataframe
	df := dataframe.ReadCSV(data)

	// imrimindo na tela
	fmt.Println(df)
}
