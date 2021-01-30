package handler

import (
	"errors"
	"github.com/desertbit/grumble"
	pkgerr "github.com/pkg/errors"
	"sync"
	"tcssh/db"
	"tcssh/model"
	"tcssh/util/config"
)

var (
	rmOnce sync.Once
	rmHan    Handler
)

type rmHandler struct {
	location *config.CurrentLocation
}

func NewRmHandler() Handler {
	rmOnce.Do(func() {
		rmHan = &rmHandler{
			location: config.GlobalLocation,
		}
	})
	return rmHan
}

func (h *rmHandler) Handle(c *grumble.Context) (err error) {
	dentryName := c.Args.String("dentry")
	if dentryName == "" {
		return errors.New("dentry name could not be null")
	}
	// first get dentry id by name
	dentryId,err := model.GetDentryIdByNameAndParentID(db.DB,dentryName,h.location.GetLocation())
	if err != nil {
		return pkgerr.Wrap(err,"get id by dentry name err ")
	}
	childDentrys,err := model.GetAllChildDentryByParentId(db.DB,dentryId.ID)
	if err != nil {
		return pkgerr.Wrap(err,"get children dentry err ")
	}
	dentrysToDel,nodesToDel := getAllDentryAndNodesToDel(childDentrys)
	tx := db.DB.Begin()
	if len(nodesToDel) > 0 {
		err = model.DeleteNodesByIds(tx,nodesToDel)
		if err != nil {
			tx.Rollback()
			return pkgerr.Wrap(err,"delete nodes err ")
		}
	}
	if len(dentrysToDel) > 0 {
		err = model.DeleteDentryByIds(tx,dentrysToDel)
		if err != nil {
			tx.Rollback()
			return pkgerr.Wrap(err,"delete dentrys err ")
		}
	}
	err = tx.Commit().Error
	return
}

func getAllDentryAndNodesToDel(d []model.Dentry) (dentrys,nodes []int) {
	for _, dentry := range d {
		if dentry.ID != 0 {
			dentrys = append(dentrys,dentry.ID)
		}
		if dentry.NodeId != 0 {
			nodes = append(nodes,dentry.NodeId)
		}
	}
	return
}
