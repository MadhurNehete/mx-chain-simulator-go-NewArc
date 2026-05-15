package multisig

import crypto "github.com/multiversx/mx-chain-crypto-go"

var _ crypto.MultiSigner = (*DisabledMultiSig)(nil)
var _ crypto.MultiSignerV2 = (*DisabledMultiSig)(nil)

const signature = "signature"

// DisabledMultiSig represents a disabled multisigner implementation
type DisabledMultiSig struct {
}

// CreateSignatureShare returns the default signature as this is a disabled component
func (dms *DisabledMultiSig) CreateSignatureShare(_ []byte, _ []byte) ([]byte, error) {
	return []byte(signature), nil
}

// CreateSignatureShareV2 returns the default signature as this is a disabled component
func (dms *DisabledMultiSig) CreateSignatureShareV2(_ crypto.PrivateKey, _ []byte) ([]byte, error) {
	return []byte(signature), nil
}

// VerifySignatureShare returns nil as this is a disabled component
func (dms *DisabledMultiSig) VerifySignatureShare(_ []byte, _ []byte, _ []byte) error {
	return nil
}

// VerifySignatureShareV2 returns nil as this is a disabled component
func (dms *DisabledMultiSig) VerifySignatureShareV2(_ crypto.PublicKey, _ []byte, _ []byte) error {
	return nil
}

// AggregateSigs returns a dummy signature, as this is a disabled component
func (dms *DisabledMultiSig) AggregateSigs(_ [][]byte, _ [][]byte) ([]byte, error) {
	return []byte(signature), nil
}

// AggregateSigsV2 returns a dummy signature, as this is a disabled component
func (dms *DisabledMultiSig) AggregateSigsV2(_ []crypto.PublicKey, _ [][]byte) ([]byte, error) {
	return []byte(signature), nil
}

// VerifyAggregatedSig returns nil as this is a disabled component
func (dms *DisabledMultiSig) VerifyAggregatedSig(_ [][]byte, _ []byte, _ []byte) error {
	return nil
}

// VerifyAggregatedSigV2 returns nil as this is a disabled component
func (dms *DisabledMultiSig) VerifyAggregatedSigV2(_ []crypto.PublicKey, _ []byte, _ []byte) error {
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (dms *DisabledMultiSig) IsInterfaceNil() bool {
	return dms == nil
}
