# foodhub
CEN3031 - group 1, foodhub project 

App name: FoodHub

Overview:
Users would be able to post pictures and comments of and about food to their profile feed, which would then be seen through a “heat map” where a more intensive color would indicate a higher user interaction. 

Utilities: 
Users would be able to “follow” other users to see what they are posting and where they are eating

Business: 
Would be able to upload their own profiles with ads and deals to engage with singular users. 

Members: 
Front-End:
Thomas McMullen, Agustin Giraldo 

Back End: 
Larry Mason, Nicholas Orth


# How to Run

### Requirements
- 64bit C/C++ Compiler
- Go 1.16 or later
- Node.js 14 or later
- Docker installed


### Backend
- You need redis installed, run ```docker pull redis``` and ```docker run -d --name my-redis-container -p 6379:6379 redis```
- Open a terminal and navigate to the backend folder
- Run `go run .` to start the backend server
- The backend server will be running on port 7100
- To access docs, go to http://localhost:7100/api/docs

### Frontend
- Open a terminal and navigate to the frontend folder
- Run `npm install` to install dependencies
- Run `npm run build` to build the release version
- Run `npm run start` to start the frontend server
