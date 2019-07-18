#!/usr/bin/env bash

# Run this if opencensus-proto is checked in the GOPATH.
# go get -d github.com/census-instrumentation/opencensus-proto
# to check in the repo to the GOAPTH.
#
# This also requires the grpc-gateway plugin.
# See: https://github.com/grpc-ecosystem/grpc-gateway#installation
#
# To generate:
#
# cd $(go env GOPATH)/census-instrumentation/opencensus-proto
# ./mkgogogen.sh

OUTDIR="$(go env GOPATH)/src"
INCLUDES="$(go env GOPATH)/src/github.com/gogo/protobuf/protobuf"

TYPES="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types"

ARGS="goproto_registration=true"


rm -rf opencensus_gogo
cp -r opencensus opencensus_gogo
find opencensus_gogo/ -type f -exec sed -i 's/gen-go/gen-gogo/g' {} +

protoc --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/stats/v1/stats.proto \
    && protoc --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/metrics/v1/metrics.proto \
    && protoc --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/resource/v1/resource.proto \
    && protoc --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/trace/v1/trace.proto \
    && protoc --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/trace/v1/trace_config.proto \
    && protoc -I=. --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/agent/common/v1/common.proto \
    && protoc -I=. --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/agent/metrics/v1/metrics_service.proto \
    && protoc -I=. --gofast_out=$TYPES,$ARGS,plugins=grpc:$OUTDIR opencensus_gogo/proto/agent/trace/v1/trace_service.proto \
    && protoc --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./opencensus_gogo/proto/agent/trace/v1/trace_service_http.yaml:$OUTDIR opencensus_gogo/proto/agent/trace/v1/trace_service.proto
