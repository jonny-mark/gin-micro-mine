package sentinel

import "github.com/alibaba/sentinel-golang/logging"

// SentinelConfig represent the general configuration of Sentinel.
type SentinelConfig struct {
	App struct {
		// Name represents the name of current running service.
		Name string
		// Type indicates the classification of the service (e.g. web service, API gateway).
		Type int32
	}
	// Log represents configuration items related to logging.
	Log LogConfig
	// Stat represents configuration items related to statistics.
	Stat StatConfig
	// UseCacheTime indicates whether to cache time(ms)
	UseCacheTime bool `yaml:"useCacheTime"`
}

// LogConfig represent the configuration of logging in Sentinel.
type LogConfig struct {
	// Logger indicates that using logger to replace default logging.
	Logger logging.Logger
	// Dir represents the log directory path.
	Dir string
	// UsePid indicates whether the filename ends with the process ID (PID).
	UsePid bool `yaml:"usePid"`
	// Metric represents the configuration items of the metric log.
	Metric MetricLogConfig
}

// MetricLogConfig represents the configuration items of the metric log.
type MetricLogConfig struct {
	SingleFileMaxSize uint64 `yaml:"singleFileMaxSize"`
	MaxFileCount      uint32 `yaml:"maxFileCount"`
	FlushIntervalSec  uint32 `yaml:"flushIntervalSec"`
}

// StatConfig represents the configuration items of statistics.
type StatConfig struct {
	// GlobalStatisticSampleCountTotal and GlobalStatisticIntervalMsTotal is the per resource's global default statistic sliding window config
	GlobalStatisticSampleCountTotal uint32 `yaml:"globalStatisticSampleCountTotal"`
	GlobalStatisticIntervalMsTotal  uint32 `yaml:"globalStatisticIntervalMsTotal"`

	// MetricStatisticSampleCount and MetricStatisticIntervalMs is the per resource's default readonly metric statistic
	// This default readonly metric statistic must be reusable based on global statistic.
	MetricStatisticSampleCount uint32 `yaml:"metricStatisticSampleCount"`
	MetricStatisticIntervalMs  uint32 `yaml:"metricStatisticIntervalMs"`

	System SystemStatConfig `yaml:"system"`
}

// SystemStatConfig represents the configuration items of system statistics.
type SystemStatConfig struct {
	// CollectIntervalMs represents the collecting interval of the system metrics collector.
	CollectIntervalMs uint32 `yaml:"collectIntervalMs"`
}
