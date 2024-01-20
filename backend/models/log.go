package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Level            string    `json:"level" gorm:"index:idx_level"`
	Message          string    `json:"message"`
	ResourceID       string    `json:"resourceId" gorm:"index:idx_resource_id"`
	Timestamp        time.Time `json:"timestamp" gorm:"index:idx_timestamp"`
	TraceID          string    `json:"traceId" gorm:"index:idx_trace_id"`
	SpanID           string    `json:"spanId" gorm:"index:idx_span_id"`
	Commit           string    `json:"commit"`
	ParentResourceID string    `json:"parentResourceId" gorm:"index:idx_parent_resource_id;column:parent_resource_id"`
}
