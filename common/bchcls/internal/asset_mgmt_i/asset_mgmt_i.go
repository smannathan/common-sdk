//go:binary-only-package

package asset_mgmt_i

import (
	"common/bchcls/asset_mgmt/asset_key_func"
	"common/bchcls/asset_mgmt/asset_manager"
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/index"
	"common/bchcls/index/table_interface"
	"common/bchcls/internal/asset_mgmt_i/asset_mgmt_c"
	"common/bchcls/internal/asset_mgmt_i/asset_mgmt_c/asset_mgmt_g"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/consent_mgmt_i/consent_mgmt_c"
	"common/bchcls/internal/datastore_i/datastore_c"
	"common/bchcls/internal/datatype_i"
	"common/bchcls/internal/key_mgmt_i"
	"common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"common/bchcls/simple_rule"
	"common/bchcls/test_utils"
	"common/bchcls/utils"

	"bytes"
	"encoding/json"
	"fmt"

	"reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)
