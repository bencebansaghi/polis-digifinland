package main

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

var surveys = make([]Survey, 0)

type Question struct {
	Hash      string
	Question  string `json:"text"`
	Votes     int    `json:"votes"`
	Voters    []string
	Answers   []string `json:"answers"`
	Answerers []string
}

type Survey struct {
	Hash        string
	Questions   []Question `json:"questions"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}

func (s *Survey) existsUsername(username string) bool {
	for _, q := range s.Questions {
		for _, a := range q.Answerers {
			if a == username {
				return true
			}
		}
	}
	return false
}

func (s *Survey) canAnswer(username string, questionHash string) bool {
	for _, q := range s.Questions {
		if q.Hash != questionHash {
			continue
		}
		for _, a := range q.Answerers {
			if a == username {
				return true
			}
		}
	}
	return false
}

var (
	ErrAlreadyAnswered     = errors.New("already answered")
	ErrNonExistentQuestion = errors.New("question does not exist")
	ErrNonExistentSurvey   = errors.New("survey does not exist")
	ErrInternalError       = errors.New("internal error")
)

func getSurvey(hash string) (survey Survey, err error) {
	for _, s := range surveys {
		if s.Hash == hash {
			return s, nil
		}
	}
	return Survey{}, ErrNonExistentSurvey
}

func generateUsername() (un string) {
	return "user" + string(rand.Intn(1000))
}

func (s *Survey) newUserName() (un string) {
	un = generateUsername()
	for s.existsUsername(un) {
		un = generateUsername()
	}
	return
}

func (s *Survey) getQuestionByHash(hash string) (question Question, err error) {
	for _, q := range s.Questions {
		if q.Hash == hash {
			return q, nil
		}
	}
	return Question{}, ErrNonExistentQuestion
}

func (s *Survey) answer(username string, questionHash string, answer string) (err error) {
	if !s.canAnswer(username, questionHash) {
		return ErrAlreadyAnswered
	}

	for i, q := range s.Questions {
		if q.Hash == questionHash {
			s.Questions[i].Answers = append(q.Answers, answer)
			s.Questions[i].Answerers = append(q.Answerers, username)
			return nil
		}
	}
	return ErrNonExistentQuestion
}

func (s *Survey) canVote(username string) bool {
	for _, q := range s.Questions {
		for _, v := range q.Voters {
			if v == username {
				return false
			}
		}
	}
	return true
}

func (s *Survey) vote(username string, questionHash string, value int) (err error) {
	if !s.canVote(username) {
		return ErrAlreadyAnswered
	}
	if value != -1 && value != 1 {
		return ErrInternalError
	}

	for i, q := range s.Questions {
		if q.Hash == questionHash {
			s.Questions[i].Votes += value
			s.Questions[i].Voters = append(q.Voters, username)
			return nil
		}
	}
	return ErrNonExistentQuestion
}

func parseSurveysFolder() {
	// Scan surveys/ directory for json files
	files, err := os.ReadDir("surveys")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		data, err := os.ReadFile(filepath.Join("surveys", file.Name()))
		if err != nil {
			log.Println("Error reading file:", file.Name(), err)
			continue
		}

		var survey Survey
		if err := json.Unmarshal(data, &survey); err != nil {
			log.Println("Error unmarshalling file:", file.Name(), err)
			continue
		}
		survey.Hash = hash(survey.Title)
		for i, q := range survey.Questions {
			q.Hash = hash(q.Question)
			survey.Questions[i] = q
		}

		surveys = append(surveys, survey)
		log.Printf("Loaded survey: %s: %s\n", survey.Title, survey.Hash)
	}
}

func hash(str string) (hash string) {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	hash = fmt.Sprintf("%x", sha1.Sum(nil))
	return
}

func getByHash(hash string) (err error, survey Survey) {
	for _, s := range surveys {
		if s.Hash == hash {
			return nil, s
		}
	}
	return ErrNonExistentSurvey, Survey{}
}
