//go:binary-only-package

package index_i

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/index/table_interface"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/rb_tree"
	"github.com/smannathan/common-sdk/common/bchcls/internal/index_i/cloudant_index"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"

	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)