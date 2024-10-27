package controllor

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sushi/service"
	"sushi/utils"
	"sushi/utils/config"
)

type Controller struct {
	service *service.Service
	log     *logrus.Logger
	conf    *config.Config
}

func NewControllor(service *service.Service, log *logrus.Logger, conf *config.Config) *Controller {
	return &Controller{service: service, log: log, conf: conf}
}

func (con *Controller) HandlePing(c *gin.Context) {
	con.log.Debug("handling ping...")
	utils.SuccessResponse(c, "pong", "")
}

type SimpleToponym struct {
	PostCode string `json:"post_code"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
}

func (con *Controller) HandleToponym(c *gin.Context) {
	id := c.Query("id")
	toponym, err := con.service.GetToponymByID(id)
	if err != nil {
		utils.ErrorResponse(c, 400, "toponym not found", err)
		return
	}
	simpleToponym := SimpleToponym{
		PostCode: id,
		Address1: toponym.C7,
		Address2: toponym.C8,
		Address3: toponym.C9,
	}
	utils.SuccessResponse(c, "success", simpleToponym)
}

type UnitJson struct {
	UnitId       string `json:"unit_id"`
	UnitName     string `json:"unit_name"`
	PostCode     string `json:"post_code"`
	Address      string `json:"unit_address"`
	Mail         string `json:"unit_mail"`
	Tel          string `json:"unit_tel"`
	CorporateNum string `json:"corporate_num"`
	Type         string `json:"unit_type"`
}

func (con *Controller) HandleRegisterUnit(c *gin.Context) {
	unit := &UnitJson{}
	err := c.BindJSON(unit)
	if err != nil {
		utils.ErrorResponse(c, 400, "invalid unit", err)
		return
	}

	err = con.service.RegisterUnit(unit.UnitId, unit.PostCode, unit.UnitName, unit.Address, unit.Type, unit.Mail, unit.Tel, unit.CorporateNum)
	if err != nil {
		utils.ErrorResponse(c, 400, "unit not registered", err)
		return
	}
	utils.SuccessResponse(c, "success", "")
}
func (con *Controller) HandleGetUnit(c *gin.Context) {
	unitId := c.Query("unit_id")
	unit, err := con.service.GetUnitByID(unitId)
	if err != nil {
		utils.ErrorResponse(c, 400, "unit not found", err)
		return
	}
	utils.SuccessResponse(c, "success", unit)
}

type RecordRequest struct {
	UnitId         string `json:"unit_id"`
	TransPortState string `json:"transport_state"`
	Category       string `json:"category"`
	TimeStamp      string `json:"time_stamp"`
	FreshState     string `json:"fresh_state"`
}
type RecordJson struct {
	Category       string `json:"category"`
	UnitAddress    string `json:"unit_address"`
	UnitName       string `json:"unit_name"`
	UnitToponym    string `json:"unit_toponym"`
	TransportState string `json:"transport_state"`
	FreshState     string `json:"fresh_state"`
	Time           string `json:"time"`
}

func (con *Controller) HandleTransRecord(c *gin.Context) {
	recordRequest := &RecordRequest{}
	err := c.BindJSON(recordRequest)
	if err != nil {
		utils.ErrorResponse(c, 400, "invalid record", err)
		return
	}

	unit, err := con.service.GetUnitByID(recordRequest.UnitId)
	if err != nil {
		utils.ErrorResponse(c, 400, "unit not found", err)
		return

	}
	toponym, err := con.service.GetToponymByID(unit.PostCode)
	if err != nil {
		utils.ErrorResponse(c, 400, "toponym not found", err)
		return
	}
	time, err := con.service.TransTime(recordRequest.TimeStamp)
	if err != nil {
		utils.ErrorResponse(c, 400, "invalid time", err)
		return

	}
	record := RecordJson{
		Category:       con.service.GetCategory(recordRequest.Category),
		UnitAddress:    unit.Address,
		UnitName:       unit.UnitName,
		UnitToponym:    toponym.C7 + " " + toponym.C8 + " " + toponym.C9,
		TransportState: con.service.GetTransportState(recordRequest.TransPortState),
		FreshState:     con.service.GetFreshState(recordRequest.FreshState),
		Time:           time,
	}
	utils.SuccessResponse(c, "success", record)
}
