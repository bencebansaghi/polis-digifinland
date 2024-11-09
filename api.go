package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenericResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ChallengeInitiationResponse struct {
	GenericResponse
	NumberA int `json:"number_a"`
	NumberB int `json:"number_b"`
}

type ChallengeSubmissionRequest struct {
	Hash             string `json:"id"`
	CalculatedResult int    `json:"calculated_result"`
}

type ChallengeSubmissionResponse struct {
	GenericResponse
	Message  string `json:"message"`
	Survey   Survey `json:"survey"`
	Username string `json:"username"`
}

type RestoreRequest struct {
	Hash     string `json:"id"`
	Username string `json:"username"`
}

type LoginResponse struct {
	GenericResponse
	Survey Survey `json:"survey"`
}

type SurveyRequest struct {
	Hash         string `json:"id"`
	QuestionHash string `json:"qid"`
	Username     string `json:"username"`
	Answer       string `json:"answer"`
}

func newCaptchaHandler(c *gin.Context) {
	fbc := generateFbChallenge()

	c.SetCookie("captcha", fbc.Id, 3600, "/", "", false, true)

	response := ChallengeInitiationResponse{
		NumberA: fbc.A,
		NumberB: fbc.B,
	}
	response.Success = true

	c.JSON(http.StatusOK, response)
}

func submitCaptchaHandler(c *gin.Context) {
	var request ChallengeSubmissionRequest
	var response ChallengeSubmissionResponse
	response.Success = false

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Message = "Bad request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	captchaCookie, err := c.Cookie("captcha")
	if err != nil {
		response.Message = "Invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if validateFbChallengeSolution(captchaCookie, request.CalculatedResult) {
		response.Message = "Bad solution"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	survey, err := getSurvey(request.Hash)
	if err != nil {
		response.Message = "Survey not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	discardFbChallenge(captchaCookie)
	response.Success = true
	response.Message = "Success"
	response.Survey = survey
	response.Username = survey.newUserName()

	c.SetCookie("username", response.Username, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, response)
}

func loginHandler(c *gin.Context) {
	var request RestoreRequest
	var response LoginResponse
	response.Success = false

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Message = "Invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	survey, err := getSurvey(request.Hash)
	if err != nil {
		response.Message = "Survey not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	if !survey.existsUsername(request.Username) {
		response.Message = "Username not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.Success = true
	response.Message = "Success"
	response.Survey = survey

	c.JSON(http.StatusOK, response)
}

func surveyAnswerHandler(c *gin.Context) {
	var request SurveyRequest
	var response GenericResponse
	response.Success = false

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Message = "Invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	survey, err := getSurvey(request.Hash)
	if err != nil {
		response.Message = "Survey not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	if !survey.canAnswer(request.Username, request.QuestionHash) {
		response.Message = "Cannot answer"
		c.JSON(http.StatusForbidden, response)
		return
	}

	question, err := survey.getQuestionByHash(request.QuestionHash)
	if err != nil {
		response.Message = "Question not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	question.Answerers = append(question.Answerers, request.Username)
	question.Answers = append(question.Answers, request.Answer)

	response.Success = true
	response.Message = "Success"

	c.JSON(http.StatusOK, response)
}

func initializeApiRoutes() {
	api := router.Group("/api")
	{
		api.GET("/newcaptcha", newCaptchaHandler)
		api.POST("/submitcaptcha", submitCaptchaHandler)
		api.POST("/login", loginHandler)
		api.POST("/survey", surveyAnswerHandler)
	}
}
