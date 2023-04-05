package botAccount

type GetModelShowData struct {
	Model string `json:"model"`
}

type SetModelShowData struct {
	Model     string `json:"model"`
	ModelShow string `json:"model_show"`
}

type GetOnlineClientsData struct {
	NoCache bool `json:"no_cache"`
}

type SetQQProfileData struct {
	Nickname     string `json:"nickname"`
	Company      string `json:"company"`
	Email        string `json:"email"`
	College      string `json:"college"`
	PersonalNote string `json:"personal_note"`
}
