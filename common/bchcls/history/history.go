//go:binary-only-package

package history

import (
	"common/bchcls/asset_mgmt/asset_manager"
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/history/history_manager"
	"common/bchcls/internal/history_i"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)