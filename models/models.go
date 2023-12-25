package models

type Thread struct {
	ID       int       `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Body     string    `json:"body,omitempty"`
	Author   User      `json:"author,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID     int    `json:"id,omitempty"`
	Body   string `json:"body,omitempty"`
	Author User   `json:"author,omitempty"`
	Thread Thread `json:"thread,omitempty"`
}
