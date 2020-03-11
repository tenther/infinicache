package server

import (
    "time"

    "github.com/mason-leap-lab/infinicache/proxy/lambdastore"
)

const LambdaMaxDeployments = 10
const NumLambdaClusters = 10
const LambdaStoreName = "LambdaStore" // replica version (no use)
const LambdaPrefix = "CacheNode"
const InstanceWarmTimout = 1 * time.Minute
const InstanceCapacity = 1536 * 1000000 // MB
const InstanceOverhead = 100 * 1000000  // MB

func init() {
    lambdastore.WarmTimout = InstanceWarmTimout
}
