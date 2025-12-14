package variation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVariationProvider_HappyPath(t *testing.T) {
	t.Parallel()

	for name := range registry {
		assert.True(t, name.IsValid())

		variation, err := Provider(name)

		require.NoError(t, err)
		assert.NotNil(t, variation)
	}
}

func TestVariationProvider_UnknownVariation(t *testing.T) {
	t.Parallel()

	name := Name("unknown")
	variation, err := Provider(name)

	require.Equal(t, NamedFunction{}, variation)
	require.Error(t, err)
	assert.ErrorContains(t, err, "unknown variation")
}
