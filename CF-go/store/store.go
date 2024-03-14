package store

import (
	model "cfapiapp/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Collection     *mongo.Collection
	UserCollection *mongo.Collection
}

// con_STR := "mongodb+srv://admin:ak1234@cluster0.ovujcuz.mongodb.net/?retryWrites=true&w=majority"
var CON_STR string = "mongodb://localhost:27017"

func (m *MongoStore) OpenConnectionWithMongoDB() {

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CON_STR).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	m.Collection = client.Database("db").Collection("cfapi")

	//	defer func() {
	//		if err = client.Disconnect(context.TODO()); err != nil {
	//			panic(err)
	//		}
	//	}()
}

func (m *MongoStore) StoreRecentActionsInTheDatabase(RecentAction []model.Result) error {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"timeseconds", -1}})
	cursor, err := m.Collection.Find(context.TODO(), filter, opts)
	var results []model.TimeStamp

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	// max := results[0].TimeSeconds
	// fmt.Println(reflect.TypeOf(max))
	var max int64
	if len(results) > 0 {
		max = results[0].TimeSeconds

		for _, result := range results {
			if max < result.TimeSeconds {
				max = result.TimeSeconds
				fmt.Printf("Got TimeStamp")
			}
		}
	} else {
		// Handle the case when results is empty, for example:
		fmt.Println("No results found")
	}
	var RecentActionInterface []interface{}

	for _, Action := range RecentAction {
		if Action.TimeSeconds > max {
			RecentActionInterface = append(RecentActionInterface, Action)
			fmt.Println("New Record Inserted")

		}

	}

	_, errInsert := m.Collection.InsertMany(context.TODO(), RecentActionInterface)

	if errInsert != nil {
		return fmt.Errorf("error inserting documents: %v", err)

	}
	// fmt.Println(result)
	// res, _ := bson.MarshalExtJSON(result, false, false)
	// fmt.Println(string(res))

	return nil

}
func (m *MongoStore) GroupBlog(actions []model.Result) map[int][]model.SubscribedBlogs {
	groupedData := make(map[int][]model.SubscribedBlogs)

	for _, action := range actions {
		blogID := action.BlogEntry.ID
		comment := action.Comment

		// Check if the blog entry is already present in the map
		if _, ok := groupedData[blogID]; !ok {
			// If not, create a new entry
			groupedData[blogID] = []model.SubscribedBlogs{{Blogs: action.BlogEntry, Comments: []model.Comment{comment}}}
		} else {
			// If yes, append the comment to the existing entry
			var found bool
			for i, entry := range groupedData[blogID] {
				if entry.Blogs.ID == blogID {
					found = true
					// Check if the comment already exists
					for _, c := range entry.Comments {
						if c.ID == comment.ID {
							found = false
							break
						}
					}
					// If the comment is not found, append it
					if found {
						groupedData[blogID][i].Comments = append(entry.Comments, comment)
					}
					break
				}
			}
			// If the blog entry is not found, create a new entry
			if !found {
				groupedData[blogID] = append(groupedData[blogID], model.SubscribedBlogs{Blogs: action.BlogEntry, Comments: []model.Comment{comment}})
			}
		}
	}

	return groupedData
}

// working groupblog
// func (m *MongoStore) GroupBlog(Actions []model.Result) map[int][]model.SubscribedBlogs {
// 	groupedAction := make(map[int][]model.SubscribedBlogs)

// 	for _, Action := range Actions {
// 		BlogID := Action.BlogEntry.ID
// 		comment := Action.Comment

// 		// Check if the blog entry is already present in the map
// 		if _, ok := groupedAction[BlogID]; !ok {
// 			// If not, create a new entry
// 			groupedAction[BlogID] = []model.SubscribedBlogs{
// 				{
// 					Blogs:    Action.BlogEntry,
// 					Comments: []model.Comment{comment},
// 				},
// 			}
// 		} else {
// 			// If yes, check if the comment is already present for that blog entry
// 			found := false
// 			for i, entry := range groupedAction[BlogID] {
// 				if entry.Blogs.ID == BlogID {
// 					// Blog entry found, check if the comment is already present
// 					for _, c := range entry.Comments {
// 						if c.ID == comment.ID {
// 							found = true
// 							break
// 						}
// 					}

