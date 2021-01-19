package table

type Relation struct {
	ID        int64    `json:"id" ,pg:"id"`
	UID1      int64    `json:"uid1" ,pg:"uid1"`
	UID2      int64    `json:"uid2" ,pg:"uid2"`
	State     int      `json:"state" ,pg:"state"`
	tableName struct{} `pg:"test.relation"`
}
