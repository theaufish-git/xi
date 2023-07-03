/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"github.com/theaufish-git/xi-config/config"
	"github.com/theaufish-git/xi-config/internal/templates"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xi-config",
	Short: "Generate xi server configuration files from environment variables.",
	Long:  `Generate xi server configuration files from environment variables..`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: generate,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xi-config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringSliceP("file", "f", []string{"logging", "login", "main", "map", "network", "search"}, "Files to generate.")
	rootCmd.Flags().StringP("outdir", "o", "./", "Directory to write generated files.")
	rootCmd.Flags().StringP("prefix", "p", "xi", "Prefix of the envvars to parse.")
}

func generate(cmd *cobra.Command, args []string) {
	prefix, err := cmd.Flags().GetString("prefix")
	if err != nil {
		log.Fatalln(err)
	}

	outdir, err := cmd.Flags().GetString("outdir")
	if err != nil {
		log.Fatalln(err)
	}

	files, err := cmd.Flags().GetStringSlice("file")
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		filename := path.Join(outdir, fmt.Sprintf("%s.lua", f))

		var cfg interface{}
		var tmpl *template.Template
		switch f {
		case "logging":
			cfg = &config.Logging{}
			if err := envconfig.Process(prefix, cfg); err != nil {
				log.Fatalln(err)
			}
			tmpl = templates.LoggingConfig
		case "login":
			cfg = &config.Login{}
			if err := envconfig.Process(prefix, cfg); err != nil {
				log.Fatalln(err)
			}
			tmpl = templates.LoginConfig
		case "main":
			cfg = &config.Main{}
			if err := envconfig.Process(prefix, cfg); err != nil {
				log.Fatalln(err)
			}
			tmpl = templates.MainConfig
		case "map":
			cfg = &config.Map{}
			if err := envconfig.Process(prefix, cfg); err != nil {
				log.Fatalln(err)
			}
			tmpl = templates.MapConfig
		case "network":
			cfg = &config.Network{}
			if err := envconfig.Process(prefix, cfg); err != nil {
				log.Fatalln(err)
			}
			tmpl = templates.NetworkConfig
		case "search":
			cfg = &config.Search{}
			if err := envconfig.Process(prefix, cfg); err != nil {
				log.Fatalln(err)
			}
			tmpl = templates.SearchConfig
		default:
			log.Fatalf("%s is not a valid file.\n", f)
		}

		log.Println("generating:", filename)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		if err := tmpl.Execute(file, cfg); err != nil {
			log.Fatalln(err)
		}
	}
}
