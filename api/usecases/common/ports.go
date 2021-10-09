package common

type OutputPort struct {
	StatusCode int16       `json:"-"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

// export interface ResultQueryDto {
// 	data: any[];
// 	totalDocs: number;
// 	limit: number;
// 	page: number;
// 	totalPages: number;
// 	hasPrevPage: boolean | null;
// 	hasNextPage: boolean | null;
// 	prevPage?: number;
// 	nextPage?: number;
//   }
