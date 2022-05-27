package main

/*
Create a web server with the following endpoints:
- POST /newGame ->
	- Creates a new game and returns the game id
- POST /joinGame/:gameID ->
	- recieves a game id as uri parameter
	- Creates a new player and adds it to the game
    - creates a cookie with the game id and the created player id
    - returns a message that the player was added to the game

- POST /add -> adds the player with the given token to the game with the given ID
- GET /list -> returns a list of all players in the game with the given ID
- DELETE /remove -> removes the player with the given token from the game with the given ID

entities are:
- Player
	- Name
	- GameID

- Game
	- Players

persistence:
- database: postgres
- ORM: gorm
- connection: postgresql://localhost:5432/postgres?sslmode=disable

server:
- port: 8080
- host: localhost
- mode: development
- libraries:
	- mux
	- cors
	- cookie-session
*/

import (
	"fmt"
	"net/http"
	"os"
	"server/controllers"
	"server/db"
	"server/router"
)

func main() {
	database := db.InitializeDB()
	r := router.SetupRouter()
	port := os.Getenv("PORT")

	controllers.InitGameController(database)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":" + port, r)
}
