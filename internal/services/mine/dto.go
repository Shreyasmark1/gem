package mine

type SaveMine struct {
	MineID   string `json:"mineId" validate:"omitempty,uuid"`
	MineName string `json:"mineName" validate:"required,min=3"`
}
