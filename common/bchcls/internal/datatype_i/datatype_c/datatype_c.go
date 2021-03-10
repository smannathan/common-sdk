//go:binary-only-package

package datatype_c

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/custom_errors"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/datatype/datatype_interface"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/graph"
	"github.com/smannathan/common-sdk/common/bchcls/utils"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)