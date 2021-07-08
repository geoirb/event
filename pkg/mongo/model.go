package mongo

// Event mongo datasheme. 
type Event struct {
	Type       string `bson:"type"`
	State      int64  `bson:"state"`
	StartedAt  int64  `bson:"started_at"`
	FinishedAt int64  `bson:"finished_at"`
}
