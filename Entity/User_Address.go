package Entity

type user_address struct {
	id_address string  `gorm:"primaryKey;type:varchar(36);"`
	Users      []users `gorm:"foreignKey:id_users"`
	street     string
	city       string
	state      string
	ZIP        int
}
