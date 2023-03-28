package app

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/wait"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/testhelp"
)

// Useful artical to help understand how this test operates
// https://medium.com/@manabie/test-coverage-of-go-services-during-integration-tests-6ff1bdbe33e0

// This test also sets up and runs its own postgres testing container
func TestAap(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	killServer := testhelp.NewKillServer(":19999", cancel)
	go killServer.Start()

	const dbname = "test-db"
	const user = "postgres"
	const password = "password"

	port, err := nat.NewPort("tcp", "5432")
	require.NoError(t, err)

	container, err := testhelp.StarPostgresContainer(ctx,
		testhelp.WithPort(port.Port()),
		testhelp.WithInitialDatabase(user, password, dbname),
		testhelp.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("failed to terminate container: %s", err)
		}
	})

	containerPort, err := container.MappedPort(ctx, port)
	assert.NoError(t, err)

	host, err := container.Host(ctx)
	assert.NoError(t, err)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, containerPort.Port(), user, password, dbname)

	Run(&domain.Config{
		App: domain.App{},
		HTTP: domain.HTTP{
			Port:            "8082",
			CookieStoreKey:  "NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==",
			SessionStoreKey: "AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==",
		},
		LOG: domain.LOG{
			Format:  "json",
			Level:   "info",
			NoColor: true,
		},
		PG: domain.PG{
			PoolMax: 0,
			URL:     connStr,
		},
	}, ctx)

	killServer.Shutdown(context.Background())
}
