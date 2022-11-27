package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type CollectionStats struct {
	DocumentCount                int64                  `bson:"size"`
	AvgObjSizeInBytes            float64                `bson:"avgObjSize"`
	CompressedStorageSizeInBytes int64                  `bson:"storageSize"`
	IndexCount                   int32                  `bson:"nIndexes"`
	TotalIndexSizeInBytes        int64                  `bson:"totalIndexSize"`
	IndexSizesInBytes            map[string]interface{} `bson:"indexSizes"`
	CollectionAgeInDays          int64
}

type DbStats struct {
	Db                string  `bson:"db"`
	CollectionsCount  int64   `bson:"collections"`
	Objects           int64   `bson:"objects"`
	AvgObjSizeInBytes float64 `bson:"avgObjSize"`
	DataSizeInBytes   int64   `bson:"dataSize"`
	IndexCount        int64   `bson:"indexes"`
	IndexSizeInBytes  int64   `bson:"indexSize"`
}

type mongoDocument struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}
