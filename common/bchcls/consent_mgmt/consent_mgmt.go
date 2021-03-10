//go:binary-only-package

package consent_mgmt

import (
	"common/bchcls/asset_mgmt/asset_key_func"
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/consent_mgmt_i"
	"common/bchcls/simple_rule"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)