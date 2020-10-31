package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
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

	question_results := map[string]int{
		"right": 0,
		"wrong": 0,
	}

	// var reader bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Let's start")
	for _, v := range question_data {
		fmt.Println(v.question)
		text, err := reader.ReadString('\n')
		check(err)
		text = clear(text)
		if strings.Compare(text, v.answer) == 0 {
			question_results["right"]++
		} else {
			question_results["wrong"]++
		}

	}
	fmt.Printf("Your results:\n right - %v \t wrong - %v ",
		question_results["right"],
		question_results["wrong"])

}
func clear(old string) (new string) {
	new = strings.Replace(old, "\n", "", -1)
	new = strings.ReplaceAll(new, " ", "")
	return new
}
