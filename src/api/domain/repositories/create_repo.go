package repositories

type CreateRepoRequest struct {
	Name		string		`json:"name"`
}

type CreateRepoResponse struct {
	Id			int64 		`json:"id"`
	Name		string		`json:"name"`
	Owner		string		`json:"owner"`
}