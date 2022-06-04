package entities

type PaginateResult struct {
	Docs        any   `json:"docs"`
	TotalDocs   int64 `json:"total_docs,omitempty"`
	Limit       int64 `json:"limit,omitempty"`
	Page        int64 `json:"page,omitempty"`
	TotalPages  int   `json:"total_pages,omitempty"`
	HasPrevPage bool  `json:"has_prev_page"`
	HasNextPage bool  `json:"has_next_page"`
	PrevPage    int64 `json:"prev_page,omitempty"`
	NextPage    int64 `json:"next_page,omitempty"`
}
