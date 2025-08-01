module go.opentelemetry.io/collector/otelcol/otelcoltest

go 1.23.0

require (
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v1.37.0
	go.opentelemetry.io/collector/confmap v1.37.0
	go.opentelemetry.io/collector/confmap/provider/envprovider v1.37.0
	go.opentelemetry.io/collector/confmap/provider/fileprovider v1.37.0
	go.opentelemetry.io/collector/confmap/provider/httpprovider v1.37.0
	go.opentelemetry.io/collector/confmap/provider/yamlprovider v1.37.0
	go.opentelemetry.io/collector/confmap/xconfmap v0.131.0
	go.opentelemetry.io/collector/connector/connectortest v0.131.0
	go.opentelemetry.io/collector/exporter/exportertest v0.131.0
	go.opentelemetry.io/collector/extension/extensiontest v0.131.0
	go.opentelemetry.io/collector/otelcol v0.131.0
	go.opentelemetry.io/collector/pipeline v0.131.0
	go.opentelemetry.io/collector/processor/processortest v0.131.0
	go.opentelemetry.io/collector/receiver/receivertest v0.131.0
	go.opentelemetry.io/collector/service v0.131.0
	go.uber.org/goleak v1.3.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/providers/confmap v1.0.0 // indirect
	github.com/knadh/koanf/v2 v2.2.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.22.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.65.0 // indirect
	github.com/prometheus/procfs v0.17.0 // indirect
	github.com/shirou/gopsutil/v4 v4.25.6 // indirect
	github.com/spf13/cobra v1.9.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.131.0 // indirect
	go.opentelemetry.io/collector/component/componenttest v0.131.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.131.0 // indirect
	go.opentelemetry.io/collector/connector v0.131.0 // indirect
	go.opentelemetry.io/collector/connector/xconnector v0.131.0 // indirect
	go.opentelemetry.io/collector/consumer v1.37.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.131.0 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.131.0 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.131.0 // indirect
	go.opentelemetry.io/collector/exporter v0.131.0 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.131.0 // indirect
	go.opentelemetry.io/collector/extension v1.37.0 // indirect
	go.opentelemetry.io/collector/extension/extensioncapabilities v0.131.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.37.0 // indirect
	go.opentelemetry.io/collector/internal/fanoutconsumer v0.131.0 // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.131.0 // indirect
	go.opentelemetry.io/collector/pdata v1.37.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.131.0 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.131.0 // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.131.0 // indirect
	go.opentelemetry.io/collector/processor v1.37.0 // indirect
	go.opentelemetry.io/collector/processor/xprocessor v0.131.0 // indirect
	go.opentelemetry.io/collector/receiver v1.37.0 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.131.0 // indirect
	go.opentelemetry.io/collector/service/hostcapabilities v0.131.0 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.12.0 // indirect
	go.opentelemetry.io/contrib/otelconf v0.16.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.36.0 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.58.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.13.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.37.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.37.0 // indirect
	go.opentelemetry.io/otel/log v0.13.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
	go.opentelemetry.io/otel/sdk v1.37.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.13.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.37.0 // indirect
	go.opentelemetry.io/proto/otlp v1.7.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	gonum.org/v1/gonum v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/grpc v1.74.2 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.5.0 // indirect
)

replace go.opentelemetry.io/collector/receiver => ../../receiver

replace go.opentelemetry.io/collector/consumer => ../../consumer

replace go.opentelemetry.io/collector/confmap/provider/httpprovider => ../../confmap/provider/httpprovider

replace go.opentelemetry.io/collector/confmap/provider/yamlprovider => ../../confmap/provider/yamlprovider

replace go.opentelemetry.io/collector => ../..

replace go.opentelemetry.io/collector/featuregate => ../../featuregate

replace go.opentelemetry.io/collector/config/configtelemetry => ../../config/configtelemetry

replace go.opentelemetry.io/collector/processor => ../../processor

replace go.opentelemetry.io/collector/pdata => ../../pdata

replace go.opentelemetry.io/collector/pdata/pprofile => ../../pdata/pprofile

replace go.opentelemetry.io/collector/connector => ../../connector

