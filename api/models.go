package api

import "time"

type Response[T any] struct {
	Data T `json:"data"`
}

type ListResponse[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}

type Agent struct {
	AccountID       string `json:"accountId"`
	Symbol          string `json:"symbol"`
	Headquarters    string `json:"headquarters"`
	Credits         int    `json:"credits"`
	StartingFaction string `json:"startingFaction"`
	ShipCount       int    `json:"shipCount"`
}

type Chart struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	SubmittedBy    string    `json:"submittedBy"`
	SubmittedOn    time.Time `json:"submittedOn"`
}

type ConnectedSystem struct {
	Symbol        string `json:"symbol"`
	SectorSymbol  string `json:"sectorSymbol"`
	Type          string `json:"type"`
	FactionSymbol string `json:"factionSymbol"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Distance      int    `json:"distance"`
}

type ConstructionMaterial struct {
	TradeSymbol string `json:"tradeSymbol"`
	Required    int    `json:"required"`
	Fulfilled   int    `json:"fulfilled"`
}

type Construction struct {
	Symbol     string                 `json:"symbol"`
	Materials  []ConstructionMaterial `json:"materials"`
	IsComplete bool                   `json:"isComplete"`
}

type ContractPayment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}

type ContractDeliverGood struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled"`
}

type ContractTerms struct {
	Deadline time.Time             `json:"deadline"`
	Payment  ContractPayment       `json:"payment"`
	Deliver  []ContractDeliverGood `json:"deliver"`
}

type Contract struct {
	ID               string        `json:"id"`
	FactionSymbol    string        `json:"factionSymbol"`
	Type             string        `json:"type"`
	Terms            ContractTerms `json:"terms"`
	Accepted         bool          `json:"accepted"`
	Fulfilled        bool          `json:"fulfilled"`
	Expiration       time.Time     `json:"expiration"`
	DeadlineToAccept time.Time     `json:"deadlineToAccept"`
}

type Cooldown struct {
	ShipSymbol       string    `json:"shipSymbol"`
	TotalSeconds     int       `json:"totalSeconds"`
	RemainingSeconds int       `json:"remainingSeconds"`
	Expiration       time.Time `json:"expiration"`
}

type ExtractionYield struct {
	Symbol string `json:"symbol"`
	Units  int    `json:"units"`
}

type Extraction struct {
	ShipSymbol string          `json:"shipSymbol"`
	Yield      ExtractionYield `json:"yield"`
}

type FactionTrait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Faction struct {
	Symbol       string         `json:"symbol"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Headquarters string         `json:"headquarters"`
	Traits       []FactionTrait `json:"traits"`
	IsRecruiting bool           `json:"isRecruiting"`
}

type JumpGate struct {
	Symbol      string   `json:"symbol"`
	Connections []string `json:"connections"`
}

type TradeGood struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MarketTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TradeSymbol    string    `json:"tradeSymbol"`
	Type           string    `json:"type"`
	Units          int       `json:"units"`
	PricePerUnit   int       `json:"pricePerUnit"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}

type MarketTradeGood struct {
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
	TradeVolume   int    `json:"tradeVolume"`
	Supply        string `json:"supply"`
	Activity      string `json:"activity"`
	PurchasePrice int    `json:"purchasePrice"`
	SellPrice     int    `json:"sellPrice"`
}

type Market struct {
	Symbol       string              `json:"symbol"`
	Exports      []TradeGood         `json:"exports"`
	Imports      []TradeGood         `json:"imports"`
	Exchange     []TradeGood         `json:"exchange"`
	Transactions []MarketTransaction `json:"transactions"`
	TradeGoods   []MarketTradeGood   `json:"tradeGoods"`
}

type Meta struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type RepairTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}

type ShipRegistration struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}

type ShipNavRouteWaypoint struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type ShipNavRoute struct {
	Destination   ShipNavRouteWaypoint `json:"destination"`
	Origin        ShipNavRouteWaypoint `json:"origin"`
	DepartureTime time.Time            `json:"departureTime"`
	Arrival       time.Time            `json:"arrival"`
}

type ShipNav struct {
	SystemSymbol   string       `json:"systemSymbol"`
	WaypointSymbol string       `json:"waypointSymbol"`
	Route          ShipNavRoute `json:"route"`
	Status         string       `json:"status"`
	FlightMode     string       `json:"flightMode"`
}

type ScannedShipFrame struct {
	Symbol string `json:"symbol"`
}

type ScannedShipReactor struct {
	Symbol string `json:"symbol"`
}

