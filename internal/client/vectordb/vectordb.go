package vectordb

import (
	"context"
	"fmt"

	"patient-chatbot/internal/config"

	"github.com/pinecone-io/go-pinecone/v4/pinecone"
)

type VectordbClient struct {
	idxConnection *pinecone.IndexConnection
	namespace     string
}

func NewVectordbClient(cfg *config.Config) (*VectordbClient, error) {
	pc, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: cfg.PineconeAPIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("pinecone NewClient: %w", err)
	}

	conn, err := pc.Index(pinecone.NewIndexConnParams{
		Host:      cfg.PineconeHost,
		Namespace: cfg.PineconeNamespace,
	})
	if err != nil {
		return nil, fmt.Errorf("IndexConnection: %w", err)
	}

	return &VectordbClient{
		idxConnection: conn,
		namespace:     cfg.PineconeNamespace,
	}, nil
}

func (v *VectordbClient) Search(ctx context.Context, userQuery string, topK ...int) (*pinecone.SearchRecordsResponse, error) {
	var k int = 5
	if len(topK) > 0 {
		k = topK[0]
	}
	topN := int32(2)
	searchWithText, err := v.idxConnection.SearchRecords(ctx, &pinecone.SearchRecordsRequest{
		Query: pinecone.SearchRecordsQuery{
			TopK: int32(k),
			Inputs: &map[string]interface{}{
				"text": userQuery,
			},
		},
		Rerank: &pinecone.SearchRecordsRerank{
			Model:      "bge-reranker-v2-m3",
			TopN:       &topN,
			RankFields: []string{"chunk_text"},
		},
		Fields: &[]string{"chunk_text", "category"},
	})
	if err != nil {
		return nil, fmt.Errorf("SearchRecords: %w", err)
	}
	return searchWithText, nil

}

func (v *VectordbClient) CreateChunks(ctx context.Context, records []*pinecone.IntegratedRecord) error {
	err := v.idxConnection.UpsertRecords(ctx, records)
	if err != nil {
		return fmt.Errorf("UpsertRecords: %w", err)
	}
	return nil
}

func (v *VectordbClient) DeleteChunks(ctx context.Context, ids []string) error {
	err := v.idxConnection.DeleteVectorsById(ctx, ids)
	if err != nil {
		return fmt.Errorf("DeleteVectorsById: %w", err)
	}
	return nil
}
