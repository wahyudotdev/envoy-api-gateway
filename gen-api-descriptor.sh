#!/bin/sh
# You must always include all proto files that included inside service definition
# otherwise Envoy will always fail to read the descriptor
# In this case, annotations.proto, http.proto, and field_behavior.proto must be included
protoc -I ./proto \
   --descriptor_set_out=./envoy/api_descriptor.pb \
   ./proto/bookstore/bookstore.proto \
   ./proto/cart/cart.proto \
   ./proto/google/api/annotations.proto \
   ./proto/google/api/http.proto \
   ./proto/google/protobuf/descriptor.proto \
   ./proto/google/api/field_behavior.proto