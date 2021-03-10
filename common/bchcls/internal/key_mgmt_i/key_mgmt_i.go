//go:binary-only-package

package key_mgmt_i

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/internal/key_mgmt_i/key_mgmt_c"
	"github.com/smannathan/common-sdk/common/bchcls/internal/key_mgmt_i/key_mgmt_c/key_mgmt_g"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)