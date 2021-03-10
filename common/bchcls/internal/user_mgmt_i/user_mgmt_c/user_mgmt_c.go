//go:binary-only-package

package user_mgmt_c

import (
	"github.com/pkg/errors"

	"common/bchcls/cached_stub"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/index"
	"common/bchcls/internal/asset_mgmt_i/asset_mgmt_c"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/common/graph"
	"common/bchcls/internal/key_mgmt_i/key_mgmt_c"
	"common/bchcls/internal/user_mgmt_i/user_mgmt_c/user_mgmt_g"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)