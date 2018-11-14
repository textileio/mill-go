package core

import (
	"crypto/rand"
	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	"gx/ipfs/QmdVrMn1LhB4ybb8hMVaMLXnA8XRSewMnK6YqXKXoTcRvN/go-libp2p-peer"
	libp2pc "gx/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5/go-libp2p-crypto"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"github.com/textileio/textile-go/repo"
)

func (a *api) addThreads(g *gin.Context) {
	args, err := a.readArgs(g)
	if err != nil {
		a.abort500(g, err)
		return
	}
	if len(args) == 0 {
		g.String(http.StatusBadRequest, "missing thread name")
		return
	}
	opts, err := a.readOpts(g)
	if err != nil {
		a.abort500(g, err)
		return
	}

	config := NewThreadConfig{
		Name: args[0],
		Join: true,
	}

	if opts["key"] != "" {
		config.Key = opts["key"]
	} else {
		config.Key = ksuid.New().String()
	}

	if opts["schema"] != "" {
		config.Schema, err = mh.FromB58String(opts["schema"])
		if err != nil {
			g.String(http.StatusBadRequest, err.Error())
		}
	}

	if opts["type"] != "" {
		var err error
		config.Type, err = repo.ThreadTypeFromString(opts["type"])
		if err != nil {
			g.String(http.StatusBadRequest, "invalid thread type")
			return
		}
	} else {
		config.Type = repo.OpenThread
	}

	// make a new secret
	sk, _, err := libp2pc.GenerateEd25519Key(rand.Reader)
	if err != nil {
		a.abort500(g, err)
		return
	}

	thrd, err := a.node.AddThread(sk, config)
	if err != nil {
		a.abort500(g, err)
		return
	}
	info, err := thrd.Info()
	if err != nil {
		a.abort500(g, err)
		return
	}
	g.JSON(http.StatusCreated, info)
}

func (a *api) lsThreads(g *gin.Context) {
	infos := make([]*ThreadInfo, 0)
	for _, thrd := range a.node.Threads() {
		info, err := thrd.Info()
		if err != nil {
			a.abort500(g, err)
			return
		}
		infos = append(infos, info)
	}
	g.JSON(http.StatusOK, infos)
}

func (a *api) getThreads(g *gin.Context) {
	id := g.Param("id")
	if id == "default" {
		id = a.node.config.Threads.Defaults.ID
	}
	thrd := a.node.Thread(id)
	if thrd == nil {
		g.String(http.StatusNotFound, ErrThreadNotFound.Error())
		return
	}
	info, err := thrd.Info()
	if err != nil {
		a.abort500(g, err)
		return
	}
	g.JSON(http.StatusOK, info)
}

func (a *api) rmThreads(g *gin.Context) {
	id := g.Param("id")
	thrd := a.node.Thread(id)
	if thrd == nil {
		g.String(http.StatusNotFound, ErrThreadNotFound.Error())
		return
	}
	if _, err := a.node.RemoveThread(id); err != nil {
		a.abort500(g, err)
		return
	}
	g.String(http.StatusOK, "ok")
}

func (a *api) createThreadInvites(g *gin.Context) {
	opts, err := a.readOpts(g)
	if err != nil {
		a.abort500(g, err)
		return
	}
	var pid peer.ID
	if opts["peer"] != "" {
		pid, err = peer.IDB58Decode(opts["peer"])
		if err != nil {
			g.String(http.StatusBadRequest, err.Error())
			return
		}
	}

	id := g.Param("id")
	if id == "default" {
		id = a.node.config.Threads.Defaults.ID
	}
	thrd := a.node.Thread(id)
	if thrd == nil {
		g.String(http.StatusNotFound, ErrThreadNotFound.Error())
		return
	}

	if pid != "" {
		if _, err := thrd.AddInvite(pid); err != nil {
			a.abort500(g, err)
			return
		}
	} else {
		if _, _, err := thrd.AddExternalInvite(); err != nil {
			a.abort500(g, err)
			return
		}
	}

	g.String(http.StatusOK, "ok")
}

func (a *api) addThreadFile(g *gin.Context) {
	opts, err := a.readOpts(g)
	if err != nil {
		a.abort500(g, err)
		return
	}

	var file repo.File
	if err := g.BindJSON(&file); err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	id := g.Param("id")
	if id == "default" {
		id = a.node.config.Threads.Defaults.ID
	}
	thrd := a.node.Thread(id)
	if thrd == nil {
		g.String(http.StatusNotFound, ErrThreadNotFound.Error())
		return
	}

	node, keys, err := a.node.AddNodeFromFiles([]repo.File{file})
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	hash, err := thrd.AddFiles(node, opts["caption"], keys)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	info, err := a.node.BlockInfo(hash.B58String())
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusCreated, info)
}

func (a *api) addThreadFiles(g *gin.Context) {
	opts, err := a.readOpts(g)
	if err != nil {
		a.abort500(g, err)
		return
	}

	var dir Directory
	if err := g.BindJSON(&dir); err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	id := g.Param("id")
	if id == "default" {
		id = a.node.config.Threads.Defaults.ID
	}
	thrd := a.node.Thread(id)
	if thrd == nil {
		g.String(http.StatusNotFound, ErrThreadNotFound.Error())
		return
	}

	node, keys, err := a.node.AddNodeFromDirs([]Directory{dir})
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	hash, err := thrd.AddFiles(node, opts["caption"], keys)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	info, err := a.node.BlockInfo(hash.B58String())
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusCreated, info)
}

func (a *api) streamThreads(g *gin.Context) {
	id := g.Param("id")
	thrd := a.node.Thread(id)
	if thrd == nil {
		g.String(http.StatusNotFound, "thread not found")
		return
	}
	opts, err := a.readOpts(g)
	if err != nil {
		a.abort500(g, err)
		return
	}

	listener := a.node.ThreadUpdateCh()
	g.Stream(func(w io.Writer) bool {
		select {
		case update, ok := <-listener.Ch:
			if !ok {
				return false
			}
			if opts["events"] == "true" {
				g.SSEvent("threadUpdate", update)
			} else {
				g.JSON(http.StatusOK, update)
			}
		default:
		}
		return true
	})

	listener.Close()
}
