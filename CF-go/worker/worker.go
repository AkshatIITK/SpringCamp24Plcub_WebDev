package worker

import (
	"cfapiapp/cfapi"
	"cfapiapp/store"
	"fmt"
	"time"
)

func WorkerFn(mongoStore *store.MongoStore) {
	for {
		fmt.Println("worker started:")
		mongoStore.OpenConnectionWithMongoDB()
		Actions, _ := cfapi.Cfapicall()
		fmt.Printf("Fetched Recent Actions\n")
		// mongoStore.GetUserASubcribedActions()
		// filter := bson.M{}
		// AllActions := mongoStore.GetActionsfromDatabase(filter)
		// fmt.Println(AllActions)
		// a := mongoStore.GroupBlog(AllActions)

		// fmt.Println(a)
		// fmt.Println(Actions)
		mongoStore.StoreRecentActionsInTheDatabase(Actions)
		fmt.Printf("Worker Stopped for 10Seconds")
		time.Sleep(100 * time.Second)

	}
}
