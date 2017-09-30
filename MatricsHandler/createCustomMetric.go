package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/monitoring/v3"
	"log"
	"fmt"
	"time"
)

// PRE-REQUISITES:
// ---------------
// 1. If not already done, enable the Google Monitoring API and check the quota for your project at
//    https://console.developers.google.com/apis/api/monitoring_component/quotas
// 2. This sample uses Application Default Credentials for Auth. If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run 'gcloud beta auth application-default login'
// 3. To install the client library and Application Default Credentials library, run:
//    'go get google.golang.org/api/monitoring/v3'
//    'go get golang.org/x/oauth2/google'



func CreateService(ctx context.Context) (*monitoring.Service, error) {
	hc, err := google.DefaultClient(ctx, monitoring.MonitoringScope)
	if err != nil {
		return nil, err
	}
	s, err := monitoring.New(hc)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func projectResource(projectID string) string {
	return "projects/" + projectID
}

// formatResource marshals a response object as JSON.
func formatResource(resource interface{}) []byte {
	b, err := json.MarshalIndent(resource, "", "    ")
	if err != nil {
		panic(err)
	}
	return b
}

// getCustomMetric reads the custom metric created.
func GetCustomMetric(s *monitoring.Service, projectID, metricType string) (*monitoring.ListMetricDescriptorsResponse, error) {
	resp, err := s.Projects.MetricDescriptors.List(projectResource(projectID)).
		Filter(fmt.Sprintf("metric.type=\"%s\"", metricType)).Do()
	if err != nil {
		return nil, fmt.Errorf("Could not get custom metric: %v", err)
	}

	log.Printf("getCustomMetric: %s\n", formatResource(resp))
	return resp, nil
}

func CreateCustomMetric(s *monitoring.Service, projectID, metricType string,matricsName string,matricsDescription string) error {
	ld := monitoring.LabelDescriptor{Key: "environment", ValueType: "STRING", Description: "An arbitrary measurement"}
	md := monitoring.MetricDescriptor{
		Type:        metricType,
		Labels:      []*monitoring.LabelDescriptor{&ld},
		MetricKind:  "GAUGE",
		ValueType:   "DOUBLE",
		Unit:        "items",
		Description: matricsDescription,
		DisplayName: matricsName,
	}


	resp, err := s.Projects.MetricDescriptors.Create(projectResource(projectID), &md).Do()

	if err != nil {
		return fmt.Errorf("Could not create custom metric: %v", err)
	}

	log.Printf("createCustomMetric: %s\n", formatResource(resp))

	// Wait until the new metric can be read back.
	for {
		resp, err := GetCustomMetric(s, projectID, metricType)
		if err != nil {
			log.Fatal(err)
		}
		if len(resp.MetricDescriptors) != 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}





//test file
/*
func main_1() {

rand.Seed(time.Now().UTC().UnixNano())

ctx := context.Background()
serv, err := CreateService(ctx)
if err != nil {
	log.Fatal(err)
}

if err := CreateCustomMetric(serv, projectID, metricType); err != nil {
	log.Fatal(err)
}


	if err := WriteTimeSeriesValue(serv, projectID, metricType); err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)


time.Sleep(2 * time.Second)

// Read the TimeSeries for the last 5 minutes for that metric.
f err := readTimeSeriesValue(serv, projectID, metricType); err != nil {
	log.Fatal(err)
}
*/

