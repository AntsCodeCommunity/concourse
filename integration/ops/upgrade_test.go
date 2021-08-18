package ops_test

import (
	"testing"

	"github.com/concourse/concourse/integration/internal/dctest"
	"github.com/concourse/concourse/integration/internal/flytest"
)

func TestUpgrade(t *testing.T) {
	t.Skip("See https://github.com/concourse/concourse/issues/7397")

	t.Parallel()

	latestDC := dctest.Init(t, "../docker-compose.yml", "overrides/named.yml", "overrides/latest.yml")

	t.Run("deploy latest", func(t *testing.T) {
		latestDC.Run(t, "up", "-d")
	})

	fly := flytest.Init(t, latestDC)
	setupUpgradeDowngrade(t, fly)

	devDC := dctest.Init(t, "../docker-compose.yml", "overrides/named.yml")

	t.Run("upgrade to dev", func(t *testing.T) {
		devDC.Run(t, "up", "-d")
	})

	fly = flytest.Init(t, devDC)
	verifyUpgradeDowngrade(t, fly)
}
