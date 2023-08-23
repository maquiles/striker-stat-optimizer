package main

// Body Parts ======================================================================================================
var HEAD string = "head"
var ARMS string = "arms"
var BODY string = "body"
var LEGS string = "legs"

var ALL_BODY_PARTS []string = []string{HEAD, ARMS, BODY, LEGS}

// Stats ======================================================================================================
var STRENGTH string = "strength"
var SPEED string = "speed"
var SHOOTING string = "shooting"
var PASSING string = "passing"
var TECHNIQUE string = "technique"

var ALL_STATS []string = []string{STRENGTH, SPEED, SHOOTING, PASSING, TECHNIQUE}

// MUSCLE GEAR SET ======================================================================================================
var MUSCLE_HELMET Gear = Gear{
	Name:     "Muscle Helmet",
	BodyPart: HEAD,
	StatsEffect: Stats{
		Strength:  2,
		Technique: -2,
	},
}

var MUSCLE_GAUNTLETS Gear = Gear{
	Name:     "Muscle Gauntlets",
	BodyPart: ARMS,
	StatsEffect: Stats{
		Strength: 2,
		Passing:  -2,
	},
}

var MUSCLE_CHEST Gear = Gear{
	Name:     "Muscle Chest",
	BodyPart: BODY,
	StatsEffect: Stats{
		Strength: 2,
		Shooting: -2,
	},
}

var MUSCLE_BOOTS Gear = Gear{
	Name:     "Muscle Boots",
	BodyPart: LEGS,
	StatsEffect: Stats{
		Strength: +2,
		Speed:    -2,
	},
}

// TURBO GEAR SET ======================================================================================================
var TURBO_HELMET Gear = Gear{
	Name:     "Turbo Helmet",
	BodyPart: HEAD,
	StatsEffect: Stats{
		Speed:     2,
		Technique: -2,
	},
}

var TURBO_GLOVES Gear = Gear{
	Name:     "Turbo Gloves",
	BodyPart: ARMS,
	StatsEffect: Stats{
		Speed:    2,
		Strength: -2,
	},
}

var TURBO_PAD Gear = Gear{
	Name:     "Turbo Pad",
	BodyPart: BODY,
	StatsEffect: Stats{
		Speed:   2,
		Passing: -2,
	},
}

var TURBO_BOOTS Gear = Gear{
	Name:     "Turbo Boots",
	BodyPart: LEGS,
	StatsEffect: Stats{
		Speed:    2,
		Shooting: -2,
	},
}

// CANNON GEAR SET ======================================================================================================
var CANNON_VISOR Gear = Gear{
	Name:     "Cannon Visor",
	BodyPart: HEAD,
	StatsEffect: Stats{
		Shooting:  2,
		Technique: -2,
	},
}

var CANNON_GLOVES Gear = Gear{
	Name:     "Cannon Gloves",
	BodyPart: ARMS,
	StatsEffect: Stats{
		Shooting: 2,
		Speed:    -2,
	},
}

var CANNON_PLATE Gear = Gear{
	Name:     "Cannon Plate",
	BodyPart: BODY,
	StatsEffect: Stats{
		Shooting: 2,
		Strength: -2,
	},
}

var CANNON_BOOTS Gear = Gear{
	Name:     "Cannon Boots",
	BodyPart: LEGS,
	StatsEffect: Stats{
		Shooting: 2,
		Passing:  -2,
	},
}

// CHAIN GEAR SET ======================================================================================================
var CHAIN_HELMET Gear = Gear{
	Name:     "Chain Helmet",
	BodyPart: HEAD,
	StatsEffect: Stats{
		Passing: 2,
		Speed:   -2,
	},
}

var CHAIN_GAUNTLETS Gear = Gear{
	Name:     "Chain Gauntlets",
	BodyPart: ARMS,
	StatsEffect: Stats{
		Passing:  2,
		Shooting: -2,
	},
}

