package pinecone

import (
	"context"
	"fmt"
	"log"

	pinecone "github.com/nekomeowww/go-pinecone"
)

type PineconeClient struct {
	Client      *pinecone.Client
	IndexClient *pinecone.IndexClient
}

func NewPineconeClient(apikey, environment, projectName string) (*PineconeClient, error) {

	indexclient, err := pinecone.NewIndexClient(
		pinecone.WithIndexName("testindex"),
		pinecone.WithAPIKey(apikey),
		pinecone.WithEnvironment(environment),
		pinecone.WithProjectName(projectName),
	)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	client, err := pinecone.New(
		pinecone.WithAPIKey(apikey),
		pinecone.WithEnvironment(environment),
		pinecone.WithProjectName(projectName),
	)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &PineconeClient{
		Client:      client,
		IndexClient: indexclient,
	}, nil
}

func (pc *PineconeClient) GetVectorIndex() string {

	indexs, err := pc.Client.ListIndexes()

	if err != nil {
		panic(err)
	} else if len(indexs) > 0 {
		fmt.Printf("Database Name is : %s\n", indexs[0])
		return indexs[0]
	} else {
		fmt.Printf("Database is null\n")
		return ""
	}
}

func (pc *PineconeClient) QueryVector(ctx context.Context, vectors []float32) (string, error) {

	params := pinecone.QueryParams{
		Vector:          vectors,
		TopK:            3,
		IncludeValues:   true,
		IncludeMetadata: true,
	}

	resp, err := pc.IndexClient.Query(ctx, params)

	if err != nil {
		log.Fatal(err)
		return "", err
	} else {
		fmt.Printf("%+v\n", resp.Matches[0].Metadata["input"])
		return resp.Matches[0].Metadata["input"].(string), nil
	}
}
