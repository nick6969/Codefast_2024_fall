package controller

import (
	"github.com/gin-gonic/gin"
)

type answer struct {
	QuestionID string `json:"question_id"`
	Answer     struct {
		OptionID string `json:"option_id"`
		Value    string `json:"value"`
	} `json:"answer"`
}

func AnswerLabor(ctx *gin.Context) {
	var answers []answer

	if err := ctx.BindJSON(&answers); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// TODO: - Do something with the answers

	ctx.JSON(200, gin.H{"message": "Answers received"})
}
