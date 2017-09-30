package main
import (
	"google.golang.org/api/monitoring/v3"
	"fmt"
	"time"
	"log"
)

// writeTimeSeriesValue writes a value for the custom metric created
func WriteTimeSeriesMatricsRequestCount(s *monitoring.Service, projectID string, metricType string,aspect CounterAspect) error {

	l :=len(aspect.Counter.Requests)
 	var timeserieslist = make([]*monitoring.TimeSeries,l)
 	i :=0
	for k, v := range aspect.Counter.Requests {
		count :=float64(v)
		now := time.Now().UTC().Format(time.RFC3339)

		timeserieslist[i]  = &monitoring.TimeSeries{
			Metric: &monitoring.Metric{
				Type: metricType,
				Labels: map[string]string{
					"api_name": k,
				},
			},
			Resource: &monitoring.MonitoredResource{
				Labels: map[string]string{
					"instance_id": "test-instancoe",
					"zone":        "us-central1-f",
				},
				Type: "gce_instance",
			},
			Points: []*monitoring.Point{
				{
					Interval: &monitoring.TimeInterval{
						//StartTime: now,
						EndTime: now,
					},
					Value: &monitoring.TypedValue{
						DoubleValue: &count,
					},
				},
			},
		}
		i++
	}
	createTimeseriesRequest := monitoring.CreateTimeSeriesRequest{
		TimeSeries: timeserieslist,
	}

	log.Printf("writeTimeseriesRequest: %s\n", formatResource(createTimeseriesRequest))
	_, err := s.Projects.TimeSeries.Create(projectResource(projectID), &createTimeseriesRequest).Do()
	if err != nil {
		return fmt.Errorf("Could not write time series value, %v ", err)
	}
	return nil
}

//Responsecode write

func WriteTimeSeriesMatricsResCodeCount(s *monitoring.Service, projectID string, metricType string,aspect RequestType) error {

	l :=len(aspect.RequestCodeCount)
	var timeserieslist = make([]*monitoring.TimeSeries,l)

	i:=0
	for api,code := range aspect.RequestCodeCount{

		count :=float64(code)
		now := time.Now().UTC().Format(time.RFC3339)

		timeserieslist[i]  = &monitoring.TimeSeries{
			Metric: &monitoring.Metric{
				Type: metricType,
				Labels: map[string]string{
					"api_name": api,
				},
			},
			Resource: &monitoring.MonitoredResource{
				Labels: map[string]string{
					"instance_id": "test-instance",
					"zone":        "us-central1-f",
				},
				Type: "gce_instance",
			},
			Points: []*monitoring.Point{
				{
					Interval: &monitoring.TimeInterval{
						//StartTime: now,
						EndTime: now,
					},
					Value: &monitoring.TypedValue{
						DoubleValue: &count,
					},
				},
			},
		}
		i++

	}
	createTimeseriesRequest := monitoring.CreateTimeSeriesRequest{
		TimeSeries: timeserieslist,//[]*monitoring.TimeSeries{&timeseries1,&timeseries2....}
	}

	log.Printf("writeTimeseriesRequest: %s\n", formatResource(createTimeseriesRequest))
	_, err := s.Projects.TimeSeries.Create(projectResource(projectID), &createTimeseriesRequest).Do()
	if err != nil {
		return fmt.Errorf("Could not write time series value for request Code, %v ", err)
	}
	return nil
}


func WriteTimeSeriesMatricsRequestTime(s *monitoring.Service, projectID string, metricType string,aspect RequestTimeAspect) error {

	l :=len(aspect.Request.AvgTime)
	var timeserieslist = make([]*monitoring.TimeSeries,l)

	i:=0
	for api,code := range aspect.Request.AvgTime{

		count :=float64(code)
		now := time.Now().UTC().Format(time.RFC3339)

		timeserieslist[i]  = &monitoring.TimeSeries{
			Metric: &monitoring.Metric{
				Type: metricType,
				Labels: map[string]string{
					"api_name": api,
				},
			},
			Resource: &monitoring.MonitoredResource{
				Labels: map[string]string{
					"instance_id": "test-instance",
					"zone":        "us-central1-f",
				},
				Type: "gce_instance",
			},
			Points: []*monitoring.Point{
				{
					Interval: &monitoring.TimeInterval{
						EndTime: now,
					},
					Value: &monitoring.TypedValue{
						DoubleValue: &count,
					},
				},
			},
		}
		i++

	}

	createTimeseriesRequest := monitoring.CreateTimeSeriesRequest{
		TimeSeries: timeserieslist,//[]*monitoring.TimeSeries{&timeseries1,&timeseries2....}
	}

	log.Printf("writeTimeseriesRequest: %s\n", formatResource(createTimeseriesRequest))
	_, err := s.Projects.TimeSeries.Create(projectResource(projectID), &createTimeseriesRequest).Do()
	if err != nil {
		return fmt.Errorf("Could not write time series value for request Code, %v ", err)
	}
	return nil
}