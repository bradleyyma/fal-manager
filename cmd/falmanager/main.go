package main

import (
	"github.com/bradleyyma/falmanager/pkg/fal"
)

func main() {
	clientId := "58f7627fe0ae31c2b7391c641a77eea5"
	fal.EvaluateFAL("team.json", clientId)
}
