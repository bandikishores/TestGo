package data

// ContextKey - Type to identify the objects saved inside Context
type ContextKey int

const (
	// Session - The Session Object saved inside Context
	Session ContextKey = iota
	// OrgID - Stores OrgID
	OrgID
)
