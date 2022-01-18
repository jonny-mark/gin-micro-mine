/**
 * @author jiangshangfang
 * @date 2022/1/10 11:39 AM
 **/
package repository

import (
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"go.opentelemetry.io/otel"
)

var _ Repository = (*repository)(nil)

// Repository 定义用户仓库接口
type Repository interface {
}

// repository mysql struct
type repository struct {
	orm *gorm.DB
	tracer    trace.Tracer
}

func New(db *gorm.DB) Repository {
	return &repository{
		orm:       db,
		tracer:    otel.Tracer("repository"),
	}
}
