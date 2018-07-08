package service

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/segmentio/ksuid"
	"github.com/textileio/textile-go/crypto"
	"github.com/textileio/textile-go/pb"
	"github.com/textileio/textile-go/repo"
	"gx/ipfs/QmZoWKhxUmZ2seW4BzX6fJkNR8hh9PsGModr7q171yq2SS/go-libp2p-peer"
	libp2pc "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"
	"gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
	"gx/ipfs/Qmej7nf81hi2x2tvjRBF3mcp74sQyuDH4VMYDGd1YtXjb2/go-block-format"
)

func (s *TextileService) HandlerForMsgType(t pb.Message_MessageType) func(peer.ID, *pb.Message, interface{}) (*pb.Message, error) {
	switch t {
	case pb.Message_PING:
		return s.handlePing
	case pb.Message_THREAD_BLOCK:
		return s.handleThreadBlock
	case pb.Message_OFFLINE_ACK:
		return s.handleOfflineAck
	case pb.Message_OFFLINE_RELAY:
		return s.handleOfflineRelay
	case pb.Message_BLOCK:
		return s.handleBlock
	case pb.Message_STORE:
		return s.handleStore
	case pb.Message_ERROR:
		return s.handleError
	default:
		return nil
	}
}

func (s *TextileService) handlePing(pid peer.ID, pmes *pb.Message, options interface{}) (*pb.Message, error) {
	log.Debugf("received PING message from %s", pid.Pretty())
	return pmes, nil
}

func (s *TextileService) handleThreadBlock(pid peer.ID, pmes *pb.Message, options interface{}) (*pb.Message, error) {
	if pmes.Payload == nil {
		return nil, errors.New("payload is nil")
	}
	log.Debug("received THREAD_BLOCK message")

	// unpack it
	signed := new(pb.SignedThreadBlock)
	err := ptypes.UnmarshalAny(pmes.Payload, signed)
	if err != nil {
		return nil, err
	}
	block := new(pb.ThreadBlock)
	err = proto.Unmarshal(signed.Data, block)
	if err != nil {
		return nil, err
	}
	thrd := s.getThread(signed.ThreadId)

	switch block.Type {
	case pb.ThreadBlock_INVITE:
		log.Debug("handling ThreadBlock_INVITE...")
		if thrd != nil {
			return nil, errors.New("thread exists")
		}
		if block.Target != s.self.Pretty() {
			// TODO: should not be error right?
			return nil, errors.New("invalid invite target")
		}
		skb, err := crypto.Decrypt(s.node.PrivateKey, block.TargetKey)
		if err != nil {
			return nil, err
		}
		sk, err := libp2pc.UnmarshalPrivateKey(skb)
		if err != nil {
			return nil, err
		}
		good, err := sk.GetPublic().Verify(signed.Data, signed.Signature)
		if err != nil || !good {
			return nil, errors.New("bad signature")
		}
		// TODO: handle when name leads to conflict (add an int)
		thrd, err = s.addThread(signed.ThreadName, sk)
		if err != nil {
			return nil, err
		}

		// add inviter as local peer
		ppk, err := libp2pc.UnmarshalPublicKey(signed.IssuerPubKey)
		if err != nil {
			return nil, err
		}
		ppkb, err := ppk.Bytes()
		if err != nil {
			return nil, err
		}
		peerId, err := peer.IDFromPublicKey(ppk)
		if err != nil {
			return nil, err
		}
		newPeer := &repo.Peer{
			Row:      ksuid.New().String(),
			Id:       peerId.Pretty(),
			ThreadId: thrd.Id,
			PubKey:   ppkb,
		}
		if err := s.datastore.Peers().Add(newPeer); err != nil {
			return nil, err
		}
	case pb.ThreadBlock_PHOTO:
		log.Debug("handling ThreadBlock_PHOTO")
		if thrd == nil {
			return nil, errors.New("thread not found")
		}
		good, err := thrd.Verify(signed.Data, signed.Signature)
		if err != nil || !good {
			return nil, errors.New("bad signature")
		}

	case pb.ThreadBlock_COMMENT:
		return nil, errors.New("TODO")
	case pb.ThreadBlock_LIKE:
		return nil, errors.New("TODO")
	}

	// handle block
	return nil, thrd.HandleBlock(signed.Id)
}

