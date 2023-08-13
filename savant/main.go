package savant

import (
	"fmt"
	"io"
	"net/http"
)

const (
	BASE_URL       = "https://baseballsavant.mlb.com/statcast_search/csv"
	PARAMETERS     = "all=true&hfPT=&hfAB=&hfGT=R%7C&hfPR=&hfZ=&stadium=&hfBBL=&hfNewZones=&hfPull=&hfC=&hfSit=&hfOuts=&opponent=&pitcher_throws=&batter_stands=&hfSA=&hfInfield=&team=&position=&hfOutfield=&hfRO=&home_road=&hfFlag=&hfBBT=&metric_1=&hfInn=&min_pitches=0&min_results=0&group_by=name&sort_col=pitches&player_event_sort=api_p_release_speed&sort_order=desc&min_pas=0&type=details&"
	QUERY_FORMAT   = "player_type=%s&game_date_gt=%s&game_date_lt=%s&hfSea=%d"
	FILE_FORMAT    = "%s/%s.csv"
	GAME_DT_FORMAT = "2006-01-02"
)

func Query(form Form) string {
	params := fmt.Sprintf(QUERY_FORMAT,
		form.PlayerType,
		form.GameDate.Format(GAME_DT_FORMAT),
		form.GameDate.Format(GAME_DT_FORMAT),
		form.Season,
	)
	return fmt.Sprintf("%s?%s&%s%s", BASE_URL, PARAMETERS, params, "%7C")
}

func Filename(form Form) string {
	filename := fmt.Sprintf(FILE_FORMAT,
		form.GameDate.Format(GAME_DT_FORMAT),
		form.PlayerType,
	)
	return filename
}

func FetchAndAsString(query string) (string, error) {
	// HTTP GETリクエストを発行
	resp, err := http.Get(query)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスボディを読み込む
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
