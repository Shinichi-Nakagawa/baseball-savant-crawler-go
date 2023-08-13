package savant

import (
	"testing"
	"time"
)

func TestQuery(t *testing.T) {
	// 正常系テスト（ValidationでOKなデータ）
	gameDate21, _ := time.Parse(GAME_DT_FORMAT, "2021-04-01")
	gameDate22, _ := time.Parse(GAME_DT_FORMAT, "2022-04-01")
	gameDate23, _ := time.Parse(GAME_DT_FORMAT, "2023-04-01")
	var datasets = []Form{
		{Season: 2021, PlayerType: PITCHER, GameDate: gameDate21},
		{Season: 2021, PlayerType: BATTER, GameDate: gameDate21},
		{Season: 2022, PlayerType: PITCHER, GameDate: gameDate22},
		{Season: 2022, PlayerType: BATTER, GameDate: gameDate22},
		{Season: 2023, PlayerType: PITCHER, GameDate: gameDate23},
		{Season: 2023, PlayerType: BATTER, GameDate: gameDate23},
	}
	var results = []string{
		"https://baseballsavant.mlb.com/statcast_search/csv?all=true&hfPT=&hfAB=&hfGT=R%7C&hfPR=&hfZ=&stadium=&hfBBL=&hfNewZones=&hfPull=&hfC=&hfSit=&hfOuts=&opponent=&pitcher_throws=&batter_stands=&hfSA=&hfInfield=&team=&position=&hfOutfield=&hfRO=&home_road=&hfFlag=&hfBBT=&metric_1=&hfInn=&min_pitches=0&min_results=0&group_by=name&sort_col=pitches&player_event_sort=api_p_release_speed&sort_order=desc&min_pas=0&type=details&&player_type=pitcher&game_date_gt=2021-04-01&game_date_lt=2021-04-01&hfSea=2021%7C",
		"https://baseballsavant.mlb.com/statcast_search/csv?all=true&hfPT=&hfAB=&hfGT=R%7C&hfPR=&hfZ=&stadium=&hfBBL=&hfNewZones=&hfPull=&hfC=&hfSit=&hfOuts=&opponent=&pitcher_throws=&batter_stands=&hfSA=&hfInfield=&team=&position=&hfOutfield=&hfRO=&home_road=&hfFlag=&hfBBT=&metric_1=&hfInn=&min_pitches=0&min_results=0&group_by=name&sort_col=pitches&player_event_sort=api_p_release_speed&sort_order=desc&min_pas=0&type=details&&player_type=batter&game_date_gt=2021-04-01&game_date_lt=2021-04-01&hfSea=2021%7C",
		"https://baseballsavant.mlb.com/statcast_search/csv?all=true&hfPT=&hfAB=&hfGT=R%7C&hfPR=&hfZ=&stadium=&hfBBL=&hfNewZones=&hfPull=&hfC=&hfSit=&hfOuts=&opponent=&pitcher_throws=&batter_stands=&hfSA=&hfInfield=&team=&position=&hfOutfield=&hfRO=&home_road=&hfFlag=&hfBBT=&metric_1=&hfInn=&min_pitches=0&min_results=0&group_by=name&sort_col=pitches&player_event_sort=api_p_release_speed&sort_order=desc&min_pas=0&type=details&&player_type=pitcher&game_date_gt=2022-04-01&game_date_lt=2022-04-01&hfSea=2022%7C",
		"https://baseballsavant.mlb.com/statcast_search/csv?all=true&hfPT=&hfAB=&hfGT=R%7C&hfPR=&hfZ=&stadium=&hfBBL=&hfNewZones=&hfPull=&hfC=&hfSit=&hfOuts=&opponent=&pitcher_throws=&batter_stands=&hfSA=&hfInfield=&team=&position=&hfOutfield=&hfRO=&home_road=&hfFlag=&hfBBT=&metric_1=&hfInn=&min_pitches=0&min_results=0&group_by=name&sort_col=pitches&player_event_sort=api_p_release_speed&sort_order=desc&min_pas=0&type=details&&player_type=batter&game_date_gt=2022-04-01&game_date_lt=2022-04-01&hfSea=2022%7C",
		"https://baseballsavant.mlb.com/statcast_search/csv?all=true&hfPT=&hfAB=&hfGT=R%7C&hfPR=&hfZ=&stadium=&hfBBL=&hfNewZones=&hfPull=&hfC=&hfSit=&hfOuts=&opponent=&pitcher_throws=&batter_stands=&hfSA=&hfInfield=&team=&position=&hfOutfield=&hfRO=&home_road=&hfFlag=&hfBBT=&metric_1=&hfInn=&min_pitches=0&min_results=0&group_by=name&sort_col=pitches&player_event_sort=api_p_release_speed&sort_order=desc&min_pas=0&type=details&&player_type=pitcher&game_date_gt=2023-04-01&game_date_lt=2023-04-01&hfSea=2023%7C",
		"https://baseballsavant.mlb.com/statcast_search/csv?all=true&hfPT=&hfAB=&hfGT=R%7C&hfPR=&hfZ=&stadium=&hfBBL=&hfNewZones=&hfPull=&hfC=&hfSit=&hfOuts=&opponent=&pitcher_throws=&batter_stands=&hfSA=&hfInfield=&team=&position=&hfOutfield=&hfRO=&home_road=&hfFlag=&hfBBT=&metric_1=&hfInn=&min_pitches=0&min_results=0&group_by=name&sort_col=pitches&player_event_sort=api_p_release_speed&sort_order=desc&min_pas=0&type=details&&player_type=batter&game_date_gt=2023-04-01&game_date_lt=2023-04-01&hfSea=2023%7C",
	}
	for i, v := range datasets {
		query := Query(v)
		if query != results[i] {
			// エラーになっちゃいけない（ここで引っかかったらアウト）
			t.Errorf("URL Error: season: %v player: %v game_date: %v query:%v", v.Season, v.PlayerType, v.GameDate, query)
		}
	}
}

func TestFilename(t *testing.T) {
	// 正常系テスト（ValidationでOKなデータ）
	gameDate21, _ := time.Parse(GAME_DT_FORMAT, "2021-04-01")
	gameDate22, _ := time.Parse(GAME_DT_FORMAT, "2022-05-31")
	gameDate23, _ := time.Parse(GAME_DT_FORMAT, "2023-10-15")
	var datasets = []Form{
		{Season: 2021, PlayerType: PITCHER, GameDate: gameDate21},
		{Season: 2021, PlayerType: BATTER, GameDate: gameDate21},
		{Season: 2022, PlayerType: PITCHER, GameDate: gameDate22},
		{Season: 2022, PlayerType: BATTER, GameDate: gameDate22},
		{Season: 2023, PlayerType: PITCHER, GameDate: gameDate23},
		{Season: 2023, PlayerType: BATTER, GameDate: gameDate23},
	}
	var results = []string{
		"2021-04-01/pitcher.csv",
		"2021-04-01/batter.csv",
		"2022-05-31/pitcher.csv",
		"2022-05-31/batter.csv",
		"2023-10-15/pitcher.csv",
		"2023-10-15/batter.csv",
	}
	for i, v := range datasets {
		query := Filename(v)
		if query != results[i] {
			// エラーになっちゃいけない（ここで引っかかったらアウト）
			t.Errorf("URL Error: season: %v player: %v game_date: %v query:%v", v.Season, v.PlayerType, v.GameDate, query)
		}
	}
}
