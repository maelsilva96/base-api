package utils

import (
	"context"
	"encoding/base64"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/secrets"
)

func GetConfigProvider(privateKeyPassword string) *common.ConfigurationProvider {
	ociFileConfig := "/.oci/config"
	if IsDevelopment() {
		ociFileConfig = "~/.oci/config"
	}
	ociConfigProvider, err := common.ConfigurationProviderFromFile(ociFileConfig, privateKeyPassword)
	ValidErrorPanic(err)
	return &ociConfigProvider
}

func GetSecreteContent(secretsClient secrets.SecretsClient, secretId string, currencyStage secrets.GetSecretBundleStageEnum) (string, error) {
	request := secrets.GetSecretBundleRequest{
		Stage:    currencyStage,
		SecretId: &secretId,
	}
	response, err := secretsClient.GetSecretBundle(context.Background(), request)
	if err != nil {
		return "", err
	}
	conn := response.SecretBundleContent.(secrets.Base64SecretBundleContentDetails).Content
	rawDecodedText, err := base64.StdEncoding.DecodeString(*conn)
	if err != nil {
		return "", err
	}
	return string(rawDecodedText), nil
}
