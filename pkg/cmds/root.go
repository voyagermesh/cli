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
	"voyagermesh.dev/apimachinery/client/clientset/versioned/scheme"

	"github.com/spf13/cobra"
	v "gomodules.xyz/x/version"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "voyager [command]",
		Short:             `Voyager cli by AppsCode`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			utilruntime.Must(scheme.AddToScheme(clientsetscheme.Scheme))
		},
	}

	// kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	// kubeConfigFlags.AddFlags(flags)
	// matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	// matchVersionKubeConfigFlags.AddFlags(flags)

	// f := cmdutil.NewFactory(matchVersionKubeConfigFlags)

	rootCmd.AddCommand(v.NewCmdVersion())
	rootCmd.AddCommand(NewCmdCompletion())

	rootCmd.AddCommand(NewCmdCheck())
	rootCmd.AddCommand(NewCmdConvert())

	return rootCmd
}
