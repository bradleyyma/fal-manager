package model

import (
	"fmt"
	"strings"

	"github.com/bradleyyma/falmanager/internal/anime"
)

type Team struct {
	Bench  []*anime.Anime
	Active []*anime.Anime
}

func (t Team) String() string {
	a := []string{}
	for _, anime := range t.Active {
		a = append(a, anime.String())
	}
	b := []string{}
	for _, anime := range t.Bench {
		b = append(b, anime.String())
	}
	output := fmt.Sprintf("\nActive: %s\nBench: %s\n", strings.Join(a, ", "), strings.Join(b, ", "))
	return output
}
