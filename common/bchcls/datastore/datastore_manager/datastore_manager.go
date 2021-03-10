//go:binary-only-package

package datastore_manager

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/datastore"
	"github.com/smannathan/common-sdk/common/bchcls/internal/datastore_i"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)