/*
Copyright © 2022 jenish HERE jenishjain@rocketmail.com

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/jenish-jain/sidekick/internal/helpers"
	"github.com/jenish-jain/sidekick/internal/mongo"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
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
	color.HiGreen("🍃 initiating mongo sizer ! 🍃")

	//TODO take inputs via flags
	mongoUri := fmt.Sprintf("mongouri")
	mongoClient := mongo.NewMongoClient(mongoUri)
	databases, err := mongoClient.ListDatabases(context.Background(), bson.D{})
	if err != nil {
		fmt.Printf("error listing databases %+v", err)
	}
	fmt.Printf("databases %+v", databases)

	for _, database := range databases.Databases {
		color.Blue("analysing db %s", database.Name)
		isTestDb, _ := regexp.MatchString("test", database.Name)
		if isTestDb {
			dbClient := mongoClient.Database(database.Name)
			time := helpers.NewTime()

			mongoService := mongo.NewMongoService(dbClient, time)
			mongoService.SizeIt()
		}
	}

}
