package services

import (
	"context"
	"main/codegen"
)

type CoreTransformer struct {
	codegen.UnimplementedDataprocessorServer
}

func (s *CoreTransformer) GetTransformationSuggestions(ctx context.Context, req *codegen.GetTransformationSuggestionsRequest) {

}
