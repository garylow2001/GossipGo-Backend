package seed

import (
	"github.com/garylow2001/GossipGo-Backend/types"
)

var SeededUsers = []types.User{
	{
		ID:       1,
		Username: "user1",
		Password: "password1",
	},
	{
		ID:       2,
		Username: "user2",
		Password: "password2",
	},
}

var SeededThreads = []types.Thread{
	{
		ID:     1,
		Title:  "Thread 1",
		Body:   "This is the first thread",
		Author: SeededUsers[0],
	},
	{
		ID:     2,
		Title:  "Thread 2",
		Body:   "This is the second thread",
		Author: SeededUsers[1],
	},
}

var SeededComments = []types.Comment{
	{
		ID:     1,
		Body:   "This is a comment",
		Author: SeededUsers[0],
		Thread: SeededThreads[0],
	},
	{
		ID:     2,
		Body:   "This is another comment",
		Author: SeededUsers[1],
		Thread: SeededThreads[1],
	},
}
