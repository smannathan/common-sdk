//go:binary-only-package

package asset_mgmt

import (
	"common/bchcls/asset_mgmt/asset_manager"
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/internal/asset_mgmt_i"
	"common/bchcls/internal/asset_mgmt_i/asset_mgmt_c"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)