#!/bin/sh
# you must always include the annotations.proto, descriptor.proto, and http.proto
# otherwise Envoy will always fail to read the descriptor
protoc -I ./proto \
   --descriptor_set_out=./envoy/api_descriptor.pb \
   ./proto/bookstore/bookstore.proto \
   ./proto/cart/cart.proto \
   ./proto/google/api/annotations.proto \
   ./proto/google/api/http.proto \
   ./proto/google/protobuf/descriptor.proto
