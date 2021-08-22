package query

type UserGameStateQuery struct {
	GamesPlayed int32 `json:"gamesPlayed"`
	Score       int64  `json:"score"`
}
