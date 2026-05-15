package p2p

import "strings"

// AddressContainsPeerID reports whether address identifies exactly peerID.
// It accepts either a bare peer ID or a multiaddress containing a /p2p/<peerID>
// or legacy /ipfs/<peerID> segment. Substring matches are intentionally rejected.
func AddressContainsPeerID(address string, peerID string) bool {
	if address == peerID {
		return true
	}

	parts := strings.Split(address, "/")
	for i := 0; i < len(parts)-1; i++ {
		switch parts[i] {
		case "p2p", "ipfs":
			if parts[i+1] == peerID {
				return true
			}
		}
	}

	return false
}
