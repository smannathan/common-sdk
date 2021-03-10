//go:binary-only-package

package user_access_ctrl

import (
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/user_access_ctrl_i"
	"common/bchcls/user_access_ctrl/user_access_manager"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)