// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clients

import "time"

// Commit represents a Git commit.
type Commit struct {
	CommittedDate       time.Time
	Message             string
	SHA                 string
	Committer           User
	AuthoredByCommitter bool
}

// User represents a Git user.
type User struct {
	Login string
}
