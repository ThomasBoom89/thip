package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"thingiverse-zip-downloader-cli/thingiverse"
)

var modelId string

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&modelId, "modelId", "m", "", "thingiverse model id")
	rootCmd.MarkPersistentFlagRequired("modelId")
}

var rootCmd = &cobra.Command{
	Use:   "thip",
	Short: "thip is a CLI-Tool to download a zip archive of all files from a thingiverse model",
	Long: `thip is a CLI-Tool to download a zip archive of all files from a thingiverse model.
thip is built with love by spf13 and friends in Go.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		model, err := thingiverse.GetModel(modelId)
		if err != nil {
			fmt.Println("something went wrong: ", err.Error())
			os.Exit(1)
		}
		filename, err := thingiverse.ToZipArchive(model)
		if err != nil {
			fmt.Println("couldnt zip: ", err.Error())
			os.Exit(1)
		}
		fmt.Println(filename + " successfull created!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
