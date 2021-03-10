//go:binary-only-package

package history_i

import (
	"common/bchcls/asset_mgmt/asset_manager"
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/history/history_manager"
	"common/bchcls/index"
	"common/bchcls/internal/asset_mgmt_i"
	"common/bchcls/internal/common/global"
	"common/bchcls/simple_rule"
	"common/bchcls/utils"

	"encoding/json"
	"math/rand"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)