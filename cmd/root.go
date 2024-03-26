package cmd

import (
	"fmt"
	"os"

	"github.com/danielmesquitta/code_to_txt/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "code_to_txt",
	Short: "Converts any codebase to a single text file.",
	Long:  `Transforms codebases into single text files, suitable for AI model training or direct integration with the ChatGPT API to adhere to codebase standards.`,

	Run: func(cmd *cobra.Command, args []string) {
		folderPath, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println(err)
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
		outputFilePath, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err)
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
		ignorePath, err := cmd.Flags().GetStringSlice("ignore")
		if err != nil {
			fmt.Println(err)
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
		separator, err := cmd.Flags().GetString("separator")
		if err != nil {
			fmt.Println(err)
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
		err = usecase.ToTtx(folderPath, outputFilePath, usecase.ToTxtConfig{IgnorePaths: ignorePath, Separator: separator})
		if err != nil {
			fmt.Println("Invalid command")
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("path", "p", "./", "Path to the folder with code files")
	rootCmd.Flags().StringP("output", "o", "output.txt", "Path to the output txt file")
	rootCmd.Flags().StringSliceP("ignore", "i", []string{".git", ".env"}, "Paths and files to ignore while converting to txt file")
	rootCmd.Flags().StringP("separator", "s", "=== %s ===\n", "Separator to use between files in the output txt file, the %s will be replaced by the relative path of the file in the folder")
}
