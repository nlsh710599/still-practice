package route

type getMemeCoinParams struct {
	ID uint `uri:"id" binding:"required"`
}

type CreateMemeCoinRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type updateMemeCoinParams struct {
	ID uint `uri:"id" binding:"required"`
}

type UpdateMemeCoinRequest struct {
	Description string `json:"description" binding:"required"`
}
