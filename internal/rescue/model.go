package rescue

type Article struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Section  string   `json:"section"`
	Summary  string   `json:"summary"`
	Keywords []string `json:"keywords"`
}

type EquipmentItem struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Purpose  string `json:"purpose"`
}

type EquipmentGroup struct {
	Name  string          `json:"name"`
	Items []EquipmentItem `json:"items"`
}

type HomePage struct {
	Title      string    `json:"title"`
	SafetyTips []string  `json:"safetyTips"`
	Featured   []Article `json:"featured"`
	AuthorTeam []string  `json:"authorTeam"`
	Sections   []string  `json:"sections"`
	BasicsPath string    `json:"basicsPath"`
}

type SearchResult struct {
	Query           string                     `json:"query"`
	Articles        []Article                  `json:"articles"`
	Recommended     map[string][]EquipmentItem `json:"recommended"`
	StarterListPath string                     `json:"starterListPath"`
}
