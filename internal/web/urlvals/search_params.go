package urlvals

type SearchCarsParams struct {
	FirstName *string `search:"first_name"`
	LastName  *string `search:"last_name"`
}
