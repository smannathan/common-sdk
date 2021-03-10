//go:binary-only-package

package datatype_i

import (
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/data_model"
	"common/bchcls/datatype/datatype_interface"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/datatype_i/datatype_c"
	"common/bchcls/internal/key_mgmt_i/key_mgmt_c"
	"common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)