package model

//type Player struct {
//	UserId    uint   `gorm:"primaryKey"`
//	Mail      string `gorm:"index"`
//	CreatedAt time.Time
//	LoginAt   time.Time
//}

type Toponym struct {
	C1 uint
	C2 uint
	C3 uint `gorm:"primaryKey"`
	C4 string
	C5 string
	C6 string
	C7 string
	C8 string
	C9 string
}

func (Toponym) TableName() string {
	return "utf_ken_all"
}

type Unit struct {
	UnitId       string `gorm:"primaryKey"`
	PostCode     string
	UnitName     string
	Address      string
	Mail         string
	Tel          string
	CorporateNum string
	Type         string
}
