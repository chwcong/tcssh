package handler

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
	"sync"
	"tcssh/db"
	"tcssh/model"
	"tcssh/util/config"
	constant "tcssh/util/const"
)

var (
	lsOnce sync.Once
	lsHan    Handler
)

type lsHandler struct {
	location *config.CurrentLocation
	printer  *printer
}

func NewLsHandler() Handler {
	lsOnce.Do(func() {
		lsHan = &lsHandler{
			location: config.GlobalLocation,
			printer:  initPrinter(),
		}
	})
	return lsHan
}

func (h *lsHandler) Handle(c *grumble.Context) (err error) {
	groups := model.GetDentryByParentID(db.DB, h.location.GetLocation())
	h.printDentry(c, groups)
	return
}

func (h *lsHandler) printDentry(c *grumble.Context, dentrys []model.Dentry) {
	nameString := ""
	for i := 0; i < len(dentrys); i++ {
		nameString = nameString + h.printer.sPrintColor(dentrys[i].Type,dentrys[i].Name) + "  "
	}
	c.App.Println(nameString)
}

type printer struct {
	typeColor map[string]*color.Color
}


func initPrinter() *printer {
	p := &printer{
		typeColor: make(map[string]*color.Color),
	}
	p.typeColor[constant.DENTRY_TYPE_NODE] = color.New(color.FgWhite)
	p.typeColor[constant.DENTRY_TYPE_GROUP] = color.New(color.FgBlue)
	p.typeColor["default"] = color.New(color.FgWhite)
	return p
}

func (p printer)sPrintColor(dentryType,content string) string {
	var colorPrinter *color.Color
	colorPrinter,ok := p.typeColor[dentryType]
	if !ok {
		colorPrinter = p.typeColor["default"]
	}
	return colorPrinter.Sprint(content)
}