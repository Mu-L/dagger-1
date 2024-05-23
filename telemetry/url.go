package telemetry

import (
	"context"
	"fmt"
	"os"

	"github.com/dagger/dagger/internal/cloud/auth"
	"go.opentelemetry.io/otel/trace"
)

func URLForTrace(ctx context.Context) (string, bool) {
	if !configuredCloudTelemetry {
		return "", false
	}

	var orgName string
	if cloudToken := os.Getenv("DAGGER_CLOUD_TOKEN"); cloudToken != "" {
		// Try token auth first
		token, ok := parseDaggerToken(cloudToken)
		if ok {
			orgName = token.orgName
		}
	} else {
		// Try OAuth next
		org, err := auth.CurrentOrg()
		if err != nil {
			return "", false
		}
		orgName = org.Name
	}

	url := fmt.Sprintf(
		"https://dagger.cloud/%s/traces/%s",
		orgName,
		trace.SpanContextFromContext(ctx).TraceID().String(),
	)
	return url, true
}
