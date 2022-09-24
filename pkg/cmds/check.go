/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmds

import (
	"fmt"
	"os"

	api "voyagermesh.dev/apimachinery/apis/voyager/v1"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

func NewCmdCheck() *cobra.Command {
	var (
		fromFile      string
		cloudProvider string
	)
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check Ingress",
		RunE: func(cmd *cobra.Command, args []string) error {
			ingBytes, err := os.ReadFile(fromFile)
			if err != nil {
				return err
			}

			var ing api.Ingress
			err = yaml.Unmarshal(ingBytes, &ing)
			if err != nil {
				return err
			}
			err = ing.IsValid(cloudProvider)
			if err != nil {
				return err
			}
			fmt.Println("No validation error was found.")

			bi, err := yaml.Marshal(ing)
			if err != nil {
				return err
			}
			fmt.Println(string(bi))
			return nil
		},
	}

	cmd.Flags().StringVar(&fromFile, "from-file", fromFile, "YAML formatted file containing ingress")
	cmd.Flags().StringVarP(&cloudProvider, "cloud-provider", "c", cloudProvider, "Name of cloud provider")
	return cmd
}