func (s *TextileService) handleOfflineAck(pid peer.ID, pmes *pb.Message, options interface{}) (*pb.Message, error) {
	if pmes.Payload == nil {
		return nil, errors.New("payload is nil")
	}
	_, err := peer.IDB58Decode(string(pmes.Payload.Value))
	if err != nil {
		return nil, err
	}
	pointer, err := s.datastore.Pointers().Get(pid)
	if err != nil {
		return nil, err
	}
	if pointer.CancelId == nil || pointer.CancelId.Pretty() != pid.Pretty() {
		return nil, errors.New("peer is not authorized to delete pointer")
	}
	err = s.datastore.Pointers().Delete(pid)
	if err != nil {
		return nil, err
	}
	log.Debugf("received OFFLINE_ACK message from %s", pid.Pretty())
	return nil, nil
}

func (s *TextileService) handleOfflineRelay(pid peer.ID, pmes *pb.Message, options interface{}) (*pb.Message, error) {
	if pmes.Payload == nil {
		return nil, errors.New("payload is nil")
	}
	plaintext, err := crypto.Decrypt(s.node.PrivateKey, pmes.Payload.Value)
	if err != nil {
		return nil, err
	}

	// Unmarshal plaintext
	env := pb.Envelope{}
	err = proto.Unmarshal(plaintext, &env)
	if err != nil {
		return nil, err
	}

	// Validate the signature
	ser, err := proto.Marshal(env.Message)
	if err != nil {
		return nil, err
	}
	pubkey, err := libp2pc.UnmarshalPublicKey(env.Pubkey)
	if err != nil {
		return nil, err
	}
	valid, err := pubkey.Verify(ser, env.Signature)
	if err != nil || !valid {
		return nil, err
	}

	id, err := peer.IDFromPublicKey(pubkey)
	if err != nil {
		return nil, err
	}

	// Get handler for this message type
	handler := s.HandlerForMsgType(env.Message.MessageType)
	if handler == nil {
		log.Debug("got back nil handler from HandlerForMsgType")
		return nil, nil
	}

	// Dispatch handler
	_, err = handler(id, env.Message, true)
	if err != nil {
		log.Errorf("handle message error: %s", err)
		return nil, err
	}
	log.Debugf("received OFFLINE_RELAY message from %s", pid.Pretty())
	return nil, nil
}

func (s *TextileService) handleBlock(pid peer.ID, pmes *pb.Message, options interface{}) (*pb.Message, error) {
	if pmes.Payload == nil {
		return nil, errors.New("payload is nil")
	}
	pbblock := new(pb.Block)
	err := ptypes.UnmarshalAny(pmes.Payload, pbblock)
	if err != nil {
		return nil, err
	}
	id, err := cid.Decode(pbblock.Cid)
	if err != nil {
		return nil, err
	}
	block, err := blocks.NewBlockWithCid(pbblock.RawData, id)
	if err != nil {
		return nil, err
	}
	if err := s.node.Blocks.AddBlock(block); err != nil {
		return nil, err
	}
	log.Debugf("received IPFS_BLOCK message from %s", pid.Pretty())
	return nil, nil
}

func (s *TextileService) handleStore(pid peer.ID, pmes *pb.Message, options interface{}) (*pb.Message, error) {
	errorResponse := func(error string) *pb.Message {
		payload := &any.Any{Value: []byte(error)}
		message := &pb.Message{
			MessageType: pb.Message_ERROR,
			Payload:     payload,
		}
		return message
	}

	if pmes.Payload == nil {
		return nil, errors.New("payload is nil")
	}
	cList := new(pb.CidList)
	err := ptypes.UnmarshalAny(pmes.Payload, cList)
	if err != nil {
		return errorResponse("could not unmarshall message"), err
	}
	var need []string
	for _, id := range cList.Cids {
		decoded, err := cid.Decode(id)
		if err != nil {
			continue
		}
		has, err := s.node.Blockstore.Has(decoded)
		if err != nil || !has {
			need = append(need, decoded.String())
		}
	}
	log.Debugf("received STORE message from %s", pid.Pretty())
	log.Debugf("requesting %d blocks from %s", len(need), pid.Pretty())

	resp := new(pb.CidList)
	resp.Cids = need
	payload, err := ptypes.MarshalAny(resp)
	if err != nil {
		return errorResponse("error marshalling response"), err
	}
	message := &pb.Message{
		MessageType: pb.Message_STORE,
		Payload:     payload,
	}
	return message, nil
}

func (s *TextileService) handleError(peer peer.ID, pmes *pb.Message, options interface{}) (*pb.Message, error) {
	if pmes.Payload == nil {
		return nil, errors.New("payload is nil")
	}
	errorMessage := new(pb.Error)
	err := ptypes.UnmarshalAny(pmes.Payload, errorMessage)
	if err != nil {
		return nil, err
	}

	// TODO

	return nil, nil
}
