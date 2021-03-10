//go:binary-only-package

package user_access_ctrl_i

import (
	"encoding/json"

	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/internal/asset_mgmt_i"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/key_mgmt_i"
	"common/bchcls/user_access_ctrl/user_access_manager"
	"common/bchcls/user_mgmt"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)