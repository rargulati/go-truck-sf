package facility

type Location struct {
	needs_recoding	bool		`json:"needs_recoding"`
	longitude				string	`json:"longitude"`
	latitude				string	`json:"latitude"`
}

type Facility struct {
	Location						*Location
	Status							string	`json:"status,omitempty"`
	Expirationdate			string	`json:"expirationdate,omitempty"`
	Permit							string	`json:"permit,omitempty"`
	Block								string	`json:"block,omitempty"`
	Received						string	`json:"received,omitempty"`
	Facilitytype				string	`json:"facilitytype,omitempty"`
	Blocklot						string	`json:"blocklot,omitempty"`
	Noisent							string	`json:"noisent,omitempty"`
	Locationdescription	string	`json:"locationdescription,omitempty"`
	Cnn									string	`json:"cnn,omitempty"`
	Priorpermit					string	`json:"priorpermit,omitempty"`
	Approved						string	`json:"approved,omitempty"`
	Schedule						string	`json:"schedule,omitempty"`
	Address							string	`json:"address,omitempty"`
	Applicant						string	`json:"applicant,omitempty"`
	Lot									string	`json:"lot,omitempty"`
	Fooditems						string	`json:"fooditems,omitempty"`
	Longitude						string	`json:"longitude,omitempty"`
	Latitude						string	`json:"latitude,omitempty"`
	Objectid						string	`json:"objectid,omitempty"`
	Y										string	`json:"y,omitempty"`
	X										string	`json:"x,omitempty"`
}
