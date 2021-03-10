//go:binary-only-package

package asset_mgmt_c

import (
	"common/bchcls/cached_stub"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/internal/asset_mgmt_i/asset_mgmt_c/asset_mgmt_g"
	"common/bchcls/internal/common/global"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)