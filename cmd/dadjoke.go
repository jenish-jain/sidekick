/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// dadjokeCmd represents the dadjoke command
var dadjokeCmd = &cobra.Command{
	Use:   "dadjoke",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(dadjokeCmd)
}

type Joke struct {
	ID     string `json:id`
	Joke   string `json:joke`
	Status int    `json:status`
}

func getRandomJoke() {
	url := "http://icanhazdadjoke.com/"
	jokeBytes := getJokeData(url)
	joke := Joke{}
	err := json.Unmarshal(jokeBytes, &joke)
	if err != nil {
		log.Printf("error unmarshaling dad joke -%v\n", err)
	}

	fmt.Print(joke.Joke)
}

func getJokeData(baseApi string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseApi, nil)
	if err != nil {
		log.Printf("error fetching dad joke - %v\n", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "SideKick CLI (github.com/jenish-jain/sidekick)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("error making request to icanhazdadjoke api - %v\n", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error reading dad joke - %v\n", err)
	}

	return responseBytes
}
