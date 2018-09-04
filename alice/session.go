package alice

type Session struct {
	New bool `json:"new"`
	MessageId int `json:"message_id"`
	SessionId string `json:"session_id"`
	SkillId string `json:"skill_id"`
	UserId string `json:"user_id"`
}
