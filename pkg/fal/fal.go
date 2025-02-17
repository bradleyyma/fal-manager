package fal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/bradleyyma/falmanager/pkg/fal/model"
)

func ParseRules() {}

func EvaluateFAL(teamInfoFile, clientID string) {
	file, err := os.Open(teamInfoFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	team := &model.Team{}
	if err := json.Unmarshal(bytes, team); err != nil {
		log.Fatal(err)
	}

	file, err = os.Open("rules.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	rules := struct {
		PointCriterias []*model.PointCriteria `json:"weeklyPointCriteria"`
	}{}
	if err := json.Unmarshal(bytes, &rules); err != nil {
		log.Fatal(err)
	}

	for _, anime := range team.Active {
		err := anime.GetInfo(clientID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Anime: %#v", anime.Title)
	}

	for _, anime := range team.Bench {
		err := anime.GetInfo(clientID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Anime: %#v", anime.Title)
	}

	fmt.Printf("\nTeam: %v", team)

	score := calculateScore(*team, 1, rules.PointCriterias)
	fmt.Printf("Score: %d", score)

}

func calculateScore(team model.Team, week int, pc []*model.PointCriteria) int {
	score := 0
	if week > len(pc)+1 {
		log.Fatal("Week out of range!")
	}
	criteria := pc[week-1]
	for _, anime := range team.Active {
		animeScore := 0
		watching, err := strconv.Atoi(anime.Statistics.Status.Watching)
		if err != nil {
			log.Fatal(err)
		}
		animeScore += int(criteria.Watching * float32(watching))
		fmt.Printf("Score for %s is: %d, from %d * %d", anime.Title, animeScore, criteria.Watching, watching)
		score += animeScore
	}
	return score
}
