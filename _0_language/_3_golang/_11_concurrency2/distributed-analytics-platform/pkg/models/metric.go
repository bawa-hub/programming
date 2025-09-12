package models

import (
	"time"
)

// MetricType defines the type of metric
type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
	MetricTypeSummary   MetricType = "summary"
)

// Metric represents a collected metric
type Metric struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Value      float64                `json:"value"`
	Type       MetricType             `json:"type"`
	Timestamp  time.Time              `json:"timestamp"`
	Source     string                 `json:"source,omitempty"`
	Dimensions map[string]interface{} `json:"dimensions,omitempty"`
}

// NewMetric creates a new metric
func NewMetric(name string, value float64, metricType MetricType) *Metric {
	return &Metric{
		ID:        generateID(),
		Name:      name,
		Value:     value,
		Type:      metricType,
		Timestamp: time.Now(),
	}
}

// SetDimensions sets dimensions for the metric
func (m *Metric) SetDimensions(dimensions map[string]interface{}) *Metric {
	m.Dimensions = dimensions
	return m
}

// SetSource sets the source for the metric
func (m *Metric) SetSource(source string) *Metric {
	m.Source = source
	return m
}

// AggregatedMetric represents an aggregated metric
type AggregatedMetric struct {
	Name        string                 `json:"name"`
	Type        MetricType             `json:"type"`
	Count       int                    `json:"count"`
	Sum         float64                `json:"sum"`
	Min         float64                `json:"min"`
	Max         float64                `json:"max"`
	Avg         float64                `json:"avg"`
	StartTime   time.Time              `json:"start_time"`
	EndTime     time.Time              `json:"end_time"`
	Dimensions  map[string]interface{} `json:"dimensions,omitempty"`
}

// NewAggregatedMetric creates a new aggregated metric
func NewAggregatedMetric(name string, metricType MetricType, startTime, endTime time.Time) *AggregatedMetric {
	return &AggregatedMetric{
		Name:      name,
		Type:      metricType,
		Count:     0,
		Sum:       0,
		Min:       0,
		Max:       0,
		Avg:       0,
		StartTime: startTime,
		EndTime:   endTime,
		Dimensions: make(map[string]interface{}),
	}
}

// SetDimensions sets dimensions for the aggregated metric
func (am *AggregatedMetric) SetDimensions(dimensions map[string]interface{}) *AggregatedMetric {
	am.Dimensions = dimensions
	return am
}

// AddValue adds a value to the aggregated metric
func (am *AggregatedMetric) AddValue(value float64) {
	am.Count++
	am.Sum += value
	
	if am.Count == 1 {
		am.Min = value
		am.Max = value
	} else {
		if value < am.Min {
			am.Min = value
		}
		if value > am.Max {
			am.Max = value
		}
	}
	
	am.Avg = am.Sum / float64(am.Count)
}
