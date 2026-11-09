type Response[T any] struct {
	Data T `json:"data"`
}

type RegisterData struct {
	Token    string   `json:"token"`
	Agent    Agent    `json:"agent"`
	Faction  Faction  `json:"faction"`
	Contract Contract `json:"contract"`
	Ships    []Ship   `json:"ships"`
}

type Ship struct {
	Symbol       string       `json:"symbol"`
	Registration Registration `json:"registration"`
	Nav          ShipNav      `json:"nav"`
	Crew         ShipCrew     `json:"crew"`
	Frame        ShipFrame    `json:"frame"`
	Reactor      ShipReactor  `json:"reactor"`
	Engine       ShipEngine   `json:"engine"`
	Modules      []ShipModule `json:"modules"`
	Mounts       []ShipMount  `json:"mounts"`
	Cargo        ShipCargo    `json:"cargo"`
	Fuel         ShipFuel     `json:"fuel"`
	Cooldown     Cooldown     `json:"cooldown"`
}

type RouteWaypoint struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type Requirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
	Slots int `json:"slots"`
}
