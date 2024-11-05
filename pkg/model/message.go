package model

type Message struct {
	PlayerUUID string `json:"player_uuid"`
	PlayerName string `json:"player_name"`
	Content    string `json:"content"`
	Timestamp  uint64 `json:"timestamp"`
}
