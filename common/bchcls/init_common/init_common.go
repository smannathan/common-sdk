//go:binary-only-package

package init_common

import (
	"github.com/smannathan/common-sdk/common/bchcls/asset_mgmt"
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/consent_mgmt"
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/datastore"
	"github.com/smannathan/common-sdk/common/bchcls/datastore/datastore_manager"
	"github.com/smannathan/common-sdk/common/bchcls/datatype"
	"github.com/smannathan/common-sdk/common/bchcls/history"
	"github.com/smannathan/common-sdk/common/bchcls/index"
	"github.com/smannathan/common-sdk/common/bchcls/user_mgmt"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)