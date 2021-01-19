package table

type Users struct {
	ID   int64  `json:"id" ,pg:"id"`
	Name string `json:"name" ,pg:"name"`

	tableName struct{} `pg:"test.tt_users,alias:tt_users"`
}
