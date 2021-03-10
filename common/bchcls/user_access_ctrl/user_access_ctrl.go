//go:binary-only-package

package user_access_ctrl

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_access_ctrl_i"
	"github.com/smannathan/common-sdk/common/bchcls/user_access_ctrl/user_access_manager"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)