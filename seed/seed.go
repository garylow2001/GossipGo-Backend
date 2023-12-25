package seed

import (
	"github.com/garylow2001/GossipGo-Backend/models"
)

var SeededUsers = []models.User{
	{
		// ID:       1,
		Username: "user1",
	},
	{
		// ID:       2,
		Username: "user2",
	},
}

// var SeededThreads = []models.Thread{
// 	{
// 		ID:     1,
// 		Title:  "Thread 1",
// 		Body:   "This is the first thread",
// 		Author: SeededUsers[0],
// 	},
// 	{
// 		ID:     2,
// 		Title:  "Thread 2",
// 		Body:   "This is the second thread",
// 		Author: SeededUsers[1],
// 	},
// }

// var SeededComments = []models.Comment{
// 	{
// 		ID:     1,
// 		Body:   "This is a comment",
// 		Author: SeededUsers[0],
// 		Thread: SeededThreads[0],
// 	},
// 	{
// 		ID:     2,
// 		Body:   "This is another comment",
// 		Author: SeededUsers[1],
// 		Thread: SeededThreads[1],
// 	},
// }
