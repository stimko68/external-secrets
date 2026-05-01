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

// Package keycard implements a Keycard provider for External Secrets.
package keycard

import (
	"context"
	"errors"

	keycard "github.com/keycardai/keycard-go"
	corev1 "k8s.io/api/core/v1"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	esv1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1"
)

// Provider satisfies the provider interface.
type Provider struct{}

// keycardBase satisfies the provider.SecretsClient interface.
type keycardBase struct {
	kube      kclient.Client
	store     *esv1.KeycardProvider
	storeKind string
	namespace string
	client    *keycard.Client
}

// Capabilities returns the provider supported capabilities (ReadOnly, WriteOnly, ReadWrite).
func (k *Provider) Capabilities() esv1.SecretStoreCapabilities { return esv1.SecretStoreReadOnly }

// NewClient creates a new Keycard client.
func (k *Provider) NewClient(ctx context.Context, store esv1.GenericStore, kube kclient.Client, namespace string) (esv1.SecretsClient, error) {
	storeSpec := store.GetSpec()
	if storeSpec == nil || storeSpec.Provider == nil || storeSpec.Provider.Keycard == nil {
		return nil, errors.New("no store type or wrong store type")
	}
	storeSpecKeycard := storeSpec.Provider.Keycard

	kl := &keycardBase{
		kube:      kube,
		store:     storeSpecKeycard,
		namespace: namespace,
		storeKind: store.GetObjectKind().GroupVersionKind().Kind,
	}

	client, err := kl.getClient(ctx, storeSpecKeycard)
	if err != nil {
		return nil, err
	}
	kl.client = client

	return kl, nil
}

func (k *keycardBase) getClient(ctx context.Context, provider *esv1.KeycardProvider) (*keycard.Client, error) {
	return new(keycard.NewClient()), nil
}

func (k *Provider) ValidateStore(store esv1.GenericStore) (admission.Warnings, error) {
	// storeSpec := store.GetSpec()
	// keycardSpec := storeSpec.Provider.Keycard

	return nil, nil
}

// NewProvider creates a new Provider instance.
func NewProvider() esv1.Provider {
	return &Provider{}
}

// GetSecret returns a single secret from the provider.
func (k *keycardBase) GetSecret(_ context.Context, _ esv1.ExternalSecretDataRemoteRef) ([]byte, error) {
	return nil, errors.New("not implemented")
}

// PushSecret writes a single secret into the provider.
func (k *keycardBase) PushSecret(_ context.Context, _ *corev1.Secret, _ esv1.PushSecretData) error {
	return errors.New("not implemented")
}

// DeleteSecret deletes the secret from the provider.
func (k *keycardBase) DeleteSecret(_ context.Context, _ esv1.PushSecretRemoteRef) error {
	return errors.New("not implemented")
}

// SecretExists checks if a secret already exists in the provider.
func (k *keycardBase) SecretExists(_ context.Context, _ esv1.PushSecretRemoteRef) (bool, error) {
	return false, errors.New("not implemented")
}

// Validate checks if the client is configured correctly.
func (k *keycardBase) Validate() (esv1.ValidationResult, error) {
	return esv1.ValidationResultUnknown, nil
}

// GetSecretMap returns multiple k/v pairs from the provider.
func (k *keycardBase) GetSecretMap(_ context.Context, _ esv1.ExternalSecretDataRemoteRef) (map[string][]byte, error) {
	return nil, errors.New("not implemented")
}

// GetAllSecrets returns multiple k/v pairs from the provider.
func (k *keycardBase) GetAllSecrets(_ context.Context, _ esv1.ExternalSecretFind) (map[string][]byte, error) {
	return nil, errors.New("not implemented")
}

// Close closes the client connection.
func (k *keycardBase) Close(_ context.Context) error {
	return nil
}
