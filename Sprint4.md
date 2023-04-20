WORK COMPLETED: Fixed post liking/unliking and following.  Added deleting posts to backend, reorganized backend for better readability, Added  branches to continue work on future features such as blocking users.

BACKEND: API documentation is avalible here, dowload HTML file and open in browser: https://drive.google.com/file/d/1RlMxF589X-dUCejbfKYFW6dImKrfXV7b/view?usp=sharing
Tests: 
   TestHashStr - tests hashing strings with Bcrypt ||  
 TestGetUsers - tests GET for getting users  ||  
 TestCreateUser - tests POST for creation of a new user  ||  
 TestCreatePost - tests POST for creating a new post  ||  
 

FRONTEND:
WORK COMPLETED: 
   Added User Verification:
      checks that user email is in our database before logging in. 
      Users now have password requirments 
      Added check to make sure that email is vaild when creating an account. 
      Fixed post sizing issue.
      Standarized post formating so that heart and image size stays the same.
      Added file type checks to only accept jpg. -- solves "cannot load" error.
      Users now have to add comments on their post before psoting.
      Profile card: 
         shows user's name 
         follower count
         following count
      Allows the user to search for other followers
      Autofills follower search for users with matching tag
      
      
