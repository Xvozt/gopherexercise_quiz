package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Question struct {
	question string
	answer   string
}

func main() {
	file_data, err := os.Open("/Users/sergeykupriyanov/go/src/githubcom/xvozt/quizgame/problems.csv")
	check(err)
	r := csv.NewReader(file_data)

	var question_data []Question

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		question := Question{}
		question.question = record[0]
		question.answer = record[1]
		question_data = append(question_data, question)
	}

	for _, v := range question_data {
		fmt.Println(v)
	}

}
