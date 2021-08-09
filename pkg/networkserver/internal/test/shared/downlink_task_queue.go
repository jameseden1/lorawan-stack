// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"context"
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

// handleDownlinkTaskQueueTest runs a test suite on q.
func handleDownlinkTaskQueueTest(ctx context.Context, q DownlinkTaskQueue) {
	t, a := test.MustNewTFromContext(ctx)

	pbs := [...]ttnpb.EndDeviceIdentifiers{
		{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
				ApplicationId: "test-app",
			},
			DeviceId: "test-dev",
		},
		{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
				ApplicationId: "test-app2",
			},
			DeviceId: "test-dev",
		},
		{
			ApplicationIdentifiers: ttnpb.ApplicationIdentifiers{
				ApplicationId: "test-app2",
			},
			DeviceId: "test-dev2",
		},
	}

	type slot struct {
		ctx    context.Context
		id     ttnpb.EndDeviceIdentifiers
		t      time.Time
		respCh chan<- TaskPopFuncResponse
	}

	if q.Shards() < 1 {
		t.Fatal("Expected at least one shard!")
	}
	popCtx := context.WithValue(ctx, &struct{}{}, "pop")
	slotCh := make(chan slot)
	errCh := make(chan error)
	nextPop := make([]chan struct{}, q.Shards())
	for s := 0; s < q.Shards(); s++ {
		nextPop[s] = make(chan struct{}, 10)
	}
	for s := 0; s < q.Shards(); s++ {
		go func(s int) {
			for {
				<-nextPop[s]
				select {
				case <-ctx.Done():
					return
				case errCh <- q.Pop(popCtx, s, func(ctx context.Context, id ttnpb.EndDeviceIdentifiers, t time.Time) (time.Time, error) {
					respCh := make(chan TaskPopFuncResponse)
					slotCh <- slot{
						ctx:    ctx,
						id:     id,
						t:      t,
						respCh: respCh,
					}
					resp := <-respCh
					return resp.Time, resp.Error
				}):
				}
			}
		}(s)
	}

	// Ensure the goroutines have started
	for s := 0; s < q.Shards(); s++ {
		nextPop[s] <- struct{}{}
	}

	// Ensure Pop is blocking on empty queue.
	select {
	case s := <-slotCh:
		t.Fatalf("Pop called f on empty schedule, slot: %+v", s)

	case err := <-errCh:
		t.Fatalf("Pop returned on empty schedule with error: %s", test.FormatError(err))

	case <-time.After(test.Delay):
	}

	err := q.Add(ctx, pbs[0], time.Unix(0, 0), false)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	select {
	case s := <-slotCh:
		a.So(s.id, should.Resemble, pbs[0])
		a.So(s.t, should.Equal, time.Unix(0, 0))
		if !a.So(s.ctx, should.HaveParentContextOrEqual, popCtx) {
			t.Fatal(s.ctx)
		}
		s.respCh <- TaskPopFuncResponse{}

	case err := <-errCh:
		a.So(err, should.BeNil)
		t.Fatal("Pop returned without calling f on non-empty schedule")

	case <-ctx.Done():
		t.Fatal("Timed out waiting for Pop to call f")
	}

	select {
	case err := <-errCh:
		if !a.So(err, should.BeNil) {
			t.FailNow()
		}

	case <-ctx.Done():
		t.Fatal("Timed out waiting for Pop to return")
	}

	err = q.Add(ctx, pbs[0], time.Unix(0, 42), false)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	err = q.Add(ctx, pbs[1], time.Now().Add(time.Hour), false)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	err = q.Add(ctx, pbs[1], time.Now().Add(2*time.Hour), true)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	err = q.Add(ctx, pbs[1], time.Unix(13, 0), false)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	err = q.Add(ctx, pbs[1], time.Unix(42, 0), true)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	err = q.Add(ctx, pbs[2], time.Now().Add(42*time.Hour), true)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	err = q.Add(ctx, pbs[2], time.Unix(42, 42), true)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}

	receivedSlots := make(map[ttnpb.EndDeviceIdentifiers][]time.Time)
	for i := 0; i < 3; i++ {
		for w := 0; w < q.Shards(); w++ {
			nextPop[w] <- struct{}{}
		}

		t.Helper()

		a := assertions.New(t)

		select {
		case s := <-slotCh:
			receivedSlots[s.id] = append(receivedSlots[s.id], s.t)
			a.So(s.ctx, should.HaveParentContextOrEqual, popCtx)
			s.respCh <- TaskPopFuncResponse{}

		case err := <-errCh:
			a.So(err, should.BeNil)
			t.Fatal("Pop returned without calling f on non-empty schedule")

		case <-ctx.Done():
			t.Fatal("Timed out waiting for Pop to call f")
		}

		select {
		case err := <-errCh:
			if !a.So(err, should.BeNil) {
				t.FailNow()
			}

		case <-ctx.Done():
			t.Fatal("Timed out waiting for Pop to return")
		}
	}

	a.So(receivedSlots, should.Resemble, map[ttnpb.EndDeviceIdentifiers][]time.Time{
		pbs[0]: {time.Unix(0, 42).UTC()},
		pbs[1]: {time.Unix(42, 0).UTC()},
		pbs[2]: {time.Unix(42, 42).UTC()},
	})
}

// HandleDownlinkTaskQueueTest runs a DownlinkTaskQueue test suite on reg.
func HandleDownlinkTaskQueueTest(t *testing.T, q DownlinkTaskQueue) {
	t.Helper()
	test.RunTest(t, test.TestConfig{
		Parallel: true,
		Func: func(ctx context.Context, a *assertions.Assertion) {
			t.Helper()
			test.RunSubtestFromContext(ctx, test.SubtestConfig{
				Name: "1st run",
				Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
					handleDownlinkTaskQueueTest(ctx, q)
				},
			})
			if t.Failed() {
				t.Skip("Skipping 2nd run")
			}
			test.RunSubtestFromContext(ctx, test.SubtestConfig{
				Name: "2st run",
				Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
					handleDownlinkTaskQueueTest(ctx, q)
				},
			})
		},
	})
}
