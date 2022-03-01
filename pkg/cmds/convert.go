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
	"bytes"
	"fmt"

	api "voyagermesh.dev/apimachinery/apis/voyager/v1"
	apiv1beta1 "voyagermesh.dev/apimachinery/apis/voyager/v1beta1"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"kmodules.xyz/client-go/tools/parser"
	"sigs.k8s.io/yaml"
)

func NewCmdConvert() *cobra.Command {
	var fromFile string
	cmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert v1beta1 to v1 api version",
		RunE: func(cmd *cobra.Command, args []string) error {
			resources, err := parser.ListPathResources(fromFile)
			if err != nil {
				return err
			}

			var buf bytes.Buffer

			ingressGVK := apiv1beta1.SchemeGroupVersion.WithKind("Ingress")
			for _, r := range resources {
				gvk := r.Object.GetObjectKind().GroupVersionKind()
				if gvk == ingressGVK {
					var obj apiv1beta1.Ingress
					if err := runtime.DefaultUnstructuredConverter.FromUnstructured(r.Object.UnstructuredContent(), &obj); err != nil {
						return err
					}
					var result api.Ingress
					if err := obj.ConvertTo(&result); err != nil {
						return err
					}
					if data, err := yaml.Marshal(result); err != nil {
						return err
					} else {
						if buf.Len() > 0 {
							buf.WriteString("\n---\n")
						}
						buf.Write(data)
					}
				}
			}
			fmt.Println(buf.String())
			return nil
		},
	}

	cmd.Flags().StringVarP(&fromFile, "from-file", "f", fromFile, "YAML formatted file containing ingress")
	return cmd
}
