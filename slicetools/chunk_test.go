package slicetools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	tCases := map[string]struct {
		inputSize      int
		inputSlice     []interface{}
		expectedChunks [][]interface{}
	}{
		"should chunk into 3": {
			inputSize:      3,
			inputSlice:     []any{nil, nil, nil, nil, nil},
			expectedChunks: [][]any{{nil, nil, nil}, {nil, nil}},
		},
		"should chunk into 2": {
			inputSize:      2,
			inputSlice:     []any{nil, nil, nil, nil, nil, nil, nil},
			expectedChunks: [][]any{{nil, nil}, {nil, nil}, {nil, nil}, {nil}},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			actual := Chunk(tc.inputSize, tc.inputSlice)
			assert.Equal(tt, len(tc.expectedChunks), len(actual))
			for i := 0; i < len(actual); i++ {
				assert.Equal(tt, len(tc.expectedChunks[i]), len(actual[i]))
			}
		})
	}
}

func TestChunkFor(t *testing.T) {
	tCases := map[string]struct {
		inputSize      int
		inputSlice     []interface{}
		expectedChunks [][]interface{}
	}{
		"should works for 3": {
			inputSize:      3,
			inputSlice:     []any{nil, nil, nil, nil, nil},
			expectedChunks: [][]any{{nil, nil}, {nil, nil}, {nil}},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			actual := ChunkFor(tc.inputSlice, tc.inputSize)
			assert.Equal(tt, len(tc.expectedChunks), len(actual))
			for i := 0; i < len(actual); i++ {
				assert.Equal(tt, len(tc.expectedChunks[i]), len(actual[i]))
			}
		})
	}
}
