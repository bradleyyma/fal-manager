package model

type PointCriteria struct {
	Watching    float32 `json: "watching"`
	Discussions int     `json: "discussions"`
	Score       int     `json: "score"`
}
