/*
   Copyright 2020 LittleBear(1018589158@qq.com)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package simulator

import (
	"container/list"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
	"reflect"
	"testing"
)

func NewNodeTest(nodeID int, numConnection int,
	region int, miningPower int64,
) *Node {

	node := &Node{
		nodeID:      nodeID,
		region:      region,
		miningPower: miningPower,
		//routingTable:      routingTable,
		//consensusAlgo:     consensusAlgo,
		block:             nil,
		orphans:           hashset.New(),
		mintingTask:       nil,
		sendingBlock:      false,
		messageQue:        new(list.List),
		downloadingBlocks: hashset.New(),
		processingTime:    2,
	}

	// Using the reflect function to find the Initial TABLE and ALGO
	node.SetRoutingTable(NewBitcoinCoreTable(node))
	node.SetConsensusAlgo(NewAbstractConsensusAlgo(node))
	node.setNumConnection(numConnection)

	return node
}

func TestBitcoinCoreTable_GetNeighbors(t *testing.T) {
	selfnode := NewNodeTest(0, 1, 1, 1)
	innode1 := NewNodeTest(1, 1, 1, 1)
	innode2 := NewNodeTest(2, 1, 1, 1)
	outnode1 := NewNodeTest(3, 1, 1, 1)
	outnode2 := NewNodeTest(4, 1, 1, 1)

	outbound := arraylist.New()
	outbound.Add(outnode1)
	outbound.Add(outnode2)

	inbound := arraylist.New()
	inbound.Add(innode1)
	inbound.Add(innode2)

	type fields struct {
		AbstractRoutingTable *AbstractRoutingTable
		outbound             *arraylist.List
		inbound              *arraylist.List
	}
	tests := []struct {
		name   string
		fields fields
		want   *arraylist.List
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				NewAbstractRoutingTable(selfnode),
				outbound,
				inbound,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bct := &BitcoinCoreTable{
				AbstractRoutingTable: tt.fields.AbstractRoutingTable,
				outbound:             tt.fields.outbound,
				inbound:              tt.fields.inbound,
			}
			if got := bct.GetNeighbors(); !reflect.DeepEqual(got, tt.want) {
				t.Logf("GetNeighbors() = %v, want %v", got, tt.want)
				if v, ok := bct.GetNeighbors().Get(0); ok {
					t.Log(v.(*Node).nodeID)
				}
				if v, ok := bct.GetNeighbors().Get(1); ok {
					t.Log(v.(*Node).nodeID)
				}
				if v, ok := bct.GetNeighbors().Get(2); ok {
					t.Log(v.(*Node).nodeID)
				}
				if v, ok := bct.GetNeighbors().Get(3); ok {
					t.Log(v.(*Node).nodeID)
				}
			}
		})
	}
}

func TestBitcoinCoreTable_InitTable(t *testing.T) {
	selfnode := NewNodeTest(1, 1, 1, 1)

	type fields struct {
		AbstractRoutingTable *AbstractRoutingTable
		outbound             *arraylist.List
		inbound              *arraylist.List
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				AbstractRoutingTable: NewAbstractRoutingTable(selfnode),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bct := &BitcoinCoreTable{
				AbstractRoutingTable: tt.fields.AbstractRoutingTable,
				outbound:             tt.fields.outbound,
				inbound:              tt.fields.inbound,
			}
			bct.InitTable()
		})
	}
}
