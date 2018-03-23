package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type config struct {
	Section1 section `toml:"section1"`
	Section2 section `toml:"section2"`
}

type section struct {
	Key string `toml:"key"`
}

var (
	c = config{
		Section1: section{
			Key: "main-1",
		},
		Section2: section{
			Key: "main-2",
		},
	}

	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("rootCmd.Run :: c.Section1.Key   = %s\t <-- Why is this still 'config-1' instead of 'LALALA'?\n", c.Section1.Key)
			fmt.Printf(`viper.GetString("section1.key") = %s%s`, viper.GetString("section1.key"), "\n")

			fmt.Println()

			fmt.Printf("rootCmd.Run :: c.Section2.Key   = %s\t <-- Why is this still 'main-DEFAULT' instead of 'LALALA'?\n", c.Section2.Key)
			fmt.Printf(`viper.GetString("section2.key") = %s%s`, viper.GetString("section2.key"), "\n")
		},
	}
)

func main() {
	/*
		// Uncommenting these lines here "fixes" the issue
		// However it also requires executing the root command
		// without previously loading the config from file..
		if err := rootCmd.Execute(); err != nil {
			panic(err)
		}
	*/

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	rootCmd.PersistentFlags().String("key", "main-DEFAULT", "Override 'Key' fields")
	viper.BindPFlag("section1.key", rootCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("section2.key", rootCmd.PersistentFlags().Lookup("key"))
}
