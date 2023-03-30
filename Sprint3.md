WORK COMPLETED
began work to implement user blocking

BACKEND 
documentation:
BlockUser: will either block or unblock a user depending if target user's ID is alread listed as blocked.  To use, input the user and the target ID, and the updated
blocked list will be returned as a slice of uints

router.GET("/ping", func(c *gin.Context) : Sends notification to user

router.GET("/user/me", func(c *gin.Context)  : loads user's profile information from db and sends name, email, and slcie of posts to frontend

router.GET("/user/posts"  : Sends slice of user's posts to front end

router.POST("/auth/login" : takes in user input at query, compares user email and password with stored email and hashed password.  If authentication is successful, generate a token.  Pass token to frontend

router.POST("/auth/signup"  :  Takes in user's name, email, and password.  hashes password then passes info to DB

router.POST("/post/create"  : Takes a user inputted caption and photo, creates a new url for the post, and saves to DB, and passes everything to frontend

router.POST("/post/like/:id"  :  When a user presses "like" check to see if post has already been liked by user, if so, remove their like, otherwise add user ID to post's "liked" slice.  send amount of likes to front end

router.GET(":/feed"  :  gets user's posts from DB and sorts by date posted, sends posts to front end

router.POST("/user/follow/:id"  :  when a user presses "follow" on a user, first check if user exists, if not send StatusNotFound error, if so checks if user is already following target user, if so send StatusBadRequest, else if DB update fails send StatusInternalServerError, else update following list and send success message to front end

tests:




TestBlockUser:  Creates a user and blocks an ID number, then unblocks it to test if both functionalities are working

FRONTEND: 
TO DO: NAV BAR
LOG IN

complete style change added highlighting buttons giving a more user friendly feel


updates to the login functinallity; users are now able to create an account 


once account is created users is able to upload an iamge file where they can later view in their feed and profile tabs respectivly


NAV bar was add for easier organization with Profile, Feed, and new as labels


more functinally needs to be added to organize post current system inplaces gird. 

additional tets have been added

