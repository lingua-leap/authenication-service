#!/bin/bash
echo "Running script with argument: $1"
CURRENT_DIR=$1
<<<<<<< HEAD
rm -rf "${CURRENT_DIR}/genproto"
find "${CURRENT_DIR}/protocol-buffers" -type f -name "*.proto" -print0 | while IFS= read -r -d '' file; do
  protoc -I=${CURRENT_DIR}/protocol-buffers -I=/usr/local/go --go_out=${CURRENT_DIR} --go-grpc_out=${CURRENT_DIR} "${file}"
done
=======
rm -rf "${CURRENT_DIR}/generated"
find "${CURRENT_DIR}/protos" -type f -name "*.proto" -print0 | while IFS= read -r -d '' file; do
  protoc -I=${CURRENT_DIR}/protos -I=/usr/local/go --go_out=${CURRENT_DIR} --go-grpc_out=${CURRENT_DIR} "${file}"
done  

>>>>>>> cd1ed2630bbb1681a55fe39ec4167e7375d69115
