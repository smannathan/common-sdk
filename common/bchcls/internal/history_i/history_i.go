//go:binary-only-package

package history_i

import (
	"github.com/smannathan/common-sdk/common/bchcls/asset_mgmt/asset_manager"
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/custom_errors"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/history/history_manager"
	"github.com/smannathan/common-sdk/common/bchcls/index"
	"github.com/smannathan/common-sdk/common/bchcls/internal/asset_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/simple_rule"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"encoding/json"
	"math/rand"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)