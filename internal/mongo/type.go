package mongo

import (
	"github.com/jenish-jain/sidekick/pkg/quantity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CollectionStats struct {
	Name                    string                       `bson:"ns"`
	UncompressedStorageSize quantity.Quantity            `bson:"size"`
	DocumentCount           int64                        `bson:"count"`
	AvgObjSize              quantity.Quantity            `bson:"avgObjSize"`
	CompressedStorageSize   quantity.Quantity            `bson:"storageSize"`
	IndexCount              int32                        `bson:"nindexes"`
	TotalIndexSize          quantity.Quantity            `bson:"totalIndexSize"`
	IndexSizes              map[string]quantity.Quantity `bson:"indexSizes"`
	CollectionAgeInDays     int64                        `bson:"collectionAgeInDays"`
	LastAccessedDate        string                       `bson:"lastAccessedDate"`
}

type DbStats struct {
	Db                           string  `bson:"db"`
	CollectionsCount             int64   `bson:"collections"`
	Objects                      int64   `bson:"objects"`
	AvgObjSizeInBytes            float64 `bson:"avgObjSize"`
	DataSizeInBytes              int64   `bson:"dataSize"`
	IndexCount                   int64   `bson:"indexes"`
	IndexSizeInBytes             int64   `bson:"indexSize"`
	CompressedStorageSizeInBytes float64 `bson:"storageSize"`
}

type DbSize struct {
	UncompressedDiscSpace quantity.Quantity `json:"uncompressedDiscSpace"`
	CompressedDiscSpace   quantity.Quantity `json:"compressedDiscSpace"`
	RAM                   quantity.Quantity `json:"ram"`
}

type mongoDocument struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}

type ServerStatus struct {
	Version       string        `bson:"version"`
	Mem           mem           `bson:"mem"`
	WiredTiger    wiredTiger    `bson:"wiredTiger"`
	StorageEngine storageEngine `bson:"storageEngine"`
}

type mem struct {
	//documentation : https://www.mongodb.com/docs/v4.2/reference/command/serverStatus/#mem
	Bits                int32 `bson:"bits"`
	ResidentMemoryInMib int32 `bson:"resident"`
	VirtualMemoryInMib  int32 `bson:"virtual"`
	Supported           bool  `bson:"supported"`
	Mapped              int32 `bson:"mapped"`
	MappedWithJournal   int32 `bson:"mappedWithJournal"`
}

type wiredTiger struct {
	Cache cache `bson:"cache"`
}

type cache struct {
	BytesCurrentlyInCache int32 `bson:"bytes currently in the cache"`
}

type storageEngine struct {
	Name string `bson:"name"`
}
