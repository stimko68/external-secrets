/*
Copyright © The ESO Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import esmeta "github.com/external-secrets/external-secrets/apis/meta/v1"

type KeycardProvider struct {
	// URL configures the Keycard API URL. Defaults to https://api.keycard.ai/.
	URL  string      `json:"url,omitempty"`
	Auth KeycardAuth `json:"auth"`
	// optional scoping fields TBD: org/zone IDs
}

type KeycardAuth struct {
	WorkloadIdentity *KeycardWorkloadIdentityAuth `json:"workloadIdentity,omitempty"`
	// future: SecretRef *KeycardTokenAuth for static-token fallback
}

type KeycardWorkloadIdentityAuth struct {
	ServiceAccountRef esmeta.ServiceAccountSelector `json:"serviceAccountRef"`
	// Audiences sent to TokenRequest; Keycard will validate `aud` claim
	Audiences []string `json:"audiences,omitempty"`
}
