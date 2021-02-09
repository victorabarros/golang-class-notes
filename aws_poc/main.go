package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	KMS "github.com/aws/aws-sdk-go/service/kms"
)

var (
	mySession        = session.Must(session.NewSession())
	kms              = KMS.New(mySession, aws.NewConfig().WithRegion("us-west-1"))
	messageType      = "RAW"
	signingAlgorithm = "RSASSA_PKCS1_V1_5_SHA_256"
	header           = map[string]string{
		"alg": "RS256",
		"typ": "JWT",
	}
	payload = map[string]interface{}{
		"iss":                  "AlloyCard",
		"custom:principalType": "com.alloycard.core.entities.recipe.Recipe",
	}
)

// BuildAlloyJWT build token on KMS
func BuildAlloyJWT(recipeID string, keyID string) (*KMS.SignOutput, error) {
	header["kid"] = fmt.Sprintf(`AlloyPrincipal-%s`, recipeID)

	payload["exp"] = time.Now().UTC().Add(time.Second * 60).UnixNano()
	payload["iat"] = time.Now().UTC().UnixNano()
	payload["custom:principalId"] = recipeID

	headerMarshal, err := json.Marshal(header)
	if err != nil {
		// TODO log err
		return nil, err
	}

	payloadMarshal, err := json.Marshal(payload)
	if err != nil {
		// TODO log err
		return nil, err
	}

	message := fmt.Sprintf("%s.%s",
		base64.StdEncoding.EncodeToString(headerMarshal),
		base64.StdEncoding.EncodeToString(payloadMarshal))

	resp, err := kms.Sign(&KMS.SignInput{
		KeyId:            &keyID,
		Message:          []byte(message),
		MessageType:      &messageType,
		SigningAlgorithm: &signingAlgorithm,
	})
	if err != nil {
		// TODO log err
		return nil, err
	}

	return resp, nil
}

func main() {
	resp, err := BuildAlloyJWT("e39d****-****-****-****-****6162d33e", "2")

	fmt.Println(resp, err)
	fmt.Printf("%+2v\n", resp)
}
