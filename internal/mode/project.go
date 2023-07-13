package mode

type ProCreate struct {
	Name string `json:"name" validate:"min=1"`
}
