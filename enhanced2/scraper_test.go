package enhanced2

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/percona/rds_exporter/client"
	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/sessions"
)

func TestScraper(t *testing.T) {
	cfg, err := config.Load("../config.yml")
	require.NoError(t, err)
	client := client.New()
	sess, err := sessions.New(cfg.Instances, client.HTTP(), false)
	require.NoError(t, err)

	session, instance := sess.GetSession("us-east-1", "rds-aurora57")
	scraper := newScraper(session, []sessions.Instance{*instance})
	scraper.scrape(context.Background())
}
