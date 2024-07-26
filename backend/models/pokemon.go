package models

type Pokemon struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"unique;not null"`
	HP     int    `gorm:"column:hp;not null"`
	Atk    int    `gorm:"not null"`
	Def    int    `gorm:"not null"`
	AtkSpe int    `gorm:"column:atk_spe;not null"`
	DefSpe int    `gorm:"column:def_spe;not null"`
	Speed  int    `gorm:"not null"`
	Types  []Type `gorm:"many2many:pokemon_type;"`
	Teams  []Team `gorm:"many2many:team_pokemon;"`
}

func (Pokemon) TableName() string {
	return "pokemon"
}

type Type struct {
	ID      uint      `gorm:"primaryKey"`
	Name    string    `gorm:"unique;not null"`
	Color   string    `gorm:"not null"`
	Pokemon []Pokemon `gorm:"many2many:pokemon_type;"`
}

func (Type) TableName() string {
	return "type"
}

type Team struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"not null"`
	Pokemon     []Pokemon `gorm:"many2many:team_pokemon;"`
}

func (Team) TableName() string {
	return "teams"
}

type PokemonType struct {
	ID        uint `gorm:"primaryKey"`
	PokemonID uint `gorm:"column:pokemon_id;not null"`
	TypeID    uint `gorm:"column:type_id;not null"`
}

func (PokemonType) TableName() string {
	return "pokemon_type"
}

type TeamPokemon struct {
	ID        uint `gorm:"primaryKey"`
	TeamID    uint `gorm:"column:team_id;not null"`
	PokemonID uint `gorm:"column:pokemon_id;not null"`
}

func (TeamPokemon) TableName() string {
	return "team_pokemon"
}
