/*
Copyright © 2022 Ashis Sharma <ashisrm9@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ashissharma97/commandlinequiz/models"
	"github.com/manifoldco/promptui"
	"github.com/wzshiming/ctc"

	"github.com/spf13/cobra"
)

const API_URL string = "https://opentdb.com/api.php?amount=10&type=multiple"

var questions []models.Question

var getCmd = &cobra.Command{
	Use:   "start",
	Short: "To start the quiz",
	Long:  `To start the quiz and you'll get the questions select the answer from the options and press enter.`,
	Run: func(cmd *cobra.Command, args []string) {
		getQuestions()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getQuestions() {
	resp, err := http.Get(API_URL)
	if err != nil {
		errors.New("Something went wrong!!!")
	}

	data, _ := io.ReadAll(resp.Body)

	var jsonData interface{}
	json.Unmarshal(data, &jsonData)

	results := jsonData.(map[string]interface{})["results"].([]interface{})

	for _, result := range results {
		question := result.(map[string]interface{})
		questions = append(questions, models.Question{
			Question:         question["question"].(string),
			CorrectAnswer:    question["correct_answer"].(string),
			IncorrectAnswers: question["incorrect_answers"].([]interface{}),
			Difficulty:       question["difficulty"].(string),
		})
	}

	templates := &promptui.SelectTemplates{
		Selected: "{{ . }}",
	}

	correct := 0

	for _, question := range questions {

		prompt := promptui.Select{
			Label:     question.GetQuestion(),
			Items:     question.GetAllOptions(),
			Templates: templates,
		}

		_, result, err := prompt.Run()

		if err != nil {
			errors.New("Something went wrong!!!")
		}

		if checkAnswer(question, result) {
			fmt.Println(ctc.ForegroundGreen, "Correct Answer", ctc.Reset)
			correct++
		} else {
			fmt.Println(ctc.ForegroundRed, "Wrong Answer", ctc.Reset)
		}
	}

	fmt.Println(ctc.ForegroundGreen, strconv.Itoa(correct)+" out of 10 is correct answers", ctc.Reset)

	defer resp.Body.Close()
}

func checkAnswer(question models.Question, answer string) bool {
	return question.GetCorrectAnswer() == answer
}