// 					// If not found, add the comment
// 					if !found {
// 						groupedAction[BlogID][i].AddComment(comment)
// 					}

// 					break
// 				}
// 			}

// 			// If the blog entry is not found, create a new entry
// 			if !found {
// 				groupedAction[BlogID] = append(groupedAction[BlogID], model.SubscribedBlogs{
// 					Blogs:    Action.BlogEntry,
// 					Comments: []model.Comment{comment},
// 				})
// 			}
// 		}
// 	}

// 	return groupedAction
// }

func (m *MongoStore) GetActionsfromDatabase(filter bson.M) []model.Result {
	opts := options.Find().SetSort(bson.D{{"timeseconds", -1}})
	cursor, err := m.Collection.Find(context.TODO(), filter, opts)
	var Actions []model.Result

	if err = cursor.All(context.TODO(), &Actions); err != nil {
		panic(err)
	}
	// fmt.Println(Actions)
	return Actions
}

func (m *MongoStore) GetUserASubcribedActions(User model.User) []int {
	// m.GroupBlog(UserActions)
	// GetActionsfromDatabase(filter bson.M)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CON_STR).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	m.UserCollection = client.Database("db").Collection("user")
	filter := bson.M{
		"email": User.Email,
	}
	var userDB model.User
	err = m.UserCollection.FindOne(context.TODO(), filter).Decode(&userDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No user //given credentials
			return nil
		}
		panic(err)
	}
	return userDB.SubscribedBlogsId

}

func (m *MongoStore) Subcribe(user model.User, blogid int) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CON_STR).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	m.UserCollection = client.Database("db").Collection("user")
	exists := false
	if len(user.SubscribedBlogsId) > 0 {
		for _, id := range user.SubscribedBlogsId {
			if id == blogid {
				exists = true
				break
			}
		}
	}
	if !exists {
		filter := bson.M{
			"email": user.Email,
		}
		update := bson.M{
			"$push": bson.M{"subscribedblogsid": blogid},
		}
		_, err = m.UserCollection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
	}

	return

}

func (m *MongoStore) UnSubcribe(user model.User, blogid int) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CON_STR).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	m.UserCollection = client.Database("db").Collection("user")
	// exists := false
	if len(user.SubscribedBlogsId) > 0 {
		for _, id := range user.SubscribedBlogsId {
			if id == blogid {
				// exists = true
				filter := bson.M{
					"email": user.Email,
				}
				update := bson.M{
					"$pull": bson.M{"subscribedblogsid": blogid},
				}
				_, err = m.UserCollection.UpdateOne(context.TODO(), filter, update)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
		}
	}
	// if !exists {

	// }

	return

}
func (m *MongoStore) InsertUserData(user model.User) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CON_STR).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	user.SubscribedBlogsId = []int{}
	m.UserCollection = client.Database("db").Collection("user")
	_, err = m.UserCollection.InsertOne(context.TODO(), user)
	return err
}

// func (m *MongoStore) UpdateUser(user model.User) (bool, model.User) {
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI(CON_STR).SetServerAPIOptions(serverAPI)
// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), opts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	m.UserCollection = client.Database("db").Collection("user")

// 	filter := bson.M{
// 		// "email":    user.Email,
// 		"username": user.Username,
// 	}
// 	var newBlogID = []int{1, 3}
// 	update := bson.M{"$push": bson.M{"subscribedblogsid": bson.M{"$each": newBlogID}}}

// 	var userDB model.User
// 	_, err = m.UserCollection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// No user found with the given credentials
// 			return false, user
// 		}
// 		panic(err)
// 	}
// 	return true, userDB
// }

func (m *MongoStore) IsUserExist(filter bson.M) (bool, model.User) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(CON_STR).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	m.UserCollection = client.Database("db").Collection("user")
	var userDB model.User
	err = m.UserCollection.FindOne(context.TODO(), filter).Decode(&userDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No user //given credentials
			return false, model.User{}
		}
		panic(err)
	}
	return true, userDB
}
