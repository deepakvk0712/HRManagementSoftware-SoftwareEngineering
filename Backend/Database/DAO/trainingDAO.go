package Dao

import (
	"fmt"
	models "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models"
	gormModels "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Models/GormModels"
	utils "hrtool.com/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	"strings"
)

func GetTrainingScore(email string) (int, int) {
	score := -1

	row := utils.Db.Raw("SELECT SCORE FROM TRAININGS WHERE LOWER(EMAIL) = ?", strings.ToLower(email)).Row()
	if row.Err() != nil {
		fmt.Println(row)
		return 0, 0
	}

	row.Scan(&score)

	return score, 1
}

func CalculateScore(posted models.TrainingQuestions, email string) (int, int) {
	actual := gormModels.Training{}
	training := gormModels.Training{}

	utils.Db.Exec("DELETE FROM TRAININGS WHERE LOWER(EMAIL) = ?", strings.ToLower(email))

	row := utils.Db.Raw("SELECT * FROM TRAININGS WHERE LOWER(EMAIL) = ?", "solution").Row()

	row.Scan(&actual.Email, &actual.QuestionOne, &actual.QuestionTwo, &actual.QuestionThree, &actual.QuestionFour, &actual.QuestionFive, &actual.QuestionSix, &actual.QuestionSeven, &actual.QuestionEight, &actual.QuestionNine, &actual.QuestionTen, &actual.Score)

	training.Email = email
	training.Score = 0

	training.QuestionOne = posted.QuestionOne
	if actual.QuestionOne == posted.QuestionOne {
		training.Score += 1
	}
	training.QuestionTwo = posted.QuestionTwo
	if actual.QuestionTwo == posted.QuestionTwo {
		training.Score += 1
	}
	training.QuestionThree = posted.QuestionThree
	if actual.QuestionThree == posted.QuestionThree {
		training.Score += 1
	}
	training.QuestionFour = posted.QuestionFour
	if actual.QuestionFour == posted.QuestionFour {
		training.Score += 1
	}
	training.QuestionFive = posted.QuestionFive
	if actual.QuestionFive == posted.QuestionFive {
		training.Score += 1
	}
	training.QuestionSix = posted.QuestionSix
	if actual.QuestionSix == posted.QuestionSix {
		training.Score += 1
	}
	training.QuestionSeven = posted.QuestionSeven
	if actual.QuestionSeven == posted.QuestionSeven {
		training.Score += 1
	}
	training.QuestionEight = posted.QuestionEight
	if actual.QuestionEight == posted.QuestionEight {
		training.Score += 1
	}
	training.QuestionNine = posted.QuestionNine
	if actual.QuestionNine == posted.QuestionNine {
		training.Score += 1
	}
	training.QuestionTen = posted.QuestionTen
	if actual.QuestionTen == posted.QuestionTen {
		training.Score += 1
	}

	result := utils.Db.Create(&training)

	if result.Error != nil {
		fmt.Println(result.Error)

		return 0, 0
	}

	return training.Score, 1
}
