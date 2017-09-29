package main1

// PRE-REQUISITES:
// ---------------
// 1. If not already done, enable the Google Monitoring API and check the quota for your project at
//    https://console.developers.google.com/apis/api/monitoring_component/quotas
// 2. This sample uses Application Default Credentials for Auth. If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run 'gcloud beta auth application-default login'
// 3. To install the client library and Application Default Credentials library, run:
//    'go get google.golang.org/api/monitoring/v3'
//    'go get golang.org/x/oauth2/google'

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/monitoring/v3"
)

func main() {
	ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, monitoring.CloudPlatformScope)
	if err != nil {
		// TODO: Handle error.
		_ = err
	}
	client, err := monitoring.New(httpClient)
	if err != nil {
		// TODO: Handle error.
		_ = err
	}

	// TODO: Change placeholders below to appropriate parameter values for the 'create' method:
	var (
		// The project on which to execute the request. The format is `"projects/{project_id_or_number}"`.
		name = ""

		requestBody = &monitoring.MetricDescriptor{}
	)

	response, err := client.Projects.MetricDescriptors.Create(name, requestBody).Context(ctx).Do()
	if err != nil {
		// TODO: Handle error.
		_ = err
	}
	// doThingsWith(response)
	_ = response
}