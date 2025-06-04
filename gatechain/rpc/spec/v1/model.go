// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

// Package v1 defines models exposed by algod rest api
package v1

import (
	"fmt"
	"strings"
	"time"

	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
)

// NodeStatus contains the information about a node status
// swagger:model NodeStatus
type NodeStatus struct {
	// LastRound indicates the last round seen
	//
	// required: true
	LastRound uint64 `json:"lastHeight"`

	// LastVersion indicates the last consensus version supported
	//
	// required: true
	LastVersion string `json:"lastConsensusVersion"`

	// NextVersion of consensus protocol to use
	//
	// required: true
	NextVersion string `json:"nextConsensusVersion"`

	// NextVersionRound is the round at which the next consensus version will apply
	//
	// required: true
	NextVersionRound uint64 `json:"nextConsensusVersionRound"`

	// NextVersionSupported indicates whether the next consensus version is supported by this node
	//
	// required: true
	NextVersionSupported bool `json:"nextConsensusVersionSupported"`

	// TimeSinceLastRound in nanoseconds
	//
	// required: true
	//TimeSinceLastRound float64 `json:"timeSinceLastRound"`
	TimeSinceLastRound uint64 `json:"timeSinceLastRound"`

	// CatchupTime in nanoseconds
	//
	// required: true
	CatchupTime int64 `json:"catchupTime"`

	// HasSyncedSinceStartup indicates whether a round has completed since startup
	// Required: true
	HasSyncedSinceStartup bool `json:"hasSyncedSinceStartup"`

	Step uint64 `json:"step"`

	Period uint64 `json:"period"`

	ZeroTimeStamp time.Time `json:"zeroTimeStamp"`

	//Deadline float64 `json:"deadline"`
	Deadline uint64 `json:"deadline"`

	//FastRecoveryDeadline float64 `json:"fastRecoveryDeadline"`
	FastRecoveryDeadline uint64 `json:"fastRecoveryDeadline"`

	MinimumTxFee uint64 `json:"minimumTxFee"`

	MinimumBlockTxFee uint64 `json:"minimumBlockTxFee"`
}

// Block contains a block information
// swagger:model Block
type Block struct {
	// Hash is the current block hash
	//
	// required: true
	Hash string `json:"hash"`

	// PreviousBlockHash is the previous block hash
	//
	// required: true
	PreviousBlockHash string `json:"previousBlockHash"`

	// Seed is the sortition seed
	//
	// required: true
	Seed string `json:"seed"`

	// Proposer is the address of this block proposer
	//
	// required: true
	Proposer string `json:"proposer"`

	// Round is the current round on which this block was appended to the chain
	//
	// required: true
	Round uint64 `json:"height"`

	// Period is the period on which the block was confirmed
	//
	// required: true
	Period uint64 `json:"period"`

	// TransactionsRoot authenticates the set of transactions appearing in the block.
	// More specifically, it's the root of a merkle tree whose leaves are the block's Txids, in lexicographic order.
	// For the empty block, it's 0.
	// Note that the TxnRoot does not authenticate the signatures on the transactions, only the transactions themselves.
	// Two blocks with the same transactions but in a different order and with different signatures will have the same TxnRoot.
	//
	// required: true
	TransactionsRoot string `json:"txnRoot"`

	// ProxyTransactions is the list of proxyTransactions in this block
	ProxyTransactions [][]byte `json:"txnps"`

	// TimeStamp in seconds since epoch
	//
	// required: true
	Timestamp int64 `json:"timestamp"`

	UpgradeState
	UpgradeVote

	CertCommitteeInfo  []CommitteeSingleInfo `json:"certCommitteeinfo"`
	BlockCommitteeInfo []CommitteeSingleInfo `json:"blockcommitteeinfo"`

	ConAccountState []string `json:"autoOfflineAccounts"`
}

type CommitteeSingleInfo struct {
	Address string `json:"singleCommitteeAddress"`
	Power   uint64 `json:"singleCommitteePower"`
}

// UpgradeState contains the information about a current state of an upgrade
// swagger:model UpgradeState
type UpgradeState struct {
	// CurrentProtocol is a string that represents the current protocol
	//
	// required: true
	CurrentProtocol string `json:"currentProtocol"`

	// NextProtocol is a string that represents the next proposed protocol
	//
	// required: true
	NextProtocol string `json:"nextProtocol"`

	// NextProtocolApprovals is the number of blocks which approved the protocol upgrade
	//
	// required: true
	NextProtocolApprovals uint64 `json:"nextProtocolApprovals"`

	// NextProtocolVoteBefore is the deadline round for this protocol upgrade (No votes will be consider after this round)
	//
	// required: true
	NextProtocolVoteBefore uint64 `json:"nextProtocolVoteBefore"`

	// NextProtocolSwitchOn is the round on which the protocol upgrade will take effect
	//
	// required: true
	NextProtocolSwitchOn uint64 `json:"nextProtocolSwitchOn"`
}

// UpgradeVote represents the vote of the block proposer with respect to protocol upgrades.
// swagger:model UpgradeVote
type UpgradeVote struct {
	// UpgradePropose indicates a proposed upgrade
	//
	// required: true
	UpgradePropose string `json:"upgradePropose"`

	// UpgradeApprove indicates a yes vote for the current proposal
	//
	// required: true
	UpgradeApprove bool `json:"upgradeApprove"`
}

type Response struct {
	Response appinterface.ResponseQuery `json:"response"`
}

// CheckTx result
type ResultBroadcastTx struct {
	Code uint32 `json:"code"`
	Data []byte `json:"data"`
	Log  string `json:"log"`
	Hash []byte `json:"hash"`
	TxId string
}
type ParticipationKeyResponse struct {
	Code             uint32                         `json:"code"`
	Log              string                         `json:"log"`
	Data             []byte                         `json:"data"`
	ParticipationKey appinterface.ParticipationData `json:"participationData"`
	FileName         string                         `json:"filename"`
}

func (pkRes ParticipationKeyResponse) String() string {
	var sb strings.Builder
	sb.WriteString("ConAccount:\n")
	sb.WriteString(fmt.Sprintf("  Code:	%d\n", pkRes.Code))
	sb.WriteString(fmt.Sprintf("  Log:	%s\n", pkRes.Log))
	if pkRes.ParticipationKey.Address != nil {
		sb.WriteString(fmt.Sprintf("  ParticipationKey:	%s\n", pkRes.ParticipationKey.String()))
	}
	if pkRes.FileName != "" {
		sb.WriteString(fmt.Sprintf("  FileName:	%s\n", pkRes.FileName))
	}
	return strings.TrimSpace(sb.String())
}

type ResultConAccount struct {
	Address []byte `json:"address"`
	Power   uint64 `json:"power"`
	Status  bool   `json:"status"`
}

func (rca ResultConAccount) String() string {
	return ""
}

func (rc ResultConAccount) IsFind() bool {
	return rc.Status
}
