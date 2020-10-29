package documents

type DocumentsResponse struct {
	ID       string `json:"id" db:"user_id"`
	DocType  string `json:"doc_type" db:"doc_type"`
	DocName  string `json:"doc_name" db:"document_name"`
	Document string `json:"document" db:"document"`
}
