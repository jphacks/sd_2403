package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
	"sushi/model"
	"sushi/utils/DB"
	"sushi/utils/config"
	"time"
)

type Service struct {
	db   *DB.DB
	log  *logrus.Logger
	conf *config.Config
	Ctx  *context.Context
}

func NewService(db *DB.DB, log *logrus.Logger, conf *config.Config, ctx context.Context) *Service {
	return &Service{
		db:   db,
		log:  log,
		conf: conf,
		Ctx:  &ctx,
	}
}

func (service *Service) GetTransportState(currentStatus string) string {
	switch currentStatus {
	case "0":
		return "生産"
	case "1":
		return "入庫"
	case "2":
		return "出庫"
	case "3":
		return "輸入"
	default:
		return "Unknown State"
	}
}
func (service *Service) GetCategory(currentStatus string) string {
	switch currentStatus {
	case "0":
		return "野菜"
	case "1":
		return "鮮魚"
	case "2":
		return "牛肉"
	default:
		return "Unknown State"
	}
}
func (service *Service) GetFreshState(currentStatus string) string {
	switch currentStatus {
	case "0":
		return "常温"
	case "1":
		return "冷蔵"
	case "2":
		return "冷凍"
	case "3":
		return "ドライアイス・氷"
	case "4":
		return "温度管理パッケージ"
	case "5":
		return "解凍"
	case "6":
		return "二酸化炭素による活魚輸送"
	default:
		return "Unknown State"
	}
}

func (service *Service) GetToponymByID(id string) (model.Toponym, error) {
	var toponym model.Toponym
	if err := service.db.DB.Model(&toponym).Where("c3 = ?", id).First(&toponym).Error; err != nil {
		return model.Toponym{}, err
	}
	return toponym, nil
}

func (service *Service) GetUnitByID(id string) (model.Unit, error) {
	var unit model.Unit
	if err := service.db.DB.Model(&unit).Where("unit_id = ?", id).First(&unit).Error; err != nil {
		return model.Unit{}, err
	}
	return unit, nil
}
func (service *Service) RegisterUnit(unitId string, postCode string, unitName string, address string, _type string, mail string, tel string, corporate_num string) error {
	unit := model.Unit{
		UnitId:       unitId,
		PostCode:     postCode,
		UnitName:     unitName,
		Address:      address,
		Mail:         mail,
		CorporateNum: corporate_num,
		Type:         _type,
		Tel:          tel,
	}
	err := service.createUnit(unit)
	if err != nil {
		return err
	}
	return nil
}
func (service *Service) createUnit(unit model.Unit) error {
	if err := service.db.DB.Create(&unit).Error; err != nil {
		return err
	}
	return nil
}
func (service *Service) TransTime(_time string) (string, error) {
	timeI, err := strconv.Atoi(_time)
	if err != nil {
		return "", err
	}
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(int64(timeI), 0).Format(timeLayout), nil
}