var CHAIN_PLATE Gear = Gear{
	Name:     "Chain Plate",
	BodyPart: BODY,
	StatsEffect: Stats{
		Passing:   2,
		Technique: -2,
	},
}

var CHAIN_BOOTS Gear = Gear{
	Name:     "Chain Boots",
	BodyPart: LEGS,
	StatsEffect: Stats{
		Passing:  2,
		Strength: -2,
	},
}

// TRICK GEAR SET ======================================================================================================
var TRICK_HELMET Gear = Gear{
	Name:     "Trick Helmet",
	BodyPart: HEAD,
	StatsEffect: Stats{
		Technique: 2,
		Passing:   -2,
	},
}

var TRICK_GLOVES Gear = Gear{
	Name:     "Trick Gloves",
	BodyPart: ARMS,
	StatsEffect: Stats{
		Technique: 2,
		Speed:     -2,
	},
}

var TRICK_PAD Gear = Gear{
	Name:     "Trick Pad",
	BodyPart: BODY,
	StatsEffect: Stats{
		Technique: 2,
		Strength:  -2,
	},
}

var TRICK_BOOTS Gear = Gear{
	Name:     "Trick Boots",
	BodyPart: LEGS,
	StatsEffect: Stats{
		Technique: 2,
		Shooting:  -2,
	},
}

// BUSHUDO GEAR SET ======================================================================================================
var BUSHIDO_HELMET Gear = Gear{
	Name:     "Bushido Helmet",
	BodyPart: HEAD,
	StatsEffect: Stats{
		Passing:   4,
		Strength:  -1,
		Speed:     -1,
		Shooting:  -1,
		Technique: -1,
	},
}

var BUSHIDO_BRACERS Gear = Gear{
	Name:     "Bushido Bracers",
	BodyPart: ARMS,
	StatsEffect: Stats{
		Passing:   -1,
		Strength:  -1,
		Speed:     -1,
		Shooting:  -1,
		Technique: 4,
	},
}

var BUSHIDO_ARMOR Gear = Gear{
	Name:     "Bushido ARMOR",
	BodyPart: BODY,
	StatsEffect: Stats{
		Passing:   -1,
		Strength:  4,
		Speed:     -1,
		Shooting:  -1,
		Technique: -1,
	},
}

var BUSHIDO_SANDALS Gear = Gear{
	Name:     "Bushido Sandals",
	BodyPart: LEGS,
	StatsEffect: Stats{
		Passing:   -1,
		Strength:  -1,
		Speed:     4,
		Shooting:  -1,
		Technique: -1,
	},
}

// GEAR SETS
var MUSCLE_GEAR_SET []Gear = []Gear{MUSCLE_HELMET, MUSCLE_GAUNTLETS, MUSCLE_CHEST, MUSCLE_BOOTS}
var TURBO_GEAR_SET []Gear = []Gear{TURBO_HELMET, TURBO_GLOVES, TURBO_PAD, TURBO_BOOTS}
var CANNON_GEAR_SET []Gear = []Gear{CANNON_VISOR, CANNON_GLOVES, CANNON_PLATE, CANNON_BOOTS}
var CHAIN_GEAR_SET []Gear = []Gear{CHAIN_HELMET, CHAIN_GAUNTLETS, CHAIN_PLATE, CHAIN_BOOTS}
var TRICK_GEAR_SET []Gear = []Gear{TRICK_HELMET, TRICK_GLOVES, TRICK_PAD, TRICK_BOOTS}
var BUSHIDO_GEAR_SET []Gear = []Gear{BUSHIDO_HELMET, BUSHIDO_BRACERS, BUSHIDO_ARMOR, BUSHIDO_SANDALS}
var HEAD_GEAR []Gear = []Gear{MUSCLE_HELMET, TURBO_HELMET, CANNON_VISOR, CHAIN_HELMET, TRICK_HELMET, BUSHIDO_HELMET}
var ARM_GEAR []Gear = []Gear{MUSCLE_GAUNTLETS, TURBO_GLOVES, CANNON_GLOVES, CHAIN_GAUNTLETS, TRICK_GLOVES, BUSHIDO_BRACERS}
var BODY_GEAR []Gear = []Gear{MUSCLE_CHEST, TURBO_PAD, CANNON_PLATE, CHAIN_PLATE, TRICK_PAD, BUSHIDO_ARMOR}
var LEG_GEAR []Gear = []Gear{MUSCLE_BOOTS, TURBO_BOOTS, CANNON_BOOTS, CHAIN_BOOTS, TRICK_BOOTS, BUSHIDO_SANDALS}

