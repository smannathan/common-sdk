//go:binary-only-package

package asset_mgmt_c

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/custom_errors"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/internal/asset_mgmt_i/asset_mgmt_c/asset_mgmt_g"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)