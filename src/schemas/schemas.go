package schemas

// represents a row in the database
type Entry struct {
	ID       int
	Service  string
	Username string
	Password string
}
