package Entity

type users struct {
	id_user   string `gorm:"primaryKey;type:varchar(36);"`
	name      string
	status    bool
	nomer_HP  string
	email     string
	user_Name string
	password  string
}
