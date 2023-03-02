Backend:

API:

We have a fuctional API for creating new users via a "sign up" form, and log in for existing users.
Our API stores user data such as email, name, hashed password, auto-generated ID number, and a following list of followed users IDs, and posts theyve made.
We also have enabled content posting, allowing users to post text and images to their profiles


Backend unit tests: 

Test HashStr function to ensure hashing works, and a hashed password can be compared with user input
Test Constains: tests the Contains function which is used to find existing user IDs within a slice, to ensure duplicate followers dont occur
Test Remove From Slice: tests the RemoveFromSlice function which is used to remove IDs from a user's following list
Test Get Users:  tests GET calls for retrieving users from the DB
Test CreateUsers:  tests POST calls for creation of new users


Frontend tests:
Test that the site pages are up and running
Tests button clicks on pages and ensures connectivity