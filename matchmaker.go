package main

import (
	"fmt"
	"math"
)

var questions = []string{
	"What is your favorite number between 1 and 5",
	"C# is the best all around programming language",
	"The Blackhawks are going to make a come back next season",
	"Python is an amazing programming language",
	"The Cheifs are going to win the Super Bowl this year",
}

var my_answers = []int{4, 5, 5, 1, 3}
var user_answers = []int{0, 0, 0, 0, 0}
var question_scores = []int{0, 0, 0, 0, 0}

var max_score = 4 * len(questions)

var total_score = 0

func calculate_score(original int, answer int) int {
	return 4 - int(math.Abs(float64(original-answer)))
}

func get_user_input(prompt string) int {
	var input int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanf("%d\n", &input)
		if err == nil {
			break
		}
		fmt.Println(err)
		fmt.Println("Invalid input. Please enter a number.")
		fmt.Scanf("%s")
	}
	return input
}

func main() {

	fmt.Println("******************************************************")
	fmt.Println("                    Matchmaker 3.0                    ")
	fmt.Println("******************************************************")

	fmt.Println("This program figures out if you and I are meant to be.")
	fmt.Println("You will answer 5 questions. To each question, answer 5")
	fmt.Println("if you strongly agree, 4 if you agree, 3 if you neither")
	fmt.Println("agree nor disagree, 2 if you disagree, and 1 if you")
	fmt.Println("strongly disagree.")
	fmt.Println("")
	fmt.Println("")

	for i := 0; i < len(questions); i++ {
		fmt.Print("Question ")
		fmt.Print(i + 1)
		fmt.Println(":")
		fmt.Println("-----------")
		user_answers[i] = get_user_input(questions[i] + ": ")
		question_scores[i] = calculate_score(my_answers[i], user_answers[i])
		total_score += question_scores[i]
		fmt.Println()
	}

	calculated_percentage := float64(total_score) / float64(max_score) * 100.0
	fmt.Print("Your score is: ")
	fmt.Print(calculated_percentage)
	fmt.Println("%")

	if calculated_percentage == 100 {
		fmt.Println("You are a perfect match!")
	} else if calculated_percentage > 80 {
		fmt.Println("You are a great match!")
	} else if calculated_percentage > 60 {
		fmt.Println("Maybe we should just be friends.")
	} else {
		fmt.Println("You are a terrible match! Get out!")
	}

	fmt.Println()

	fmt.Println("Overview")
	fmt.Println("--------")
	for i := 0; i < len(questions); i++ {
		fmt.Print(questions[i])
		fmt.Println(": ")
		fmt.Print("You answered ")
		fmt.Print(user_answers[i])
		fmt.Print(" and my answer is ")
		fmt.Println(my_answers[i])
		fmt.Print("Your compatibility score for this question is ")
		fmt.Print(float64(question_scores[i]) / 4.0 * 100)
		fmt.Println("%")
		fmt.Println()
	}
}
