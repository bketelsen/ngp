// Copyright © 2017 Brian Ketelsen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "newgo",
	Short: "A tool for initializing a new Go project",
	Long: `newgo is a tool for initializing a new Go project.  It creates a dockerfile and Makefile 
appropriate for your project.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newgo executed")
		fmt.Println(ProjectPath())
		data := make(map[string]string)
		data["ProjectPath"] = guessImportPath()
		data["DockerImage"] = DockerImage()
		data["GithubRepo"] = strings.Replace(guessImportPath(), "github.com/", "", -1)

		dockerfile := filepath.Join(getSrcPath(), "github.com", "bketelsen", "ngp", "templates", "Dockerfile.tmpl")
		rt, err := template.ParseFiles(dockerfile)
		if err != nil {
			fmt.Println(errors.Wrap(err, "reading dockerfile template"))
			return
		}
		rm, err := os.Create(filepath.Join(ProjectPath(), "Dockerfile"))
		if err != nil {
			fmt.Println("create dockerfile: ", err)
			return
		}
		defer rm.Close()
		err = rt.Execute(rm, data)
		if err != nil {
			fmt.Print("execute dockerfile template: ", err)
			return
		}

		makefile := filepath.Join(getSrcPath(), "github.com", "bketelsen", "ngp", "templates", "Makefile.tmpl")
		rt, err = template.ParseFiles(makefile)
		if err != nil {
			fmt.Println(errors.Wrap(err, "reading makefile template"))
			return
		}
		rm, err = os.Create(filepath.Join(ProjectPath(), "Makefile"))
		if err != nil {
			fmt.Println("create Makefile: ", err)
			return
		}
		defer rm.Close()
		err = rt.Execute(rm, data)
		if err != nil {
			fmt.Print("execute makefile template: ", err)
			return
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("package", "p", false, "this project is a package, not a command")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".newgo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".newgo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
