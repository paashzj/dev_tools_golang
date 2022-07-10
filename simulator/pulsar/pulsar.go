// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pulsar

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
	"time"
)

func buildPulsarClient(url string, authType string, jwtToken string) (pulsar.Client, error) {
	if authType == AuthTypeJwt {
		return pulsar.NewClient(pulsar.ClientOptions{
			URL:            url,
			Authentication: pulsar.NewAuthenticationToken(jwtToken),
		})
	} else {
		return pulsar.NewClient(pulsar.ClientOptions{
			URL: url,
		})
	}
}

func buildPulsarConsumer(client pulsar.Client, topic string) (pulsar.Consumer, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return client.Subscribe(pulsar.ConsumerOptions{
		Type:             pulsar.Failover,
		SubscriptionName: id.String(),
		Topic:            topic,
	})
}

func receiveMsg(consumer pulsar.Consumer) (pulsar.Message, error) {
	ctx, cancelFunc := context.WithTimeout(context.TODO(), 10*time.Millisecond)
	defer cancelFunc()
	return consumer.Receive(ctx)
}
