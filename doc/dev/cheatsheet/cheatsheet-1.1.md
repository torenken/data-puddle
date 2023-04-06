Cheat sheet for coding session
---

### data-puddle-stack
src/data-puddle-stack.ts

```ts
export class DataPuddleStack extends Stack {
  constructor(scope: Construct, id: string, props: DataPuddleStackProps) {
    super(scope, id, props);
    ...
    secret.grantRead(provideTicketDataFunc);
  }
}
```

### Secret
src/data-puddle-secret.ts

```ts
export class DataPuddleSecret extends Secret {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      removalPolicy: RemovalPolicy.DESTROY,
      secretStringValue: SecretValue.unsafePlainText('Add a 32bit encryption key as base64 encoding later'),
    });
  }
}
```

### provide-ticket-data
foundation/encrypt/encrypt.go

```go
// cryptopasta - basic cryptography examples
//
// Written in 2015 by George Tankersley <george.tankersley@gmail.com> (https://github.com/gtank/cryptopasta)
//

// Package encrypt provides symmetric authenticated encryption using 256-bit AES-GCM with a random nonce.
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// Encrypt encrypts data using 256-bit AES-GCM.
func Encrypt(plaintext []byte, key *[32]byte) ([]byte, error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts data using 256-bit AES-GCM.
func Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

// NewEncryptionKey generates a random 256-bit key for Encrypt() and  Decrypt().
func NewEncryptionKey() (*[32]byte, error) {
	key := [32]byte{}
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}
	return &key, nil
}

```

### provide-ticket-data
app/services/provide-ticket-data/main.go

```go
package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/torenken/data-puddle/foundation/encrypt"
)

//current impl for coding session not for production (only prototyping - no logging/error handling)
//programming mode: make it work and then ...

var (
	s3s *s3.Client
	smc *secretsmanager.Client
)

func init() {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	s3s = s3.NewFromConfig(cfg)
	smc = secretsmanager.NewFromConfig(cfg)
}

func handle(ctx context.Context, s3e events.S3Event) error {
	fmt.Printf("starting service with %v s3event records.\n", len(s3e.Records))
	for _, s3Event := range s3e.Records {
		s3Entity := s3Event.S3
		fmt.Printf("executing s3event from %v bucket and %v object.\n", s3Entity.Bucket.Name, s3Entity.Object.Key)

		if err := executing(ctx, s3Entity); err != nil {
			return fmt.Errorf("executing s3 copy with encryption: %w", err)
		}
	}
	fmt.Println("successful processing of the service.")
	return nil
}

func executing(ctx context.Context, entity events.S3Entity) error {
	targetBucket := os.Getenv("TICKET_SYS_OUT_BUCKET")

	//fetching data from s3 bucket
	data, err := fetchData(ctx, entity)
	if err != nil {
		return err
	}

	key, err := getEncryptionKey(ctx)
	if err != nil {
		return err
	}

	ciphertext, err := encrypt.Encrypt(data, key)
	if err != nil {
		return err
	}

	/*	plaintext, err := encrypt.Decrypt(ciphertext, key)
		if err != nil {
			return err
		}*/

	//importing data to another s3 bucket
	if err := importData(ctx, entity, ciphertext); err != nil {
		return err
	}

	fmt.Printf("copying data from %v to %v successful.\n", entity.Bucket.Name, targetBucket)

	return nil
}

func fetchData(ctx context.Context, entity events.S3Entity) ([]byte, error) {
	object, err := s3s.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(entity.Bucket.Name),
		Key:    aws.String(entity.Object.Key),
	})
	if err != nil {
		return nil, fmt.Errorf("get object from s3 bucket: %w", err)
	}

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(object.Body)
	if err != nil {
		return nil, fmt.Errorf("reading data from s3-object: %w", err)
	}
	data := buf.Bytes()
	fmt.Printf("fetch data from %v successful. reading %v bytes from %v.\n",
		entity.Bucket.Name, len(data), entity.Object.Key)

	return data, nil
}

func getEncryptionKey(ctx context.Context) (*[32]byte, error) {
	value, err := smc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(os.Getenv("SECRET_STORE_ARM")),
	})
	if err != nil {
		return nil, fmt.Errorf("get secret value from secret vault: %w", err)
	}
	keyValue := *value.SecretString //base64 format
	decodeKeyValue, err := base64.StdEncoding.DecodeString(keyValue)
	if err != nil {
		return nil, fmt.Errorf("base64 decoding: %w", err)
	}
	key := (*[32]byte)(decodeKeyValue)
	return key, nil
}

func importData(ctx context.Context, entity events.S3Entity, data []byte) error {
	bucket := os.Getenv("TICKET_SYS_OUT_BUCKET")

	_, err := s3s.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(entity.Object.Key),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return fmt.Errorf("importing data into bucket: %w", err)
	}

	fmt.Printf("importing data to %v successful. writing %v bytes to %v.\n",
		bucket, len(data), entity.Object.Key)

	return nil
}

func main() {
	lambda.Start(handle)
}
```