var GEAR_MAP map[string][]Gear = map[string][]Gear{
	"muscle":  MUSCLE_GEAR_SET,
	"turbo":   TURBO_GEAR_SET,
	"cannon":  CANNON_GEAR_SET,
	"chain":   CHAIN_GEAR_SET,
	"trick":   TRICK_GEAR_SET,
	"bushido": BUSHIDO_GEAR_SET,
	"head":    HEAD_GEAR,
	"arm":     ARM_GEAR,
	"body":    BODY_GEAR,
	"leg":     LEG_GEAR,
}

// Strikers ======================================================================================================
var MARIO Striker = Striker{
	Name: "Mario",
	BaseStats: Stats{
		Strength:  11,
		Speed:     12,
		Shooting:  14,
		Passing:   10,
		Technique: 16,
	},
}

var LUIGI Striker = Striker{
	Name: "Luigi",
	BaseStats: Stats{
		Strength:  11,
		Speed:     11,
		Shooting:  10,
		Passing:   14,
		Technique: 17,
	},
}

var BOWSER Striker = Striker{
	Name: "Bowser",
	BaseStats: Stats{
		Strength:  17,
		Speed:     9,
		Shooting:  17,
		Passing:   11,
		Technique: 9,
	},
}

var PEACH Striker = Striker{
	Name: "Peach",
	BaseStats: Stats{
		Strength:  9,
		Speed:     17,
		Shooting:  9,
		Passing:   13,
		Technique: 15,
	},
}

var ROSALINA Striker = Striker{
	Name: "Rosalina",
	BaseStats: Stats{
		Strength:  14,
		Speed:     9,
		Shooting:  17,
		Passing:   10,
		Technique: 13,
	},
}

var TOAD Striker = Striker{
	Name: "Toad",
	BaseStats: Stats{
		Strength:  9,
		Speed:     17,
		Shooting:  11,
		Passing:   15,
		Technique: 11,
	},
}

var YOSHI Striker = Striker{
	Name: "Yoshi",
	BaseStats: Stats{
		Strength:  10,
		Speed:     10,
		Shooting:  17,
		Passing:   17,
		Technique: 9,
	},
}

var DONKEY_KONG Striker = Striker{
	Name: "Donkey Kong",
	BaseStats: Stats{
		Strength:  16,
		Speed:     9,
		Shooting:  13,
		Passing:   16,
		Technique: 9,
	},
}

var WARIO Striker = Striker{
	Name: "Wario",
	BaseStats: Stats{
		Strength:  17,
		Speed:     9,
		Shooting:  15,
		Passing:   13,
		Technique: 9,
	},
}

var WALUIGI Striker = Striker{
	Name: "Waluigi",
	BaseStats: Stats{
		Strength:  15,
		Speed:     16,
		Shooting:  9,
		Passing:   9,
		Technique: 14,
	},
}

var ALL_STRIKERS map[string]Striker = map[string]Striker{
	MARIO.Name:       MARIO,
	LUIGI.Name:       LUIGI,
	BOWSER.Name:      BOWSER,
	PEACH.Name:       PEACH,
	ROSALINA.Name:    ROSALINA,
	TOAD.Name:        TOAD,
	YOSHI.Name:       YOSHI,
	DONKEY_KONG.Name: DONKEY_KONG,
	WARIO.Name:       WARIO,
	WALUIGI.Name:     WALUIGI,
}
