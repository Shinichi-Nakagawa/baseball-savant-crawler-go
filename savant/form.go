package savant

import (
	"encoding/json"
	"time"
)

type PlayerType string

const (
	PITCHER PlayerType = "pitcher"
	BATTER  PlayerType = "batter"
)

func CreateForm(value string) (Form, error) {
	var form Form
	if err := json.Unmarshal([]byte(value), &form); err != nil {
		return form, err

	}
	return form, nil
}

type Form struct {
	Season     int        `validate:"required, min=2015,max=2999" json:"season"`
	PlayerType PlayerType `validate:"required" json:"player_type"`
	GameDate   time.Time  `validate:"required" json:"game_date"`
}
