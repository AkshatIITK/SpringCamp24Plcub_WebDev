package Route

import (
	"cfapiapp/cfapi"
	"cfapiapp/controllers"
	model "cfapiapp/models"
	"cfapiapp/store"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Route() {
	fmt.Printf("Server Connected")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // List of allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable sending credentials
	}))

	// Initialize MongoStore
	mongoStore := &store.MongoStore{}
	mongoStore.OpenConnectionWithMongoDB()

	// mongoStore.OpenConnectionWithMongoDB()

	router.GET("/RecentActions", func(c *gin.Context) {
		Actions, _ := cfapi.Cfapicall()
		c.IndentedJSON(http.StatusOK, Actions)
	})
	router.GET("/Action", func(c *gin.Context) {
		controllers.ShowActions(c, mongoStore)
	})
	router.GET("/GroupedAction", func(c *gin.Context) {
		controllers.ShowBlogs(c, mongoStore)
	})

	router.POST("/user/action", func(c *gin.Context) {
		controllers.ShowUserActions(c, mongoStore)
	})
	router.POST("/register", func(ctx *gin.Context) {
		controllers.RegisterHelper(ctx, mongoStore)
	})
	// router.POST("/user/subscribe", func(c *gin.Context) {
	// 	var user model.User
	// 	if err := c.ShouldBindJSON(&user); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// 	// mongoStore.UpdateUser(user)
	// 	c.IndentedJSON(http.StatusAccepted, gin.H{"Subsribe": "Working on it"})
	// })

	// Login route
	router.POST("/login", func(c *gin.Context) {
		controllers.LoginHelper(c, mongoStore)
	})
	router.GET("/user", func(c *gin.Context) {
		controllers.GetCookieHandler(c)
		controllers.UserHelper(c, mongoStore)

	})

	router.POST("/SubcribedBlogs", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		blogid := mongoStore.GetUserASubcribedActions(user)
		// filter := bson.M{
		// 	// "subscribedblogsid": bson.M{
		// 	// 	"$in": blogid,
		// 	// },
		// 	"email": user.Email,
		// }
		// var userData model.User
		// var isUser bool
		// c.JSON(http.StatusAccepted, blogid)
		filter := bson.M{
			"blogentry.id": bson.M{
				"$in": blogid,
			},
		}

		action := mongoStore.GetActionsfromDatabase(filter)
		groupedBlog := mongoStore.GroupBlog(action)
		c.JSON(http.StatusAccepted, groupedBlog)

	})
	router.POST("/subscribe", func(ctx *gin.Context) {
		controllers.SubsribeHelper(ctx, mongoStore)
	})
	router.POST("/unsubscribe", func(ctx *gin.Context) {
		controllers.UnSubsribeHelper(ctx, mongoStore)
	})
	router.POST("/logout", func(c *gin.Context) {
		var secretKey = []byte("secret-key")
		claims := jwt.MapClaims{}
		claims["exp"] = time.Now().Add(-time.Hour * 24 * 30).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(secretKey)
		c.SetCookie("jwt", tokenString, 3600, "/", "", false, true)
	})
	router.Run(":8080")
}
