package gormModels

type Training struct {
	Email         string `json:"email"`
	QuestionOne   int    `json:"questionOne"`
	QuestionTwo   int    `json:"questionTwo"`
	QuestionThree int    `json:"questionThree"`
	QuestionFour  int    `json:"questionFour"`
	QuestionFive  int    `json:"questionFive"`
	QuestionSix   int    `json:"questionSix"`
	QuestionSeven int    `json:"questionSeven"`
	QuestionEight int    `json:"questionEight"`
	QuestionNine  int    `json:"questionNine"`
	QuestionTen   int    `json:"questionTen"`
	Score         int    `json:"score"`
}
