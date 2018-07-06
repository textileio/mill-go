package mobile

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/op/go-logging"
	"github.com/textileio/textile-go/central/models"
	tcore "github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/net"
	"github.com/textileio/textile-go/repo"
	"github.com/textileio/textile-go/wallet"
	"github.com/textileio/textile-go/wallet/thread"
	libp2pc "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"
	"time"
)

var log = logging.MustGetLogger("mobile")

// Message is a generic go -> bridge message structure
type Event struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

// newEvent transforms an event name and structured data in Event
func newEvent(name string, payload map[string]interface{}) *Event {
	event := &Event{Name: name}
	jsn, _ := toJSON(payload)
	event.Payload = jsn
	return event
}

// Messenger is used to inform the bridge layer of new data waiting to be queried
type Messenger interface {
	Notify(event *Event)
}

// NodeConfig is used to configure the mobile node
// NOTE: logLevel is one of: CRITICAL ERROR WARNING NOTICE INFO DEBUG
type NodeConfig struct {
	RepoPath      string
	CentralApiURL string
	LogLevel      string
	LogFiles      bool
}

// Mobile is the name of the framework (must match package name)
type Mobile struct {
	RepoPath  string
	Mnemonic  string
	Online    <-chan struct{} // not readable from bridges
	messenger Messenger
}

// ThreadItem is a simple meta data wrapper around a Thread
type ThreadItem struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Peers int    `json:"peers"`
}

// Threads is a wrapper around a list of Threads
type Threads struct {
	Items []ThreadItem `json:"items"`
}

// DeviceItem is a simple meta data wrapper around a Device
type DeviceItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Devices is a wrapper around a list of Devices
type Devices struct {
	Items []DeviceItem `json:"items"`
}

// BlockItem is a simple meta data wrapper around a Device
type BlockItem struct {
	Id      string         `json:"id"`
	Target  string         `json:"target"`
	Parents []string       `json:"parents"`
	Type    repo.BlockType `json:"type"`
	Date    time.Time      `json:"date"`
}

// Blocks is a wrapper around a list of Blocks
type Blocks struct {
	Items []BlockItem `json:"items"`
}

// PinRequests is a wrapper around multiple PinRequests
type PinRequests struct {
	Items []net.PinRequest `json:"items"`
}

// Create a gomobile compatible wrapper around TextileNode
func NewNode(config *NodeConfig, messenger Messenger) (*Mobile, error) {
	ll, err := logging.LogLevel(config.LogLevel)
	if err != nil {
		ll = logging.INFO
	}
	cconfig := tcore.NodeConfig{
		LogLevel: ll,
		LogFiles: config.LogFiles,
		WalletConfig: wallet.Config{
			RepoPath:   config.RepoPath,
			CentralAPI: config.CentralApiURL,
			IsMobile:   true,
		},
	}
	node, mnemonic, err := tcore.NewNode(cconfig)
	if err != nil {
		return nil, err
	}
	tcore.Node = node

	return &Mobile{RepoPath: config.RepoPath, Mnemonic: mnemonic, messenger: messenger}, nil
}

// Start the mobile node
func (m *Mobile) Start() error {
	online, err := tcore.Node.StartWallet()
	if err != nil {
		if err == wallet.ErrStarted {
			return nil
		}
		return err
	}
	m.Online = online

	go func() {
		<-online
		// subscribe to thread updates
		for _, thrd := range tcore.Node.Wallet.Threads() {
			go func(t *thread.Thread) {
				m.subscribe(t)
			}(thrd)
		}

		// notify UI we're ready
		m.messenger.Notify(newEvent("onOnline", map[string]interface{}{}))

		// publish
		tcore.Node.Wallet.PublishThreads()
	}()

	return nil
}

// Stop the mobile node
func (m *Mobile) Stop() error {
	if err := tcore.Node.StopWallet(); err != nil && err != wallet.ErrStopped {
		return err
	}
	return nil
}

// SignUpWithEmail creates an email based registration and calls core signup
func (m *Mobile) SignUpWithEmail(username string, password string, email string, referral string) error {
	// build registration
	reg := &models.Registration{
		Username: username,
		Password: password,
		Identity: &models.Identity{
			Type:  models.EmailAddress,
			Value: email,
		},
		Referral: referral,
	}
	return tcore.Node.Wallet.SignUp(reg)
}

