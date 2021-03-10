//go:binary-only-package

package consent_mgmt_i

import (
	"github.com/smannathan/common-sdk/common/bchcls/asset_mgmt/asset_key_func"
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/custom_errors"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/datatype/datatype_interface"
	"github.com/smannathan/common-sdk/common/bchcls/index"
	"github.com/smannathan/common-sdk/common/bchcls/internal/asset_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/consent_mgmt_i/consent_mgmt_c"
	"github.com/smannathan/common-sdk/common/bchcls/internal/datatype_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/key_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"github.com/smannathan/common-sdk/common/bchcls/simple_rule"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"encoding/json"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)