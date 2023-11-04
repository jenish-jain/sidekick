package mongo

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jenish-jain/sidekick/internal/constants"
	"github.com/jenish-jain/sidekick/internal/helpers"
	q "github.com/jenish-jain/sidekick/pkg/quantity"
	"github.com/rodaine/table"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
)

type Service interface {
	SizeIt()
}

type service struct {
	mongoRepository repository
}

func (s service) SizeIt() {
	collectionList := s.mongoRepository.GetCollectionNames()

	serverStatus := s.mongoRepository.GetServerStatus()
	color.Blue("server status \n %+v \n", serverStatus)

	dbStats := s.mongoRepository.GetDBStats()
	color.HiRed("DB STATUS \n %+v \n", dbStats)
	color.HiRed("_________________________________________________________")
	color.HiRed("| UNCOMPRESSED SIZE IN GB PRESENT RAW %+v GB |", float64(dbStats.DataSizeInBytes)/constants.BytesInOneGb)
	color.HiRed("_________________________________________________________")
	color.HiRed("| COMPRESSED SIZE IN GB PRESENT RAW %+v GB |", dbStats.CompressedStorageSizeInBytes/constants.BytesInOneGb)
	color.HiRed("_________________________________________________________")

	//dbSize := &DbSize{
	//	UncompressedDiscSpaceInBytes: 0,
	//	CompressedDiscSpaceInBytes:   0,
	//	RAMInBytes:                   0,
	//}

	color.Cyan("Analysing individual collections in database \n")
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Index",
		"Collection Name",
		"Document Count",
		"Avg Obj size (KB)",
		"Uncompressed Size (GB)",
		"Compressed Size (GB)",
		"Index count",
		"Total index size (MB)",
		"Records maintained for (Days)",
		"Last accessed date",
		"Required Disc space comp (GB)",
	)
	forecastTbl := table.New("Collection Name",
		"Uncompressed disc space (GB)",
		"CompressedDiscSpace(GB)",
		"RAM (GB)",
	)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	forecastTbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	collectionForecastMap := map[string]DbSize{}

	for index, collection := range collectionList {
		isNotSystemCollection, _ := regexp.MatchString("system", collection)
		if !isNotSystemCollection {
			collStats := s.mongoRepository.GetCollectionStats(collection)
			//requiredSpaceForCollection := getRequiredCollectionSizePerMonth(collStats)
			//dbSize.UncompressedDiscSpaceInBytes += requiredSpaceForCollection.UncompressedDiscSpaceInBytes
			//dbSize.CompressedDiscSpaceInBytes += requiredSpaceForCollection.CompressedDiscSpaceInBytes
			collectionForecast := estimateRequiredCollectionSize(collStats)
			collectionForecastMap[collection] = collectionForecast
			tbl.AddRow(index+1,
				collection,
				collStats.DocumentCount,
				collStats.AvgObjSize.ConvertTo(q.KiloBytes).Value(q.WithPrecision(2)),
				collStats.UncompressedStorageSize.ConvertTo(q.GigaBytes).Value(q.WithPrecision(2)),
				collStats.CompressedStorageSize.ConvertTo(q.GigaBytes).Value(q.WithPrecision(2)),
				collStats.IndexCount,
				collStats.TotalIndexSize.ConvertTo(q.MegaBytes).Value(q.WithPrecision(2)),
				collStats.CollectionAgeInDays,
				collStats.LastAccessedDate,
				//toFixed(float64(requiredSpaceForCollection.CompressedDiscSpaceInBytes)/float64(constants.BytesInOneGb), 2),
			)

			forecastTbl.AddRow(
				collection,
				collectionForecast.UncompressedDiscSpace.ConvertTo(q.GigaBytes).Value(q.WithPrecision(2)),
				collectionForecast.CompressedDiscSpace.ConvertTo(q.GigaBytes).Value(q.WithPrecision(2)),
				collectionForecast.RAM.ConvertTo(q.GigaBytes).Value(q.WithPrecision(2)),
			)
		}
	}

	tbl.Print()
	color.HiRed("FORECASTED SUGGESTIONS")
	forecastTbl.Print()

	//color.HiRed("_________________________________________________________")
	//color.HiRed("| UNCOMPRESSED SIZE IN GB REQUIRED RAW %f GB  |", dbSize.UncompressedDiscSpace.ConvertTo(q.GigaBytes).Value(q.WithPrecision(2)))
	//color.HiRed("_________________________________________________________")
	//color.HiRed("| COMPRESSED SIZE IN GB REQUIRED RAW %f GB |", dbSize.CompressedDiscSpace.ConvertTo(q.GigaBytes).Value(q.WithPrecision(2)))
	//color.HiRed("_________________________________________________________")
	//color.HiCyan("TOTAL DB SIZE %+v", dbSize)
	fmt.Printf("STATS + %+v\n", dbStats)

}

func estimateRequiredCollectionSize(collectionStats CollectionStats) DbSize {
	compressedAvgDocSize := float64(0)
	if collectionStats.DocumentCount != 0 {
		compressedAvgDocSize = collectionStats.CompressedStorageSize.Value() / float64(collectionStats.DocumentCount)
	}

	avgDocCountPerMonth := float64(0)
	if collectionStats.CollectionAgeInDays != 0 {
		avgDocCountPerMonth = float64((collectionStats.DocumentCount / collectionStats.CollectionAgeInDays) * constants.DaysInMonth)
	}

	//fmt.Printf("DocumentCount %+v \n", collectionStats.DocumentCount)
	//fmt.Printf("CollectionAgeInDays %+v \n", collectionStats.CollectionAgeInDays)
	//fmt.Printf("avgDocCountPerMonth %+v \n", avgDocCountPerMonth)
	//fmt.Printf("compressed average doc size %+v \n", compressedAvgDocSize)

	return DbSize{
		UncompressedDiscSpace: q.New(collectionStats.AvgObjSize.Value()*avgDocCountPerMonth, q.Bytes),
		CompressedDiscSpace:   q.New(compressedAvgDocSize*avgDocCountPerMonth, q.Bytes),
		RAM:                   q.New(0, q.Bytes),
	}
}

func NewMongoService(database *mongo.Database, timeUtil helpers.Time) Service {
	repo := InitMongoRepository(database, timeUtil)
	return &service{
		mongoRepository: repo,
	}
}
