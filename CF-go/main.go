package main

import (
	Route "cfapiapp/routes"
	"cfapiapp/store"
	"cfapiapp/worker"
	"sync"
)

func main() {
	var w sync.WaitGroup
	w.Add(2)

	go func() {
		Route.Route()
		w.Done()
	}()
	go func() {
		worker.WorkerFn(&store.MongoStore{})
		w.Done()
	}()

	w.Wait()

	// workermain(&store.MongoStore{})

}

// func workermain(mongoStore *store.MongoStore) {
// 	mongoStore.OpenConnectionWithMongoDB()
// 	Actions, _ := cfapi.Cfapicall()
// 	// fmt.Println(Actions)
// 	mongoStore.StoreRecentActionsInTheDatabase(Actions)

// }
