package dw

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

var searcher engine.Engine

func init() {
	searcher.Init(types.EngineInitOptions{SegmenterDictionaries: "./data/dictionary.txt",
		StopTokenFile:           "./data/stop_tokens.txt",
		UsePersistentStorage:    true,
		PersistentStorageFolder: "./data/index",
		PersistentStorageShards: 10,
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.LocationsIndex,
		}})
}

func Add(dockId uint64, s string) {
	searcher.IndexDocument(dockId, types.DocumentIndexData{
		Content: s,
	})
	searcher.FlushIndex()
}

func Query(key string) types.SearchResponse {
	return searcher.Search(types.SearchRequest{
		Text: key,
	})
}
