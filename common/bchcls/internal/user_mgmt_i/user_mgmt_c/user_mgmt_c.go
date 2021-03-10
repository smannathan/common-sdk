//go:binary-only-package

package user_mgmt_c

import (
	"github.com/pkg/errors"

	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/custom_errors"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/index"
	"github.com/smannathan/common-sdk/common/bchcls/internal/asset_mgmt_i/asset_mgmt_c"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/graph"
	"github.com/smannathan/common-sdk/common/bchcls/internal/key_mgmt_i/key_mgmt_c"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i/user_mgmt_c/user_mgmt_g"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)