package types

type User struct {
	ID       string    `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	Threads  []Thread  `json:"threads,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

type Thread struct {
	ID       string    `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Body     string    `json:"body,omitempty"`
	Author   User      `json:"author,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID     string `json:"id,omitempty"`
	Body   string `json:"body,omitempty"`
	Author User   `json:"author,omitempty"`
	Thread Thread `json:"thread,omitempty"`
}
