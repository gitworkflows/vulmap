// keys package contains the public key for verifying digital signature of templates
package keys

import _ "embed"

//go:embed vulmap.crt
var VulmapCert []byte // public key for verifying digital signature of templates