type ScannedShipEngine struct {
	Symbol string `json:"symbol"`
}

type ScannedShipMount struct {
	Symbol string `json:"symbol"`
}

type ScannedShip struct {
	Symbol       string             `json:"symbol"`
	Registration ShipRegistration   `json:"registration"`
	Nav          ShipNav            `json:"nav"`
	Frame        ScannedShipFrame   `json:"frame"`
	Reactor      ScannedShipReactor `json:"reactor"`
	Engine       ScannedShipEngine  `json:"engine"`
	Mounts       []ScannedShipMount `json:"mounts"`
}

type ScannedSystem struct {
	Symbol       string `json:"symbol"`
	SectorSymbol string `json:"sectorSymbol"`
	Type         string `json:"type"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Distance     int    `json:"distance"`
}

type WaypointOrbital struct {
	Symbol string `json:"symbol"`
}

type WaypointFaction struct {
	Symbol string `json:"symbol"`
}

type WaypointTrait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ScannedWaypoint struct {
	Symbol       string            `json:"symbol"`
	Type         string            `json:"type"`
	SystemSymbol string            `json:"systemSymbol"`
	X            int               `json:"x"`
	Y            int               `json:"y"`
	Orbitals     []WaypointOrbital `json:"orbitals"`
	Faction      WaypointFaction   `json:"faction"`
	Traits       []WaypointTrait   `json:"traits"`
	Chart        Chart             `json:"chart"`
}

type ScrapTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}

type ShipCrew struct {
	Current  int    `json:"current"`
	Required int    `json:"required"`
	Capacity int    `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int    `json:"morale"`
	Wages    int    `json:"wages"`
}

type ShipRequirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
	Slots int `json:"slots"`
}

type ShipFrame struct {
	Symbol         string           `json:"symbol"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Condition      float64          `json:"condition"`
	Integrity      float64          `json:"integrity"`
	ModuleSlots    int              `json:"moduleSlots"`
	MountingPoints int              `json:"mountingPoints"`
	FuelCapacity   int              `json:"fuelCapacity"`
	Requirements   ShipRequirements `json:"requirements"`
	Quality        float64          `json:"quality"`
}

type ShipReactor struct {
	Symbol       string           `json:"symbol"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Condition    float64          `json:"condition"`
	Integrity    float64          `json:"integrity"`
	PowerOutput  int              `json:"powerOutput"`
	Requirements ShipRequirements `json:"requirements"`
	Quality      float64          `json:"quality"`
}

type ShipEngine struct {
	Symbol       string           `json:"symbol"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Condition    float64          `json:"condition"`
	Integrity    float64          `json:"integrity"`
	Speed        int              `json:"speed"`
	Requirements ShipRequirements `json:"requirements"`
	Quality      float64          `json:"quality"`
}

type ShipModule struct {
	Symbol       string           `json:"symbol"`
	Capacity     int              `json:"capacity"`
	Range        int              `json:"range"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Requirements ShipRequirements `json:"requirements"`
}

type ShipMount struct {
	Symbol       string           `json:"symbol"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Strength     int              `json:"strength"`
	Deposits     []string         `json:"deposits"`
	Requirements ShipRequirements `json:"requirements"`
}

type ShipCargoItem struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       int    `json:"units"`
}

type ShipCargo struct {
	Capacity  int             `json:"capacity"`
	Units     int             `json:"units"`
	Inventory []ShipCargoItem `json:"inventory"`
}

type ShipFuelConsumed struct {
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type ShipFuel struct {
	Current  int              `json:"current"`
	Capacity int              `json:"capacity"`
	Consumed ShipFuelConsumed `json:"consumed"`
}

type Ship struct {
	Symbol       string           `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav          ShipNav          `json:"nav"`
	Crew         ShipCrew         `json:"crew"`
	Frame        ShipFrame        `json:"frame"`
	Reactor      ShipReactor      `json:"reactor"`
	Engine       ShipEngine       `json:"engine"`
	Cooldown     Cooldown         `json:"cooldown"`
	Modules      []ShipModule     `json:"modules"`
	Mounts       []ShipMount      `json:"mounts"`
	Cargo        ShipCargo        `json:"cargo"`
	Fuel         ShipFuel         `json:"fuel"`
}

type ShipConditionEvent struct {
	Symbol      string `json:"symbol"`
	Component   string `json:"component"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ShipModificationTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TradeSymbol    string    `json:"tradeSymbol"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}

type ShipyardShipType struct {
	Type string `json:"type"`
}

type ShipyardTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	ShipType       string    `json:"shipType"`
	Price          int       `json:"price"`
	AgentSymbol    string    `json:"agentSymbol"`
	Timestamp      time.Time `json:"timestamp"`
}

type ShipyardShipCrew struct {
	Required int `json:"required"`
	Capacity int `json:"capacity"`
}

type ShipyardShip struct {
	Type          string           `json:"type"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Supply        string           `json:"supply"`
	Activity      string           `json:"activity"`
	PurchasePrice int              `json:"purchasePrice"`
	Frame         ShipFrame        `json:"frame"`
	Reactor       ShipReactor      `json:"reactor"`
	Engine        ShipEngine       `json:"engine"`
	Modules       []ShipModule     `json:"modules"`
	Mounts        []ShipMount      `json:"mounts"`
	Crew          ShipyardShipCrew `json:"crew"`
}

type Shipyard struct {
	Symbol           string                `json:"symbol"`
	ShipTypes        []ShipyardShipType    `json:"shipTypes"`
	Transactions     []ShipyardTransaction `json:"transactions"`
	Ships            []ShipyardShip        `json:"ships"`
	ModificationsFee int                   `json:"modificationsFee"`
}

type SiphonYield struct {
	Symbol string `json:"symbol"`
	Units  int    `json:"units"`
}

type Siphon struct {
	ShipSymbol string      `json:"shipSymbol"`
	Yield      SiphonYield `json:"yield"`
}

type SurveyDeposit struct {
	Symbol string `json:"symbol"`
}

type Survey struct {
	Signature  string          `json:"signature"`
	Symbol     string          `json:"symbol"`
	Deposits   []SurveyDeposit `json:"deposits"`
	Expiration time.Time       `json:"expiration"`
	Size       string          `json:"size"`
}

type SystemWaypoint struct {
	Symbol   string            `json:"symbol"`
	Type     string            `json:"type"`
	X        int               `json:"x"`
	Y        int               `json:"y"`
	Orbitals []WaypointOrbital `json:"orbitals"`
	Orbits   string            `json:"orbits"`
}

type SystemFaction struct {
	Symbol string `json:"symbol"`
}

type System struct {
	Symbol        string           `json:"symbol"`
	SectorSymbol  string           `json:"sectorSymbol"`
	Constellation string           `json:"constellation"`
	Name          string           `json:"name"`
	Type          string           `json:"type"`
	X             int              `json:"x"`
	Y             int              `json:"y"`
	Waypoints     []SystemWaypoint `json:"waypoints"`
	Factions      []SystemFaction  `json:"factions"`
}

type WaypointModifier struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Waypoint struct {
	Symbol              string             `json:"symbol"`
	Type                string             `json:"type"`
	SystemSymbol        string             `json:"systemSymbol"`
	X                   int                `json:"x"`
	Y                   int                `json:"y"`
	Orbitals            []WaypointOrbital  `json:"orbitals"`
	Orbits              string             `json:"orbits"`
	Faction             WaypointFaction    `json:"faction"`
	Traits              []WaypointTrait    `json:"traits"`
	Modifiers           []WaypointModifier `json:"modifiers"`
	Chart               Chart              `json:"chart"`
	IsUnderConstruction bool               `json:"isUnderConstruction"`
}

type StatusResponseStats struct {
	Accounts  int `json:"accounts"`
	Agents    int `json:"agents"`
	Ships     int `json:"ships"`
	Systems   int `json:"systems"`
	Waypoints int `json:"waypoints"`
}

type StatusResponseLeaderboardsMostCredit struct {
	AgentSymbol string `json:"agentSymbol"`
	Credits     int    `json:"credits"`
}

type StatusResponseLeaderboardsMostSubmittedChart struct {
	AgentSymbol string `json:"agentSymbol"`
	ChartCount  int    `json:"chartCount"`
}

type StatusResponseLeaderboards struct {
	MostCredits         []StatusResponseLeaderboardsMostCredit         `json:"mostCredits"`
	MostSubmittedCharts []StatusResponseLeaderboardsMostSubmittedChart `json:"mostSubmittedCharts"`
}

type StatusResponseServerResets struct {
	Next      string `json:"next"`
	Frequency string `json:"frequency"`
}

type StatusResponseAnnouncement struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type StatusResponseLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type StatusResponse struct {
	Status        string                       `json:"status"`
	Version       string                       `json:"version"`
	ResetDate     string                       `json:"resetDate"`
	Description   string                       `json:"description"`
	Stats         StatusResponseStats          `json:"stats"`
	Leaderboards  StatusResponseLeaderboards   `json:"leaderboards"`
	ServerResets  StatusResponseServerResets   `json:"serverResets"`
	Announcements []StatusResponseAnnouncement `json:"announcements"`
	Links         []StatusResponseLink         `json:"links"`
}

type RegisterRequest struct {
	Faction string `json:"faction"`
	Symbol  string `json:"symbol"`
	Email   string `json:"email,omitempty"`
}

type RegisterData struct {
	Agent    Agent    `json:"agent"`
	Contract Contract `json:"contract"`
	Faction  Faction  `json:"faction"`
	Ships    []Ship   `json:"ships"`
	Token    string   `json:"token"`
}

type SupplyConstructionRequest struct {
	ShipSymbol  string `json:"shipSymbol"`
	TradeSymbol string `json:"tradeSymbol"`
	Units       int    `json:"units"`
}

type SupplyConstructionData struct {
	Construction Construction `json:"construction"`
	Cargo        ShipCargo    `json:"cargo"`
}

type AcceptContractData struct {
	Agent    Agent    `json:"agent"`
	Contract Contract `json:"contract"`
}

type DeliverContractRequest struct {
	ShipSymbol  string `json:"shipSymbol"`
	TradeSymbol string `json:"tradeSymbol"`
	Units       int    `json:"units"`
}

type DeliverContractData struct {
	Contract Contract  `json:"contract"`
	Cargo    ShipCargo `json:"cargo"`
}

type FulfillContractData struct {
	Agent    Agent    `json:"agent"`
	Contract Contract `json:"contract"`
}

type PurchaseShipRequest struct {
	ShipType       string `json:"shipType"`
	WaypointSymbol string `json:"waypointSymbol"`
}

type PurchaseShipData struct {
	Agent       Agent               `json:"agent"`
	Ship        Ship                `json:"ship"`
	Transaction ShipyardTransaction `json:"transaction"`
}

type OrbitShipData struct {
	Nav ShipNav `json:"nav"`
}

type ShipRefineRequest struct {
	Produce string `json:"produce"`
}

type ShipRefineDataProduced struct {
	TradeSymbol string `json:"tradeSymbol"`
	Units       int    `json:"units"`
}

type ShipRefineDataConsumed struct {
	TradeSymbol string `json:"tradeSymbol"`
	Units       int    `json:"units"`
}

type ShipRefineData struct {
	Cargo    ShipCargo                `json:"cargo"`
	Cooldown Cooldown                 `json:"cooldown"`
	Produced []ShipRefineDataProduced `json:"produced"`
	Consumed []ShipRefineDataConsumed `json:"consumed"`
}

type CreateChartData struct {
	Chart    Chart    `json:"chart"`
	Waypoint Waypoint `json:"waypoint"`
	Agent    Agent    `json:"agent"`
}

type DockShipData struct {
	Nav ShipNav `json:"nav"`
}

type CreateSurveyData struct {
	Cooldown Cooldown `json:"cooldown"`
	Surveys  []Survey `json:"surveys"`
}

type ExtractResourcesRequest struct {
	Survey Survey `json:"survey,omitempty"`
}

type ExtractResourcesData struct {
	Cooldown   Cooldown             `json:"cooldown"`
	Extraction Extraction           `json:"extraction"`
	Cargo      ShipCargo            `json:"cargo"`
	Modifiers  []WaypointModifier   `json:"modifiers"`
	Events     []ShipConditionEvent `json:"events"`
}

type SiphonResourcesData struct {
	Cooldown Cooldown             `json:"cooldown"`
	Siphon   Siphon               `json:"siphon"`
	Cargo    ShipCargo            `json:"cargo"`
	Events   []ShipConditionEvent `json:"events"`
}

type ExtractResourcesWithSurveyData struct {
	Cooldown   Cooldown             `json:"cooldown"`
	Extraction Extraction           `json:"extraction"`
	Cargo      ShipCargo            `json:"cargo"`
	Events     []ShipConditionEvent `json:"events"`
}

type JettisonRequest struct {
	Symbol string `json:"symbol"`
	Units  int    `json:"units"`
}

type JettisonData struct {
	Cargo ShipCargo `json:"cargo"`
}

type JumpShipRequest struct {
	WaypointSymbol string `json:"waypointSymbol"`
}

type JumpShipData struct {
	Nav         ShipNav           `json:"nav"`
	Cooldown    Cooldown          `json:"cooldown"`
	Transaction MarketTransaction `json:"transaction"`
	Agent       Agent             `json:"agent"`
}

type NavigateShipRequest struct {
	WaypointSymbol string `json:"waypointSymbol"`
}

type NavigateShipData struct {
	Fuel   ShipFuel             `json:"fuel"`
	Nav    ShipNav              `json:"nav"`
	Events []ShipConditionEvent `json:"events"`
}

type PatchShipNavRequest struct {
	FlightMode string `json:"flightMode,omitempty"`
}

type PatchShipNavData struct {
	Nav    ShipNav              `json:"nav"`
	Fuel   ShipFuel             `json:"fuel"`
	Events []ShipConditionEvent `json:"events"`
}

type WarpShipRequest struct {
	WaypointSymbol string `json:"waypointSymbol"`
}

type WarpShipData struct {
	Fuel ShipFuel `json:"fuel"`
	Nav  ShipNav  `json:"nav"`
}

type SellCargoRequest struct {
	Symbol string `json:"symbol"`
	Units  int    `json:"units"`
}

type SellCargoData struct {
	Agent       Agent             `json:"agent"`
	Cargo       ShipCargo         `json:"cargo"`
	Transaction MarketTransaction `json:"transaction"`
}

type CreateShipSystemScanData struct {
	Cooldown Cooldown        `json:"cooldown"`
	Systems  []ScannedSystem `json:"systems"`
}

type CreateShipWaypointScanData struct {
	Cooldown  Cooldown          `json:"cooldown"`
	Waypoints []ScannedWaypoint `json:"waypoints"`
}

type CreateShipShipScanData struct {
	Cooldown Cooldown      `json:"cooldown"`
	Ships    []ScannedShip `json:"ships"`
}

type RefuelShipRequest struct {
	Units     int  `json:"units,omitempty"`
	FromCargo bool `json:"fromCargo,omitempty"`
}

type RefuelShipData struct {
	Agent       Agent             `json:"agent"`
	Fuel        ShipFuel          `json:"fuel"`
	Transaction MarketTransaction `json:"transaction"`
	Cargo       ShipCargo         `json:"cargo"`
}

type PurchaseCargoRequest struct {
	Symbol string `json:"symbol"`
	Units  int    `json:"units"`
}

type PurchaseCargoData struct {
	Agent       Agent             `json:"agent"`
	Cargo       ShipCargo         `json:"cargo"`
	Transaction MarketTransaction `json:"transaction"`
}

type TransferCargoRequest struct {
	TradeSymbol string `json:"tradeSymbol"`
	Units       int    `json:"units"`
	ShipSymbol  string `json:"shipSymbol"`
}

type TransferCargoData struct {
	Cargo ShipCargo `json:"cargo"`
}

type NegotiateContractData struct {
	Contract Contract `json:"contract"`
}

type InstallMountRequest struct {
	Symbol string `json:"symbol"`
}

type InstallMountData struct {
	Agent       Agent                       `json:"agent"`
	Mounts      []ShipMount                 `json:"mounts"`
	Cargo       ShipCargo                   `json:"cargo"`
	Transaction ShipModificationTransaction `json:"transaction"`
}

type RemoveMountRequest struct {
	Symbol string `json:"symbol"`
}

type RemoveMountData struct {
	Agent       Agent                       `json:"agent"`
	Mounts      []ShipMount                 `json:"mounts"`
	Cargo       ShipCargo                   `json:"cargo"`
	Transaction ShipModificationTransaction `json:"transaction"`
}

type GetScrapShipData struct {
	Transaction ScrapTransaction `json:"transaction"`
}

type ScrapShipData struct {
	Agent       Agent            `json:"agent"`
	Transaction ScrapTransaction `json:"transaction"`
}

type GetRepairShipData struct {
	Transaction RepairTransaction `json:"transaction"`
}

type RepairShipData struct {
	Agent       Agent             `json:"agent"`
	Ship        Ship              `json:"ship"`
	Transaction RepairTransaction `json:"transaction"`
}

type GetSupplyChainData struct {
	ExportToImportMap map[string][]string `json:"exportToImportMap"`
}

type InstallShipModuleRequest struct {
	Symbol string `json:"symbol"`
}

type InstallShipModuleData struct {
	Agent       Agent                       `json:"agent"`
	Modules     []ShipModule                `json:"modules"`
	Cargo       ShipCargo                   `json:"cargo"`
	Transaction ShipModificationTransaction `json:"transaction"`
}

type RemoveShipModuleRequest struct {
	Symbol string `json:"symbol"`
}

type RemoveShipModuleData struct {
	Agent       Agent                       `json:"agent"`
	Modules     []ShipModule                `json:"modules"`
	Cargo       ShipCargo                   `json:"cargo"`
	Transaction ShipModificationTransaction `json:"transaction"`
}
