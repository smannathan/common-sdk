//go:binary-only-package

package datastore_manager

import (
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/datastore"
	"common/bchcls/internal/datastore_i"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)