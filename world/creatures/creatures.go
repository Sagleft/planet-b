package creatures

// Creature - strange biorobots
type Creature struct {
	Eyes []CreatureEye
	Legs []CreatureLeg
}

// CreatureEye - pieces for the perception of the visible spectrum of radiation
type CreatureEye struct{}

// CreatureLeg - primitive things for movement in space
type CreatureLeg struct{}
