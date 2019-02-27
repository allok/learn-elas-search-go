package main

import (
	"github.com/spf13/cobra"
)

/*

 $ hanap --help

 $ hanap server start
 $ hanap server stop
 $ hanap server --help
 $
 $ hanap client search index phrase -i golang -p '8080'
 $ hanap client destroy index -i golang
 $ hanap client reindex file -f ./index_file_go.csv -i golang -s .go
 $ hanap client reindex file -f ./index_file_solidity.csv -i solidity -s .sol
 $ hanap client reindex file -f ./index_file_rust.csv -i rust -s .rs
 $
*/

// Cobra is both a library for creating powerful modern CLI
// applications as well as a program to generate applications
// and command files

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {

	rootCmd = &cobra.Command{
		Use:   "hanap",
		Short: "Search application for finding application source code files",
	}
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)

	// viper.SetDefault("location", os.Getenv("HOME"))
	// viper.SetConfigName("config") // refers to ./config.yaml file
	// viper.AddConfigPath(".")      // refers to ./config.yaml file
	// viper.ReadInConfig()
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("No configuration file found")
	// }
	// viper.SetDefault("location", os.Getenv("HOME")) // see pork.yaml file for 'location' setup. If none, this set to user's home directory
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
