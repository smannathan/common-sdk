//go:binary-only-package

package user_mgmt

import (
	"common/bchcls/asset_mgmt/asset_manager"
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/user_mgmt_i"
	"common/bchcls/simple_rule"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)