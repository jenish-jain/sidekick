/*
Copyright © 2022 jenish HERE jenishjain@rocketmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jenish-jain/sidekick/internal/constants"
	"github.com/jenish-jain/sidekick/internal/helpers"
	"github.com/jenish-jain/sidekick/internal/mongo"
	"github.com/spf13/cobra"
	"regexp"
)

// sizeMyMongoCmd represents the sizeMyMongo command
var sizeMyMongoCmd = &cobra.Command{
	Use:   "sizeMyMongo",
	Short: "A tool to help you size your mongodb",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sizeMyMongo called")
	},
}

func init() {
	rootCmd.AddCommand(sizeMyMongoCmd)
	color.HiGreen("Initiating mongo sizer !")

	//TODO take inputs via flags
	mongoUri := fmt.Sprintf("mongodb://user:pass@localhost:27017/dbName?readPreference=secondaryPreferred&authSource=dbName")
	dbName := "dbName"
	//mongoUri := helpers.GetLineInput("Enter mongo uri")
	//dbName := helpers.GetLineInput("Enter database name")

	fmt.Printf("Mongo uri: %s\n", mongoUri)
	mongoClient := mongo.NewMongoClient(mongoUri)

	database := mongoClient.Database(dbName)
	time := helpers.NewTime()

	mongoRepository := mongo.InitMongoRepository(database, time)
	collectionList := mongoRepository.GetCollectionNames()

	dbStats := mongoRepository.GetDBStats()
	fmt.Printf("STATS + %+v\n", dbStats)

	for i := 0; i < int(dbStats.CollectionsCount); i++ {
		isNotSystemCollection, _ := regexp.MatchString("system", collectionList[i])
		if !isNotSystemCollection {
			collectionName := collectionList[i]
			collStats := mongoRepository.GetCollectionStats(collectionName)
			collStats.CollectionAgeInDays = mongoRepository.GetCollectionAgeInDays(collectionName)
			color.Blue("Stats for collection : %s \n", collectionName)
			color.HiMagenta("%+v \n", collStats)

			color.HiRed("SPACE : %+v \n\n", getRequiredCollectionSizePerMonth(collStats))
			color.Yellow("_____________________________________________________")
		}
	}

}

func getRequiredCollectionSizePerMonth(collectionStats mongo.CollectionStats) *dbSize {
	compressedAvgDocSize := int64(0)
	if collectionStats.DocumentCount != 0 {
		compressedAvgDocSize = collectionStats.CompressedStorageSizeInBytes / collectionStats.DocumentCount
	}

	avgDocCountPerMonth := int64(0)
	if collectionStats.CollectionAgeInDays != 0 {
		fmt.Printf("DocumentCount %+v \n", collectionStats.DocumentCount)
		fmt.Printf("CollectionAgeInDays %+v \n", collectionStats.CollectionAgeInDays)
		avgDocCountPerMonth = (collectionStats.DocumentCount / collectionStats.CollectionAgeInDays) * constants.DaysInMonth
		fmt.Printf("avgDocCountPerMonth %+v \n", avgDocCountPerMonth)
	}

	return &dbSize{
		UncompressedDiscSpaceInBytes: int64(collectionStats.AvgObjSizeInBytes) * avgDocCountPerMonth,
		CompressedDiscSpaceInBytes:   compressedAvgDocSize * avgDocCountPerMonth,
		RAMInBytes:                   int64(0),
	}
}

type dbSize struct {
	UncompressedDiscSpaceInBytes int64 `json:"uncompressedDiscSpaceInBytes"`
	CompressedDiscSpaceInBytes   int64 `json:"compressedDiscSpaceInMb"`
	RAMInBytes                   int64 `json:"ramInMb"`
}
