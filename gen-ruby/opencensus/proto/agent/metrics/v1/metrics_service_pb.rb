# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: opencensus/proto/agent/metrics/v1/metrics_service.proto

require 'google/protobuf'

require 'opencensus/proto/agent/common/v1/common_pb'
require 'opencensus/proto/metrics/v1/metrics_pb'
require 'opencensus/proto/resource/v1/resource_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("opencensus/proto/agent/metrics/v1/metrics_service.proto", :syntax => :proto3) do
    add_message "opencensus.proto.agent.metrics.v1.ExportMetricsServiceRequest" do
      optional :node, :message, 1, "opencensus.proto.agent.common.v1.Node"
      repeated :metrics, :message, 2, "opencensus.proto.metrics.v1.Metric"
      optional :resource, :message, 3, "opencensus.proto.resource.v1.Resource"
    end
    add_message "opencensus.proto.agent.metrics.v1.ExportMetricsServiceResponse" do
    end
  end
end

module OpenCensus
  module Proto
    module Agent
      module Metrics
        module V1
          ExportMetricsServiceRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("opencensus.proto.agent.metrics.v1.ExportMetricsServiceRequest").msgclass
          ExportMetricsServiceResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("opencensus.proto.agent.metrics.v1.ExportMetricsServiceResponse").msgclass
        end
      end
    end
  end
end
