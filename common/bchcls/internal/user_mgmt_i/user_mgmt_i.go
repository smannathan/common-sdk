//go:binary-only-package

package user_mgmt_i

import (
	"common/bchcls/asset_mgmt/asset_manager"
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/internal/asset_mgmt_i"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/key_mgmt_i"
	"common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"common/bchcls/internal/user_mgmt_i/user_mgmt_c/user_mgmt_g"
	"common/bchcls/simple_rule"
	"common/bchcls/utils"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)