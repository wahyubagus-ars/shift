package request

type CreateWorkspace struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	ProfilePicture string `json:"profile_picture"`
}
