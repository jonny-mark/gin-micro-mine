/**
 * @author jiangshangfang
 * @date 2022/4/15 12:00 AM
 **/
package sql

import (
	"database/sql"
	"go.opentelemetry.io/otel/trace"
	"context"
)

// tx 事务
type Tx struct {
	tx    *sql.Tx
	trace trace.Tracer
}

// Commit事务
func (tx *Tx) Commit(ctx context.Context) (err error) {
	err = tx.tx.Commit()
	if tx.trace != nil {
		ctx, span := tx.trace.Start(ctx,"tx.Commit")
		defer span.End()

	}
	return
}
