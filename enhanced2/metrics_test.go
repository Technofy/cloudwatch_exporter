package enhanced2

import (
	"strings"
	"testing"
	"time"

	"github.com/percona/exporter_shared/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:lll
func TestParse(t *testing.T) {
	m, err := parseOSMetrics(dataMySQL57)
	require.NoError(t, err)
	timestamp := time.Date(2018, 1, 8, 14, 35, 21, 0, time.UTC)
	assert.Equal(t, timestamp, m.Timestamp)

	metrics := m.originalMetrics("us-east-1")
	assert.Len(t, metrics, expectedMetrics)
	actual := strings.Join(helpers.Format(metrics), "\n")

	assert.Equal(t, dataMySQL57Expected, actual)
}
