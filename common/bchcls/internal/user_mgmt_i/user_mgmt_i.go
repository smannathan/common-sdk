//go:binary-only-package

package user_mgmt_i

import (
	"github.com/smannathan/common-sdk/common/bchcls/asset_mgmt/asset_manager"
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/custom_errors"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/internal/asset_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/key_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i/user_mgmt_c/user_mgmt_g"
	"github.com/smannathan/common-sdk/common/bchcls/simple_rule"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)