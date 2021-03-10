//go:binary-only-package

package user_groups

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/graph"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i/user_mgmt_c"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)