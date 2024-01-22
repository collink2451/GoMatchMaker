package main

import (
	"fmt"
	"math"
)

// create array of questions that also includes a number for the correct answer
var questions = []string{
	"This is Question 1",
	"This is Question 2",
	"This is Question 3",
	"This is Question 4",
	"This is Question 5",
}

var answers = []int{1, 2, 3, 4, 5}
var user_answers = []int{0, 0, 0, 0, 0}

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
	fmt.Println("Matchmaker 3.0")
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
		fmt.Println("----------")
		user_answers[i] = get_user_input(questions[i] + ": ")
		total_score += calculate_score(answers[i], user_answers[i])
		fmt.Println()
	}

	calculated_percentage := float64(total_score) / float64(max_score) * 100.0
	fmt.Print("Your score is: ")
	fmt.Print(calculated_percentage)
	fmt.Println("%")

	if calculated_percentage == 100 {
		fmt.Println("You are a perfect match!")
	} else if calculated_percentage > 90 {
		fmt.Println("You are a great match!")
	} else if calculated_percentage > 70 {
		fmt.Println("You are an ok match!")
	} else {
		fmt.Println("You are a terrible match! Get out!")
	}
}
