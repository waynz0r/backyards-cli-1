// Copyright © 2019 Banzai Cloud
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

package graphql

import (
	"context"

	"github.com/MakeNowJust/heredoc"
)

type DisablePolicyPeersInput struct {
	Selector *PolicySelectorInput `json:"selector"`
}

func (c *client) DisablePolicyPeers(input DisablePolicyPeersInput) (bool, error) {
	request := heredoc.Doc(`
	  mutation disablePolicyPeers(
        $input: DisablePolicyPeersInput!
      ) {
        disablePolicyPeers(
          input: $input
        )
      }
`)

	r := c.NewRequest(request)
	r.Var("input", input)

	// run it and capture the response
	var respData map[string]bool
	if err := c.client.Run(context.Background(), r, &respData); err != nil {
		return false, err
	}

	return respData["disablePolicyPeers"], nil
}
