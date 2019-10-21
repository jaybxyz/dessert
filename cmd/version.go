package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/sirkon/goproxy/gomod"
	"github.com/spf13/cobra"

	yaml "gopkg.in/yaml.v2"
)

const (
	flagFormat = "format"
)

// Version defines the application version (defined at compile time)
// GitCommit defines the application commit hash (defined at compile time)
var (
	Version       = ""
	Commit        = ""
	versionFormat string
)

type versionInfo struct {
	Version string `json:"version" yaml:"version"`
	Commit  string `json:"commit" yaml:"commit"`
	SDK     string `json:"sdk" yaml:"sdk"`
	Go      string `json:"go" yaml:"go"`
}

func getVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print the version of Dessert",
		RunE: func(cmd *cobra.Command, args []string) error {
			modBz, err := ioutil.ReadFile("go.mod")
			if err != nil {
				return err
			}

			mod, err := gomod.Parse("go.mod", modBz)
			if err != nil {
				return err
			}

			verInfo := versionInfo{
				Version: Version,
				Commit:  Commit,
				SDK:     mod.Require["github.com/cosmos/cosmos-sdk"],
				Go:      fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
			}

			var bz []byte

			switch versionFormat {
			case "json":
				bz, err = json.Marshal(verInfo)

			default:
				bz, err = yaml.Marshal(&verInfo)
			}

			if err != nil {
				return err
			}

			_, err = fmt.Println(string(bz))
			return err
		},
	}

	versionCmd.Flags().StringVar(&versionFormat, flagFormat, "text", "Print the version in the given format (text | json)")

	return versionCmd
}
