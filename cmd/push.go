/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright © 2017 Zain Malik
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.git-pif/utils"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Create fork and push current branch to the fork",
	Long: `The command checks if the forks exists. Creates the fork if it doesn't
exists. Push and set the remote for current branch the fork.

git pif push`,
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo, user, err := utils.CreateFork()
		if err != nil {
			fmt.Println(err)
		}
		err = utils.Push(owner, repo, user)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
