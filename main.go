package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-analytics-plugin] Starting Analytics plugin...")
	plugin := sdk.Init("hc-analytics-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("AnalyticsStatus", "Analytics plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getAnalyticsStatus",
		sdk.ComplexObjectField("Get analytics plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-analytics-plugin] Plugin ready")
	plugin.Serve()
}
