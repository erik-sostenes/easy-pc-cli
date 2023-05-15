package query

// Query represents the intention to request any data from our system without altering the state of our universe
//
// Query is a DTO(Data Transfer Object) able of representing the request you want to query
type Query interface {
	// QueryId method that implements all queries(Data Transfer Object)
	// returns a string representing the type of query to be performed
	QueryId() string
}
