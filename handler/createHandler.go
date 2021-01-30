package handler

import (
	"github.com/desertbit/grumble"
	pkgerr "github.com/pkg/errors"
	"os/user"
	"strings"
	"sync"
	"tcssh/db"
	"tcssh/model"
	"tcssh/util/config"
	constant "tcssh/util/const"
	"tcssh/util/passwd"
	"time"
)

var (
	createOnce sync.Once
	createHan  Handler
)

type createHandler struct {
	location *config.CurrentLocation
}

func NewCreateHandler() Handler {
	createOnce.Do(func() {
		createHan = &createHandler{
			location: config.GlobalLocation,
		}
	})
	return createHan
}

func (c *createHandler) Handle(ctx *grumble.Context) (err error) {
	address := ctx.Args.String("address")
	userName, ip := splitAddress(address)
	port := ctx.Flags.Int("port")
	hostName := ctx.Flags.String("name")
	description := ctx.Flags.String("description")
	if hostName == "" {
		hostName = ip
	}
	pass, err := passwd.GetPasswd()
	if err != nil {
		return pkgerr.Wrap(err, "get passwd err")
	}
	tx := db.DB.Begin()
	node := &model.Node{
		DentryID:    c.location.GetLocation(),
		Name:        hostName,
		Ip:          ip,
		UserName:    userName,
		Password:    string(pass),
		Port:        port,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = model.CreateNode(tx, node)
	if err != nil {
		tx.Rollback()
		return pkgerr.Wrap(err, "create node err ")
	}
	dentry := &model.Dentry{
		ParentID:    c.location.GetLocation(),
		Name:        hostName,
		NodeId:      node.ID,
		Description: description,
		Type:        constant.DENTRY_TYPE_NODE,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = model.CreateDentry(tx, dentry)
	if err != nil {
		tx.Rollback()
		return pkgerr.Wrap(err, "create dentry err ")
	}
	err = tx.Commit().Error
	return
}

func splitAddress(add string) (user, ip string) {
	result := strings.Split(add, "@")
	if len(result) == 1 {
		// TODO if need to validate the ip?
		ip = result[0]
		user = getCurrentUser()
		return
	}
	user = result[0]
	ip = result[1]
	return
}

func getCurrentUser() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return u.Name
}