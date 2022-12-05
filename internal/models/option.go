package models

type Option struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Variant    string     `json:"variant"`
	Name       string     `json:"name"`
	Price      float32    `json:"price"`
	Quantity   uint32     `json:"quantity"`
	Categories []Category `gorm:"many2many:category_option;" json:"categories"`
}

func initOption(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Option{}) {
		manager.GetConn().AutoMigrate(&Option{})
	}
}
