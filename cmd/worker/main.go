package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Ricardolv/drive-app/internal/bucket"
	"github.com/Ricardolv/drive-app/internal/queue"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {

	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	qc, err := queue.New(qcfg, queue.RABBITMQ)
	if err != nil {
		panic(err)
	}

	chanel := make(chan queue.QueueResponse)
	qc.Consume(chanel)

	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region:      aws.String("AWS_REGION"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY_PWD"), os.Getenv("AWS_SECRET"), ""),
		},
		BucketDownload: "drive-app-raw",
		BucketUpload:   "drive-app-gzip",
	}

	buck, err := bucket.New(bcfg, bucket.AwsProvider)
	if err != nil {
		panic(err)
	}

	for msg := range chanel {
		src := fmt.Sprintf("%s/%s", msg.Path, msg.Filename)
		dst := fmt.Sprintf("%d_%s", msg.ID, msg.Filename)

		file, err := buck.Download(src, dst)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		body, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		if _, err := zw.Write(body); err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		if err := zw.Close(); err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		err = buck.Upload(zr, src)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		err = os.Remove(dst)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

	}

}
