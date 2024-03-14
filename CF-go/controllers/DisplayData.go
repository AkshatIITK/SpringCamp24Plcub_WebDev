package controllers

import (
	model "cfapiapp/models"
	"cfapiapp/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ShowActions(c *gin.Context, mongoStore *store.MongoStore) {
	filter := bson.M{}
	AllActions := mongoStore.GetActionsfromDatabase(filter)
	// fmt.Println(AllActions)
	c.IndentedJSON(http.StatusOK, AllActions)
}

func ShowBlogs(c *gin.Context, mongoStore *store.MongoStore) {
	// fmt.Printf("SHowing Blogs")
	filter := bson.M{}
	AllActions := mongoStore.GetActionsfromDatabase(filter)
	// fmt.Println(AllActions)
	groupedBlog := mongoStore.GroupBlog(AllActions)
	c.IndentedJSON(http.StatusOK, groupedBlog)
}

func ShowUserActions(c *gin.Context, mongoStore *store.MongoStore) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NOT Found"})
		return
	}
	// blogid := mongoStore.GetUserASubcribedActions(user)
	// filter := bson.M{
	// 	"subscribedblogsid": bson.M{
	// 		"$in": blogid,
	// 	},
	// 	// "email": user.Email,
	// }

	// c.JSON(http.StatusAccepted, blogid)
	// filter := bson.M{"id": bson.M{"$in": []string{"id1", "id2", "id3"}}}
	// AllActions := mongoStore.GetActionsfromDatabase(filter)
	// groupedBlog := mongoStore.GroupBlog(AllActions)
	// c.IndentedJSON(http.StatusOK, groupedBlog)
	c.IndentedJSON(http.StatusOK, user)
}

func SubsribeHelper(c *gin.Context, mongoStore *store.MongoStore) {
	var UserSubscribe model.UserSubscribe
	if err := c.ShouldBindJSON(&UserSubscribe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NOT Found"})
		return
	}
	var userData model.User
	filter := bson.M{
		"email": UserSubscribe.Email,
	}
	c.JSON(http.StatusAccepted, UserSubscribe)
	_, userData = mongoStore.IsUserExist(filter)
	// c.JSON(http.StatusAccepted, userData)
	mongoStore.Subcribe(userData, UserSubscribe.BlogId)
	c.JSON(http.StatusAccepted, userData)

	// c.IndentedJSON(http.StatusOK, UserSubscribe)

}
func UnSubsribeHelper(c *gin.Context, mongoStore *store.MongoStore) {
	var UserSubscribe model.UserSubscribe
	if err := c.ShouldBindJSON(&UserSubscribe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NOT Found"})
		return
	}
	var userData model.User
	filter := bson.M{
		"email": UserSubscribe.Email,
	}
	c.JSON(http.StatusAccepted, UserSubscribe)
	_, userData = mongoStore.IsUserExist(filter)
	// c.JSON(http.StatusAccepted, userData)
	mongoStore.UnSubcribe(userData, UserSubscribe.BlogId)
	c.JSON(http.StatusAccepted, userData.SubscribedBlogsId)

	// c.IndentedJSON(http.StatusOK, UserSubscribe)

}
