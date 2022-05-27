package controllers

import (
	"fmt"
	"gorm.io/gorm"
)

// GameController implements the Controller interface

type Game struct {
	*gorm.Model
	Players []Player
}

type GameController struct {
	DB *gorm.DB
	PlayerController PlayerController
}

var GameControllerInstance GameController

func (c GameController) GetAll() ([]Game, error) {
	// get all games from database
	var games []Game
	c.DB.Find(&games)
	return games, nil
}

func (c GameController) Get(id interface{}) (Game, error) {
	// get game from database
	var game Game
	// find game with association to players
	c.DB.First(&game, id)
	err := c.DB.Model(&game).Association("Players").Find(&game.Players)
	if err != nil {
		fmt.Print(err)
	}
	return game, nil
}

func (c GameController) Create() (Game, error) {
	// create game in database
	var game Game
	c.DB.Create(&game)
	game.Players = make([]Player, 0)
	return game, nil
}

func (c GameController) Update(id interface{}, element interface{}) error {
	// update game in database
	game, err := c.Get(id)
	if err != nil {
		return err
	}

	c.DB.Model(&game).Updates(element)
	return nil
}

func (c GameController) Delete(id interface{}) error {
	// delete game from database
	game, err := c.Get(id)

	if err != nil {
		return err
	}

	c.DB.Delete(&game)
	return nil
}

func (c GameController) AddNewPlayerToGame(gameID interface{}, player Player) (Game, Player, error) {
	game, err := c.Get(gameID)

	//create player
	player.GameID = game.ID
	_player, _ := c.PlayerController.Create(player)

	if err != nil {
		return game, _player, err
	}

	c.DB.Model(&game).Association("players").Append(_player)

	game.Players = append(game.Players, _player)

	return game, _player, nil
}

func (c GameController) RemovePlayerFromGame(gameID interface{}, playerID interface{}) (Game, error) {
	c.PlayerController.Delete(playerID)
	game, _ := c.Get(gameID)

	return game, nil
}

func InitGameController(db *gorm.DB) {
	db.AutoMigrate(&Game{})
	GameControllerInstance = GameController{
		DB: db,
		PlayerController: NewPlayerController(db),
	}
}
