package anime

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/apex/log"
)

const (
	endpoint    = "https://api.myanimelist.net/v2/anime/"
	queryParams = "?fields=id,title,mean,statistics"
)

type Anime struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Score      float32    `json:"mean"`
	Statistics Statistics `json:"statistics"`
}

func (a *Anime) GetInfo(clientID string) error {
	url := fmt.Sprintf("%s%v%s", endpoint, a.Id, queryParams)
	log.Infof("Getting anime info: %s", url)
	client := http.Client{}
	anime, err := getAnime(&client, url, clientID)
	*a = *anime
	return err
}

func (a Anime) String() string {
	return a.Title
}

type Statistics struct {
	Status Status `json:"status"`
}

type Status struct {
	Watching    string `json:"watching"`
	Completed   string `json:"completed"`
	OnHold      string `json:"on_hold"`
	PlanToWatch string `json:"plan_to_watch"`
}

func getAnime(client *http.Client, url, clientID string) (*Anime, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-MAL-CLIENT-ID", clientID)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	anime := &Anime{}
	if resp.StatusCode == http.StatusOK {
		if err := json.Unmarshal(body, anime); err != nil {
			return nil, err
		}
		return anime, nil
	} else {
		return nil, fmt.Errorf("error getting anime details %v", string(body))
	}

}
