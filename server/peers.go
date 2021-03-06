package server

import (
	pb "github.com/PomeloCloud/BFTRaft4go/proto/server"
	"github.com/PomeloCloud/BFTRaft4go/utils"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	"log"
	"strconv"
)

func GetGroupPeersFromKV(txn *badger.Txn, group uint64) map[uint64]*pb.Peer {
	peers := map[uint64]*pb.Peer{}
	keyPrefix := ComposeKeyPrefix(group, GROUP_PEERS)
	iter := txn.NewIterator(badger.IteratorOptions{})
	iter.Seek(append(keyPrefix, utils.U64Bytes(0)...)) // seek the head
	for iter.ValidForPrefix(keyPrefix) {
		item := iter.Item()
		item_data := ItemValue(item)
		peer := pb.Peer{}
		proto.Unmarshal(*item_data, &peer)
		peers[peer.Id] = &peer
		iter.Next()
	}
	iter.Close()
	return peers
}

func (m *RTGroup) PeerUncommittedLogEntries(peer *pb.Peer) ([]*pb.LogEntry, *pb.LogEntry) {
	entries_ := []*pb.LogEntry{}
	prevEntry := &pb.LogEntry{
		Term:  0,
		Index: 0,
	}
	m.Server.DB.View(func(txn *badger.Txn) error {
		entries := []*pb.LogEntry{}
		iter := m.ReversedLogIterator(txn)
		nextLogIdx := peer.NextIndex
		for true {
			entry := iter.Current()
			if entry == nil {
				break
			}
			prevEntry = entry
			if entry.Index < nextLogIdx {
				break
			}
			entries = append(entries, entry)
			iter.Next()
		}
		if peer.NextIndex == 0 && peer.MatchIndex == 0 {
			// new peer, should set prevEntry = 0
			prevEntry = &pb.LogEntry{
				Term:  0,
				Index: 0,
			}
		}
		// reverse so the first will be the one with least index
		if len(entries) > 1 {
			for i := 0; i < len(entries)/2; i++ {
				j := len(entries) - i - 1
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
		entries_ = entries
		return nil
	})
	log.Println("prev index for", peer.Id, "is", prevEntry.Index, peer.MatchIndex, "/", peer.NextIndex)
	return entries_, prevEntry
}

func (s *BFTRaftServer) ScanHostedGroups(serverId uint64) map[uint64]*RTGroup {
	scanKey := utils.U32Bytes(GROUP_PEERS)
	res := map[uint64]*RTGroup{}
	s.DB.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.IteratorOptions{})
		iter.Seek(scanKey)
		groups := map[uint64]*RTGroup{}
		for iter.ValidForPrefix(scanKey) {
			item := iter.Item()
			val := ItemValue(item)
			peer := &pb.Peer{}
			proto.Unmarshal(*val, peer)
			if peer.Id == serverId {
				group := GetGroupFromKV(txn, peer.Group)
				if group != nil {
					defaultLeader := uint64(0)
					if len(s.GetGroupHosts(txn, group.Id)) < 2 {
						defaultLeader = s.Id
					}
					groups[peer.Group] = NewRTGroup(
						s, defaultLeader,
						GetGroupPeersFromKV(txn, peer.Group),
						group, FOLLOWER,
					)
				}
			}
			iter.Next()
		}
		iter.Close()
		res = groups
		return nil
	})
	log.Println("found", len(res), "hosted groups")
	for groupId, meta := range res {
		log.Println("scanned group:", meta.Group.Id)
		k := strconv.Itoa(int(groupId))
		s.GroupsOnboard.Set(k, meta)
	}
	return res
}

func (m *RTGroup) OnboardGroupPeersSlice() []*pb.Peer {
	peers := []*pb.Peer{}
	for _, peer := range m.GroupPeers {
		peers = append(peers, peer)
	}
	return peers
}

func (s *BFTRaftServer) SavePeer(txn *badger.Txn, peer *pb.Peer) error {
	if data, err := proto.Marshal(peer); err == nil {
		dbKey := append(ComposeKeyPrefix(peer.Group, GROUP_PEERS), utils.U64Bytes(peer.Id)...)
		return txn.Set(dbKey, data, 0x00)
	} else {
		return err
	}
}
