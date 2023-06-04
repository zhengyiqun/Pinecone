package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	tp "testapi/pinecone"
)

type VectorIndexBody struct {
	IndexName string `json:"indexname"`
}

type QueryParams struct {
	Embedding []float32 `json:"embedding"`
}

type QueryBody struct {
	Input string `json:"input"`
}

// @Summary GetIndexName
// @Description Get Pinecone Index Name
// @Success 200 {string} string	"ok"
// @Router /api/v1/index [GET][个人测试]
func GetIndexName(w http.ResponseWriter, r *http.Request) {

	_apikey := os.Getenv("PINECONE_APIKEY")
	_environment := os.Getenv("PINECONE_ENVIRONMENT")
	_projectname := os.Getenv("PINECONE_PROJECT_NAME")

	config, err := tp.NewPineconeClient(_apikey, _environment, _projectname)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	indexname := config.GetVectorIndex()

	VectorIndexBody := VectorIndexBody{
		IndexName: indexname,
	}
	VectorIndexBodyJson, err := json.Marshal(VectorIndexBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(VectorIndexBodyJson)
}

// @Summary QueryIndex
// @Description Query Pinecone Index
// @Success 200 {string} string	"ok"
// @Router /api/v1/query [GET][个人测试]
func QueryIndex(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	_apikey := os.Getenv("PINECONE_APIKEY")
	_environment := os.Getenv("PINECONE_ENVIRONMENT")
	_projectname := os.Getenv("PINECONE_PROJECT_NAME")

	var params QueryParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	config, err := tp.NewPineconeClient(_apikey, _environment, _projectname)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	input, _ := config.QueryVector(ctx, params.Embedding)

	QueryBody := QueryBody{
		Input: input,
	}
	VectorIndexBodyJson, err := json.Marshal(QueryBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(VectorIndexBodyJson)
}
