/**
 * @author jiangshangfang
 * @date 2022/2/21 4:39 PM
 **/
package issuer

import (
	"context"
	"gin/internal/model"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func (r *repository) GetIssuerOrderInPush(ctx context.Context, status int) (*model.IssuerOrderModel, error) {
	ctx, span := r.tracer.Start(ctx, "GetIssuerOrderInPush", oteltrace.WithAttributes(attribute.String("param.status", string(status))))
	defer span.End()
	//r.issuerOrderCache.Ge

	var data *model.IssuerOrderModel

	return data, nil
}
