//go:binary-only-package

package key_mgmt_i

import (
	"common/bchcls/cached_stub"
	"common/bchcls/data_model"
	"common/bchcls/internal/key_mgmt_i/key_mgmt_c"
	"common/bchcls/internal/key_mgmt_i/key_mgmt_c/key_mgmt_g"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)