replace go.opentelemetry.io/collector/connector/connectortest => ../../connector/connectortest

replace go.opentelemetry.io/collector/config/configretry => ../../config/configretry

replace go.opentelemetry.io/collector/confmap/provider/fileprovider => ../../confmap/provider/fileprovider

replace go.opentelemetry.io/collector/confmap/provider/envprovider => ../../confmap/provider/envprovider

replace go.opentelemetry.io/collector/component => ../../component

replace go.opentelemetry.io/collector/component/componenttest => ../../component/componenttest

replace go.opentelemetry.io/collector/extension => ../../extension

replace go.opentelemetry.io/collector/exporter => ../../exporter

replace go.opentelemetry.io/collector/consumer/xconsumer => ../../consumer/xconsumer

replace go.opentelemetry.io/collector/consumer/consumertest => ../../consumer/consumertest

replace go.opentelemetry.io/collector/component/componentstatus => ../../component/componentstatus

replace go.opentelemetry.io/collector/extension/extensioncapabilities => ../../extension/extensioncapabilities

replace go.opentelemetry.io/collector/receiver/xreceiver => ../../receiver/xreceiver

replace go.opentelemetry.io/collector/receiver/receivertest => ../../receiver/receivertest

replace go.opentelemetry.io/collector/processor/xprocessor => ../../processor/xprocessor

replace go.opentelemetry.io/collector/connector/xconnector => ../../connector/xconnector

replace go.opentelemetry.io/collector/exporter/xexporter => ../../exporter/xexporter

replace go.opentelemetry.io/collector/pipeline => ../../pipeline

replace go.opentelemetry.io/collector/pipeline/xpipeline => ../../pipeline/xpipeline

replace go.opentelemetry.io/collector/exporter/exportertest => ../../exporter/exportertest

replace go.opentelemetry.io/collector/consumer/consumererror => ../../consumer/consumererror

replace go.opentelemetry.io/collector/config/configopaque => ../../config/configopaque

replace go.opentelemetry.io/collector/config/configauth => ../../config/configauth

replace go.opentelemetry.io/collector/config/configtls => ../../config/configtls

replace go.opentelemetry.io/collector/config/configcompression => ../../config/configcompression

replace go.opentelemetry.io/collector/config/confighttp => ../../config/confighttp

replace go.opentelemetry.io/collector/client => ../../client

replace go.opentelemetry.io/collector/extension/extensionauth => ../../extension/extensionauth

replace go.opentelemetry.io/collector/otelcol => ../

replace go.opentelemetry.io/collector/extension/zpagesextension => ../../extension/zpagesextension

replace go.opentelemetry.io/collector/service => ../../service

replace go.opentelemetry.io/collector/service/hostcapabilities => ../../service/hostcapabilities

replace go.opentelemetry.io/collector/internal/fanoutconsumer => ../../internal/fanoutconsumer

replace go.opentelemetry.io/collector/extension/extensiontest => ../../extension/extensiontest

replace go.opentelemetry.io/collector/extension/extensionauth/extensionauthtest => ../../extension/extensionauth/extensionauthtest

replace go.opentelemetry.io/collector/extension/xextension => ../../extension/xextension

replace go.opentelemetry.io/collector/confmap => ../../confmap

replace go.opentelemetry.io/collector/confmap/xconfmap => ../../confmap/xconfmap

replace go.opentelemetry.io/collector/processor/processortest => ../../processor/processortest

replace go.opentelemetry.io/collector/pdata/testdata => ../../pdata/testdata

replace go.opentelemetry.io/collector/internal/telemetry => ../../internal/telemetry

replace go.opentelemetry.io/collector/config/configmiddleware => ../../config/configmiddleware

replace go.opentelemetry.io/collector/extension/extensionmiddleware/extensionmiddlewaretest => ../../extension/extensionmiddleware/extensionmiddlewaretest

replace go.opentelemetry.io/collector/extension/extensionmiddleware => ../../extension/extensionmiddleware

replace go.opentelemetry.io/collector/config/configoptional => ../../config/configoptional

replace go.opentelemetry.io/collector/pdata/xpdata => ../../pdata/xpdata
