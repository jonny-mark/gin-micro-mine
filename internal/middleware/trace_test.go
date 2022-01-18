/**
 * @author jiangshangfang
 * @date 2021/11/6 4:58 PM
 **/
package middleware

import (
	"testing"
	"go.opentelemetry.io/otel"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"go.opentelemetry.io/otel/oteltest"
	oteltrace "go.opentelemetry.io/otel/trace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
)

// nolint
func init() {
	gin.SetMode(gin.DebugMode) // silence annoying log msgs
}

func TestChildSpanFromGlobalTracer(t *testing.T) {
	otel.SetTracerProvider(oteltest.NewTracerProvider())

	router := gin.New()
	router.Use(Tracing("foobar"))
	router.GET("/user/:id", func(c *gin.Context) {
		span := oteltrace.SpanFromContext(c.Request.Context())
		_, ok := span.(*oteltest.Span)
		assert.True(t, ok)
	})

	r := httptest.NewRequest("GET", "/user/123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)
}

func TestTrace200(t *testing.T) {
	sr := new(oteltest.SpanRecorder)
	provider := oteltest.NewTracerProvider(oteltest.WithSpanRecorder(sr))

	router := gin.New()
	router.Use(Tracing("foobar", WithTracerProvider(provider)))
	router.GET("/user/:id", func(c *gin.Context) {
		span := oteltrace.SpanFromContext(c.Request.Context())
		mspan, ok := span.(*oteltest.Span)
		require.True(t, ok)
		assert.Equal(t, attribute.StringValue("foobar"), mspan.Attributes()["http.server_name"])
		id := c.Param("id")
		_, _ = c.Writer.Write([]byte(id))
	})

	r := httptest.NewRequest("GET", "/user/123", nil)
	w := httptest.NewRecorder()

	// do and verify the request
	router.ServeHTTP(w, r)
	response := w.Result()
	require.Equal(t, http.StatusOK, response.StatusCode)

	// verify traces look good
	spans := sr.Completed()
	require.Len(t, spans, 1)
	span := spans[0]
	assert.Equal(t, "/user/:id", span.Name())
	assert.Equal(t, oteltrace.SpanKindServer, span.SpanKind())
	assert.Equal(t, attribute.StringValue("foobar"), span.Attributes()["http.server_name"])
	assert.Equal(t, attribute.IntValue(http.StatusOK), span.Attributes()["http.status_code"])
	assert.Equal(t, attribute.StringValue("GET"), span.Attributes()["http.method"])
	assert.Equal(t, attribute.StringValue("/user/123"), span.Attributes()["http.target"])
	assert.Equal(t, attribute.StringValue("/user/:id"), span.Attributes()["http.route"])
}