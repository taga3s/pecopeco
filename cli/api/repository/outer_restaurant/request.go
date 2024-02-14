package restaurant

type ListRequest struct {
	City  string
	Genre string
}

type NotifyToLINERequest struct {
	Name           string
	Address        string
	NearestStation string
	Genre          string
	URL            string
}
