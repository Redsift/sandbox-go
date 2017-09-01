#!/bin/bash
rm -rf ./sift ./sift.go

cat > sift.go << EOF
package sandbox

import (
	rpc "github.com/redsift/go-sandbox-rpc"
)

type RedsiftFunc func(rpc.ComputeRequest) ([]rpc.ComputeResponse, error)

var Computes = map[int]RedsiftFunc{}
EOF