package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SurveyPageData struct {
	SurveyTitle       string
	SurveyDescription string
	QuestionsAnswered int
	QuestionsTotal    int
	Username          string
	IsAuth            bool
}

func loadTemplates(router *gin.Engine) {
	router.LoadHTMLGlob("./public/*.html")

	router.Static("/static", "./public")
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func surveyHandler(c *gin.Context) {
	// Get survey ID from query parameters
	surveyID := c.Query("id")

	survey, err := getSurvey(surveyID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	username, err := c.Cookie("username")
	if err != nil {
		username = ""
	}

	pageData := SurveyPageData{
		SurveyTitle:       survey.Title,
		SurveyDescription: survey.Description,
		QuestionsAnswered: 0,
		QuestionsTotal:    len(survey.Questions),
		Username:          username,
		IsAuth:            len(username) > 0,
	}

	if surveyID == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Survey ID is required",
		})
		return
	}

	c.HTML(http.StatusOK, "survey.html", pageData)
}

func initializeRoutes() {
	// Load templates
	loadTemplates(router)

	// Public routes
	router.GET("/", indexHandler)
	router.GET("/survey", surveyHandler)

	setupErrorTemplates(router)
}

func setupErrorTemplates(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "Page not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.HTML(http.StatusMethodNotAllowed, "error.html", gin.H{
			"error": "Method not allowed",
		})
	})
}
