package github

type CreateRepoRequestData struct {
	Name 		string 	`json:"name"`
	Description string	`json:"description"`
	Homepage 	string 	`json:"homepage"`
	Private		bool 	`json:"private"`
	HasIssues	bool 	`json:"has_issues"`
	HasProjects bool 	`json:"has_projects"`
	HasWiki		bool 	`json:"has_wiki"`
}

type CreateRepoResponseData struct {
	Id 			int64 			`json:"id"`
	Name 		string			`json:"name"`
	FullName	string 			`json:"full_name"`
	Owner		RepoOwner 		`json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

type RepoOwner struct {
	Id 		int64 	`json:"id"`
	Login 	string 	`json:"login"`
	Url 	string 	`json:"url"`
	HtmlUrl string	`json:"html_url"`
}

type RepoPermissions struct {
	Admin	bool `json:"admin"`
	Push	bool `json:"push"`
	Pull	bool `json:"pull"`
}