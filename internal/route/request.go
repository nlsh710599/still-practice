package route

type getMemeCoinParams struct {
	ID uint `uri:"id" binding:"required"`
}
