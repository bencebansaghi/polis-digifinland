package main

import "math/rand"

type ScryptParams struct {
	Cost      int `json:"N"`
	BlockSize int `json:"r"`
	Parallel  int `json:"p"`
	KeyLen    int `json:"l"`
}

type Challenge struct {
	ScryptParams
	Preimage        string `json:"preimage"`
	Difficulty      string `json:"difficulty"`
	DifficultyLevel int    `json:"difficulty_level"`
}

var fbChallenges = make(map[string]FbChallenge)

type FbChallenge struct {
	Id string `json:"id"`
	A  int    `json:"a"`
	B  int    `json:"b"`
}

func generateFbChallenge() (fbc FbChallenge) {
	fbc = FbChallenge{
		Id: string(len(fbChallenges)),
		A:  rand.Intn(100),
		B:  rand.Intn(100),
	}
	return
}

func validateFbChallengeSolution(id string, solution int) bool {
	fbc, ok := fbChallenges[id]
	if !ok {
		return false
	}
	return fbc.A+fbc.B == solution
}

func discardFbChallenge(id string) {
	delete(fbChallenges, id)
}
