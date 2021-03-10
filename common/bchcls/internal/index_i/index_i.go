//go:binary-only-package

package index_i

import (
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/index/table_interface"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/common/rb_tree"
	"common/bchcls/internal/index_i/cloudant_index"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"

	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)