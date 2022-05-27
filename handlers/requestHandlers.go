package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"server/controllers"
	"server/utils"
)

var gameController = &controllers.GameControllerInstance
var store = &utils.StoreInstance

func NewGame(w http.ResponseWriter, r *http.Request) {
	game, _ := gameController.Create()
	store.SetCookie(w, r, "gameID", game.ID)
	utils.JsonResponse(w, game)
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	gameID := mux.Vars(r)["gameID"]

	var player controllers.Player
	json.NewDecoder(r.Body).Decode(&player)

	//add the player to the game
	game, player, _ := gameController.AddNewPlayerToGame(gameID, player)

	//setup cookies
	store.SetCookie(w, r, "playerID", player.ID)
	store.SetCookie(w, r, "gameID", game.ID)

	utils.JsonResponse(w, game)
}

func List(w http.ResponseWriter, r *http.Request) {
	gameID := store.GetCookie(r, "gameID")
	var game, _ = gameController.Get(gameID)
	utils.JsonResponse(w, game)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	playerID := store.GetCookie(r, "playerID")
	gameID := store.GetCookie(r, "gameID")

	game, _ := gameController.RemovePlayerFromGame(gameID, playerID)

	utils.JsonResponse(w, game)
}
