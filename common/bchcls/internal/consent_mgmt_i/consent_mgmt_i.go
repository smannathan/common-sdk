//go:binary-only-package

package consent_mgmt_i

import (
	"common/bchcls/asset_mgmt/asset_key_func"
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/datatype/datatype_interface"
	"common/bchcls/index"
	"common/bchcls/internal/asset_mgmt_i"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/consent_mgmt_i/consent_mgmt_c"
	"common/bchcls/internal/datatype_i"
	"common/bchcls/internal/key_mgmt_i"
	"common/bchcls/internal/user_mgmt_i"
	"common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"common/bchcls/simple_rule"
	"common/bchcls/utils"

	"encoding/json"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)