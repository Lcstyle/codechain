package hashchain

import (
	"os"

	"github.com/frankbraun/codechain/hashchain/internal/state"
	"github.com/frankbraun/codechain/util/lockfile"
)

// HashChain of threshold signatures over a chain of code changes.
type HashChain struct {
	lock  lockfile.Lock
	fp    *os.File
	chain []*link
	state *state.State
}

// Close the underlying file pointer of hash chain and release lock.
func (c *HashChain) Close() error {
	if c.fp == nil {
		return c.lock.Release()
	}
	err := c.fp.Close()
	if err != nil {
		c.lock.Release()
		return err
	}
	c.fp = nil
	return c.lock.Release()
}

// M returns the signature threshold.
func (c *HashChain) M() int {
	return c.state.M()
}

// N returns the total weight of all signers.
func (c *HashChain) N() int {
	return c.state.N()
}

// LastEntryHash returns the hash of the last entry.
func (c *HashChain) LastEntryHash() [32]byte {
	return c.chain[len(c.chain)-1].Hash()
}

// LastTreeHash returns the most current tree hash (can be unsigned).
func (c *HashChain) LastTreeHash() string {
	return c.state.LastTreeHash()
}

// LastSignedTreeHash returns the last signed tree hash.
func (c *HashChain) LastSignedTreeHash() (string, int) {
	return c.state.LastSignedTreeHash()
}

// TreeHashes returns a list of all tree hashes in order (starting from
// tree.EmptyHash).
func (c *HashChain) TreeHashes() []string {
	return c.state.TreeHashes()
}

// TreeComments returns a list of all tree comments in order (starting from
// tree.EmptyHash).
func (c *HashChain) TreeComments() []string {
	return c.state.TreeComments()
}

// Signer returns a map containing all active signers for hash chain.
func (c *HashChain) Signer() map[string]bool {
	return c.state.Signer()
}

// SignerComment returns the signer comment for given pubKey.
func (c *HashChain) SignerComment(pubKey string) string {
	return c.state.SignerComment(pubKey)
}

// SignerWeight returns the signer weight for given pubKey.
func (c *HashChain) SignerWeight(pubKey string) int {
	return c.state.SignerWeight(pubKey)
}

// SignerInfo returns signer pubKey and comment for patch with given treeHash.
func (c *HashChain) SignerInfo(treeHash string) (string, string) {
	link := c.chain[c.state.SourceLine(treeHash)]
	pubKey := link.typeFields[1]
	return pubKey, c.state.SignerComment(pubKey)
}

// LinkHash returns the link hash corresponding to given treeHash.
func (c *HashChain) LinkHash(treeHash string) [32]byte {
	return c.state.LinkHash(treeHash)
}

// EntryHash returns the entry hash for the given treeHash.
func (c *HashChain) EntryHash(treeHash [32]byte) [32]byte {
	var h [32]byte
	// TODO: implement
	return h
}

// UnsignedInfo returns a string slice with information about all unsigned
// entries suitable for printing.
// If TreeHash is defined it returns info until that treeHash.
// If omitSource is true source lines are omitted
func (c *HashChain) UnsignedInfo(treeHash string, omitSource bool) ([]string, error) {
	return c.state.UnsignedInfo(treeHash, omitSource)
}
