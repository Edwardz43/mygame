package betRecord

// Repository ...
type Repository interface {
	//TODO
	CreateOne(gameID int8, run int64, inn int, memberID int, distinctID int, amount int) (int, error)
}
