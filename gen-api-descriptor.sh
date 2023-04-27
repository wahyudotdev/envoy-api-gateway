protoc -I ./proto \
   --descriptor_set_out=api_descriptor.pb \
   ./proto/bookstore/bookstore.proto \
   ./proto/cart/cart.proto \
   ./proto/google/api/annotations.proto \
   ./proto/google/api/http.proto \
   ./proto/google/protobuf/descriptor.proto