// SignIn build credentials and calls core SignIn
func (m *Mobile) SignIn(username string, password string) error {
	// build creds
	creds := &models.Credentials{
		Username: username,
		Password: password,
	}
	return tcore.Node.Wallet.SignIn(creds)
}

// SignOut calls core SignOut
func (m *Mobile) SignOut() error {
	return tcore.Node.Wallet.SignOut()
}

// IsSignedIn calls core IsSignedIn
func (m *Mobile) IsSignedIn() bool {
	si, _ := tcore.Node.Wallet.IsSignedIn()
	return si
}

// GetId calls core GetId
func (m *Mobile) GetId() (string, error) {
	return tcore.Node.Wallet.GetId()
}

// GetUsername calls core GetUsername
func (m *Mobile) GetUsername() (string, error) {
	return tcore.Node.Wallet.GetUsername()
}

// GetAccessToken calls core GetAccessToken
func (m *Mobile) GetAccessToken() (string, error) {
	return tcore.Node.Wallet.GetAccessToken()
}

// Threads lists all threads
func (m *Mobile) Threads() (string, error) {
	var threads Threads
	for _, thrd := range tcore.Node.Wallet.Threads() {
		peers := thrd.Peers("", -1)
		item := ThreadItem{Id: thrd.Id, Name: thrd.Name, Peers: len(peers)}
		threads.Items = append(threads.Items, item)
	}

	// build json
	jsn, err := toJSON(threads)
	if err != nil {
		return "", err
	}
	return jsn, nil
}

// AddThread adds a new thread with the given name
func (m *Mobile) AddThread(name string, mnemonic string) (string, error) {
	var mnem *string
	if mnemonic != "" {
		mnem = &mnemonic
	}
	thrd, _, err := tcore.Node.Wallet.AddThreadWithMnemonic(name, mnem)
	if err != nil {
		return "", err
	}

	// subscribe to updates
	go m.subscribe(thrd)

	// build json
	peers := thrd.Peers("", -1)
	item := ThreadItem{
		Id:    thrd.Id,
		Name:  thrd.Name,
		Peers: len(peers),
	}
	jsn, err := toJSON(item)
	if err != nil {
		return "", err
	}
	return jsn, nil
}

// RemoveThread call core RemoveDevice
func (m *Mobile) RemoveThread(name string) error {
	return tcore.Node.Wallet.RemoveThread(name)
}

// Devices lists all devices
func (m *Mobile) Devices() (string, error) {
	var devices Devices
	for _, dev := range tcore.Node.Wallet.Devices() {
		item := DeviceItem{Id: dev.Id, Name: dev.Name}
		devices.Items = append(devices.Items, item)
	}

	// build json
	jsn, err := toJSON(devices)
	if err != nil {
		return "", err
	}
	return jsn, nil
}

// AddDevice calls core AddDevice
func (m *Mobile) AddDevice(name string, pubKey string) error {
	pkb, err := libp2pc.ConfigDecodeKey(pubKey)
	if err != nil {
		return err
	}
	pk, err := libp2pc.UnmarshalPublicKey(pkb)
	if err != nil {
		return err
	}
	return tcore.Node.Wallet.AddDevice(name, pk)
}

// RemoveDevice call core RemoveDevice
func (m *Mobile) RemoveDevice(name string) error {
	return tcore.Node.Wallet.RemoveDevice(name)
}

// AddPhoto adds a photo by path and shares it to a thread
func (m *Mobile) AddPhoto(path string, threadName string, caption string) (string, error) {
	_, thrd := tcore.Node.Wallet.GetThreadByName(threadName)
	if thrd == nil {
		return "", errors.New(fmt.Sprintf("could not find thread: %s", threadName))
	}
	added, err := tcore.Node.Wallet.AddPhoto(path)
	if err != nil {
		return "", err
	}
	shared, err := thrd.AddPhoto(added.Id, caption, added.Key)
	if err != nil {
		return "", err
	}

	// build json
	requests := PinRequests{}
	requests.Items = append(requests.Items, *added.RemoteRequest)
	requests.Items = append(requests.Items, *shared.RemoteRequest)
	jsn, err := toJSON(requests)
	if err != nil {
		return "", err
	}
	return jsn, nil
}

