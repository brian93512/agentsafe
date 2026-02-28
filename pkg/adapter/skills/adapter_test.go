package skills_test

import (
	"context"
	"testing"

	"github.com/brian93512/agentsafe/pkg/adapter/skills"
	"github.com/brian93512/agentsafe/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestAdapter_Protocol(t *testing.T) {
	a := skills.NewAdapter()
	assert.Equal(t, model.ProtocolSkills, a.Protocol())
}

func TestAdapter_Parse_NotImplemented(t *testing.T) {
	t.Skip("Skills adapter not yet implemented")
	_, err := skills.NewAdapter().Parse(context.Background(), []byte(""))
	assert.NoError(t, err)
}
