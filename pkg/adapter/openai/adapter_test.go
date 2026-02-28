package openai_test

import (
	"context"
	"testing"

	"github.com/brian93512/agentsafe/pkg/adapter/openai"
	"github.com/brian93512/agentsafe/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestAdapter_Protocol(t *testing.T) {
	a := openai.NewAdapter()
	assert.Equal(t, model.ProtocolOpenAI, a.Protocol())
}

func TestAdapter_Parse_NotImplemented(t *testing.T) {
	t.Skip("OpenAI adapter not yet implemented")
	_, err := openai.NewAdapter().Parse(context.Background(), []byte("{}"))
	assert.NoError(t, err)
}