// SharePhoto adds an existing photo to a new thread
func (m *Mobile) SharePhoto(id string, threadName string, caption string) (string, error) {
	block, err := tcore.Node.Wallet.GetBlockByTarget(id)
	if err != nil {
		return "", err
	}
	fromThread := tcore.Node.Wallet.GetThread(block.ThreadPubKey)
	if fromThread == nil {
		return "", errors.New(fmt.Sprintf("could not find thread %s", block.ThreadPubKey))
	}
	_, toThread := tcore.Node.Wallet.GetThreadByName(threadName)
	if toThread == nil {
		return "", errors.New(fmt.Sprintf("could not find thread named %s", threadName))
	}
	key, err := fromThread.Decrypt(block.TargetKey)
	if err != nil {
		return "", err
	}

	// TODO: owner challenge
	shared, err := toThread.AddPhoto(id, caption, key)
	if err != nil {
		return "", err
	}

	// build json
	requests := PinRequests{}
	requests.Items = append(requests.Items, *shared.RemoteRequest)
	jsn, err := toJSON(requests)
	if err != nil {
		return "", err
	}
	return jsn, nil
}

// PhotoBlocks returns thread photo blocks with json encoding
func (m *Mobile) PhotoBlocks(offsetId string, limit int, threadName string) (string, error) {
	_, thrd := tcore.Node.Wallet.GetThreadByName(threadName)
	if thrd == nil {
		return "", errors.New(fmt.Sprintf("thread not found: %s", threadName))
	}

	// build json
	blocks := &Blocks{}
	for _, b := range thrd.Blocks(offsetId, limit, repo.PhotoBlock) {
		blocks.Items = append(blocks.Items, BlockItem{
			Id:      b.Id,
			Target:  b.Target,
			Parents: b.Parents,
			Type:    b.Type,
			Date:    b.Date,
		})
	}
	jsn, err := toJSON(blocks)
	if err != nil {
		return "", err
	}
	return jsn, nil
}

// GetBlockData calls GetBlockDataBase64 on a thread
func (m *Mobile) GetBlockData(id string, path string) (string, error) {
	block, err := tcore.Node.Wallet.GetBlock(id)
	if err != nil {
		log.Errorf("could not find block %s for path %s: %s", id, path, err)
		return "", err
	}
	thrd := tcore.Node.Wallet.GetThread(block.ThreadPubKey)
	if thrd == nil {
		err := errors.New(fmt.Sprintf("could not find thread: %s", block.ThreadPubKey))
		log.Error(err.Error())
		return "", err
	}

	return thrd.GetBlockDataBase64(fmt.Sprintf("%s/%s", id, path), block)
}

// GetFileData calls GetFileDataBase64 on a thread
func (m *Mobile) GetFileData(id string, path string) (string, error) {
	block, err := tcore.Node.Wallet.GetBlockByTarget(id)
	if err != nil {
		log.Errorf("could not find block for target %s: %s", id, err)
		return "", err
	}
	thrd := tcore.Node.Wallet.GetThread(block.ThreadPubKey)
	if thrd == nil {
		err := errors.New(fmt.Sprintf("could not find thread: %s", block.ThreadPubKey))
		log.Error(err.Error())
		return "", err
	}

	return thrd.GetFileDataBase64(fmt.Sprintf("%s/%s", id, path), block)
}

// subscribe to thread and pass updates to messenger
func (m *Mobile) subscribe(thrd *thread.Thread) {
	for {
		select {
		case update, ok := <-thrd.Updates():
			if !ok {
				return
			}
			m.messenger.Notify(newEvent("onThreadUpdate", map[string]interface{}{
				"id":          update.Id,
				"type":        int(update.Type),
				"target_id":   update.TargetId,
				"thread_id":   update.ThreadId,
				"thread_name": update.ThreadName,
			}))
		}
	}
}

// toJSON returns a json string and logs errors
func toJSON(any interface{}) (string, error) {
	jsonb, err := json.Marshal(any)
	if err != nil {
		log.Errorf("error marshaling json: %s", err)
		return "", err
	}
	return string(jsonb), nil
}
