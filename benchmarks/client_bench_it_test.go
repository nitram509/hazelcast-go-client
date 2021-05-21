/*
 * Copyright (c) 2008-2021, Hazelcast, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package benchmarks_test

import (
	"testing"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/internal/it"
)

func BenchmarkCreateShutdownClient(b *testing.B) {
	it.Benchmarker(b, func(b *testing.B, cb *hazelcast.ConfigBuilder) {
		for i := 0; i < b.N; i++ {
			client := it.MustClient(hazelcast.StartNewClientWithConfig(cb))
			it.Must(client.Shutdown())
		}
	})
}
