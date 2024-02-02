package restaurant

type ListRequest struct {
	City  string
	Genre string
}

type NotifyToLINERequest struct {
	Name        string
	Address     string
	StationName string
	GenreName   string
	URL         string
}
