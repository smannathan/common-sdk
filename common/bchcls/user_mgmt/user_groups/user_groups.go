//go:binary-only-package

package user_groups

import (
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/common/graph"
	"common/bchcls/internal/user_mgmt_i"
	"common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"common/bchcls/utils"

	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)