//go:binary-only-package

package init_common

import (
	"common/bchcls/asset_mgmt"
	"common/bchcls/cached_stub"
	"common/bchcls/consent_mgmt"
	"common/bchcls/crypto"
	"common/bchcls/data_model"
	"common/bchcls/datastore"
	"common/bchcls/datastore/datastore_manager"
	"common/bchcls/datatype"
	"common/bchcls/history"
	"common/bchcls/index"
	"common/bchcls/user_mgmt"
	"common/bchcls/utils"

	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)