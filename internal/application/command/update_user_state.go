package command

type UpdateUserState struct {
	GamesPlayed uint8 `json:"gamesPlayed"`
	Score       uint  `json:"score"`
}
