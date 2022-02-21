package blockinfodatabase

import (
	"Chain/pkg/pro"
	"Chain/pkg/utils"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"
)

// BlockInfoDatabase is a wrapper for a levelDB
type BlockInfoDatabase struct {
	db *leveldb.DB
}

// New returns a BlockInfoDatabase given a Config
func New(config *Config) *BlockInfoDatabase {
	db, err := leveldb.OpenFile(config.DatabasePath, nil)
	if err != nil {
		utils.Debug.Printf("Unable to initialize BlockInfoDatabase with path {%v}", config.DatabasePath)
	}
	return &BlockInfoDatabase{db: db}
}

// StoreBlockRecord stores a BlockRecord in the BlockInfoDatabase.
func (blockInfoDB *BlockInfoDatabase) StoreBlockRecord(hash string, blockRecord *BlockRecord) {
	//TODO
	blockByte, _ := proto.Marshal(EncodeBlockRecord(blockRecord))
	if err := blockInfoDB.db.Put([]byte(hash), blockByte, nil); err != nil {
		utils.Debug.Printf("[StoreBlockRecord] Unable to store block record for block {%v}", hash)
	}
}

// GetBlockRecord returns a BlockRecord from the BlockInfoDatabase given
// the relevant block's hash.
func (blockInfoDB *BlockInfoDatabase) GetBlockRecord(hash string) *BlockRecord {
	//TODO
	if data, err := blockInfoDB.db.Get([]byte(hash), nil); err != nil {
		utils.Debug.Printf("[GetBlockRecord] block not in leveldb")
		return nil
	} else {
		pbr := &pro.BlockRecord{}
		if err := proto.Unmarshal(data, pbr); err != nil {
			utils.Debug.Printf("[GetBlockRecord] Failed to unmarshal record from Block {%v}:", hash, err)
		}
		br := DecodeBlockRecord(pbr)
		return br
	}
}
