package controllers

import "gorm.io/gorm"

// PlayerController implements the Controller interface

type Player struct {
	*gorm.Model
	Name   string "name"
	GameID uint
}

type PlayerController Controller

func (c PlayerController) GetAll() ([]Player, error) {
	var players []Player
	c.DB.Find(&players)
	return players, nil
}

func (c PlayerController) Get(id interface{}) (Player, error) {
	// get game from database
	var player Player
	c.DB.First(&player, id)
	return player, nil
}

func (c PlayerController) Create(player Player) (Player, error) {
	// create game in database
	var _player = player
	c.DB.Create(&_player)
	return _player, nil
}

func (c PlayerController) Delete(id interface{}) (Player, error) {
	player, _ := c.Get(id)
	c.DB.Delete(&player)
	return player, nil
}

func NewPlayerController(db *gorm.DB) PlayerController {
	db.AutoMigrate(&Player{})
	return PlayerController{DB: db}
}
