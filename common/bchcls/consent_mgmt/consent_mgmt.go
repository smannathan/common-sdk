//go:binary-only-package

package consent_mgmt

import (
	"github.com/smannathan/common-sdk/common/bchcls/asset_mgmt/asset_key_func"
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/consent_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/simple_rule"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)