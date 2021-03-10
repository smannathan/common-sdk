//go:binary-only-package

package datastore_i

import (
	"common/bchcls/cached_stub"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/datastore"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/datastore_i/datastore_c"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)