# GO-BLOG blogging platform RESTful API service

Just another fun project from me 👻... This is a simple REST API service built for a blogging platform where a anybody can sign-up, and start their own blog by creating a post. The two major models on this platform are the Users and Blog posts... Feel free to check it out and of course, play around it!

### Getting started
To run this API locally, you must have Go installed on your machine, PostgresSQL is the database used, but you may want to try out other SQL DBMSs. To do this, you would need to tweak the Init() method present in the controllers/base.go file. Not forgetting to install the required gorm dialect package for  whichever DBMS you'd be using. Follow the steps below to get started 👇

#### Run the following commands
```
git clone https://github.com/Tee-Stark/go-blog.git
cd go-blog
go mod download
```

#### Set Environment Variables
You may now go ahead to set your environment variables by copying the content of the `envsample` file into a .env file in the root directory. Ensure the content is set to your local database configurations ⚠ .
`
#### Create Local Database
Create a new database with the name and owner as specified in your `DB_NAME` & `DB_USER` variables. 

#### Start the Application
Once done, you may now start the Go application by running the command below in the project's root directory:
`go run main.go`

The above command will run auto-migrations, and get the server running on your specified port.

#### Making Requests (API docs still cooking 🙃)
kindly bear with me...

while I complete the docs and add tests...you can always checkout the code and know your way around, it's very understandable I promise... 💓
### Enjoy !!!