package main

import (
	"fmt"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilog"
)

func start(_ *astilectron.Astilectron, iw *astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
	astilog.Info("TEXTILE STARTED")

	// check for an existing paired mobile id
	room, err := textile.GetRoomID()
	if err != nil {
		return err
	}
	if room != nil {
		// if we have one, start syncing
		astilog.Info("FOUND ROOM ID")

		// tell app what peer id we're gonna sync with
		sendData(iw, "sync.ready", map[string]interface{}{
			"pairedID": room.Pretty(),
			"html":     getPhotosHTML(),
		})

	} else {
		// otherwise, start onboaring
		astilog.Info("NO MOBILE PEER ID FOUND")
		astilog.Info("STARTING PAIRING")

		// sub to own peer id for pairing setup
		go func() {
			var idc = make(chan string)
			var errc = make(chan error)
			go func() {
				errc <- textile.StartPairing(idc)
			}()
			select {
			case id := <-idc:
				if id == "" {
					astilog.Errorf("failed to parse new paired id: %s", err)
					return
				}

				// let the app know we're done pairing
				sendMessage(iw, "onboard.complete")

				// and that we're ready to go
				sendData(iw, "sync.ready", map[string]interface{}{
					"pairedID": id,
					"html":     getPhotosHTML(),
				})
			case err := <-errc:
				astilog.Errorf("error while pairing: %s", err)
			}
		}()
		sendMessage(iw, "onboard.start")
	}

	return nil
}

func startSyncing(iw *astilectron.Window) error {
	astilog.Info("STARTING SYNC")

	// start subscription
	// TODO: expose cancel somehow
	cancel := make(chan struct{})
	datac, errc, err := textile.JoinRoom(cancel)
	if err != nil {
		return err
	}

	for {
		select {
		case hash := <-datac:
			sendData(iw, "sync.data", map[string]interface{}{
				"hash": hash,
			})
		case err := <-errc:
			astilog.Errorf("error while syncing: %s", err)
		}
	}
}

func getPhotosHTML() string {
	var html string
	for _, photo := range textile.Datastore.Photos().GetPhotos("", -1) {
		ph := fmt.Sprintf("http://localhost:9192/ipfs/%s/photo", photo.Cid)
		th := fmt.Sprintf("http://localhost:9192/ipfs/%s/thumb", photo.Cid)
		md := fmt.Sprintf("http://localhost:9192/ipfs/%s/meta", photo.Cid)
		img := fmt.Sprintf("<img src=\"%s\" />", th)
		html += fmt.Sprintf("<div class=\"grid-item\" data-url=\"%s\" data-meta=\"%s\">%s</div>", ph, md, img)
	}
	return html
}

func sendMessage(iw *astilectron.Window, name string) {
	iw.SendMessage(map[string]string{"name": name})
}

func sendData(iw *astilectron.Window, name string, data map[string]interface{}) {
	data["name"] = name
	iw.SendMessage(data)
}
