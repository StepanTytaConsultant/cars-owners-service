package urlvals

type PaginationParams struct {
	Limit  *int32 `page:"limit"`
	Offset *int32 `page:"offset"`
}
