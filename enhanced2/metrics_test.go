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
	assert.Equal(t, time.Date(2018, 9, 25, 8, 7, 3, 0, time.UTC), m.Timestamp)

	metrics := m.originalMetrics("us-east-1")
	actual := strings.Join(helpers.Format(metrics), "\n")
	assert.Equal(t, dataMySQL57Expected, actual, "Actual:\n%s", actual)

	m, err = parseOSMetrics(dataAurora57)
	require.NoError(t, err)
	assert.Equal(t, time.Date(2018, 9, 25, 8, 16, 20, 0, time.UTC), m.Timestamp)

	metrics = m.originalMetrics("us-east-1")
	actual = strings.Join(helpers.Format(metrics), "\n")
	assert.Equal(t, dataAurora57Expected, actual, "Actual:\n%s", actual)
}
