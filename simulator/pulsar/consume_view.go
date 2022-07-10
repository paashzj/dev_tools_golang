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
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var (
	pulsarUrlText       = canvas.NewText("pulsar url", color.Black)
	pulsarUrlEntry      = widget.NewEntry()
	pulsarTopicText     = canvas.NewText("pulsar topic", color.Black)
	pulsarTopicEntry    = widget.NewEntry()
	pulsarJwtTokenText  = canvas.NewText("pulsar jwt token", color.Black)
	pulsarJwtTokenEntry = widget.NewEntry()
	pulsarMsgText       = canvas.NewText("receive message", color.Black)
)

func ConsumeConfigView() *fyne.Container {
	pulsarUrlEntry.SetText("pulsar://localhost:6650")
	vBox := container.NewVBox()
	pulsarTlsInsecure := widget.NewCheck("pulsar tls insecure", func(b bool) {
	})
	pulsarTopicEntry.SetText("default")
	vBox.Add(pulsarUrlText)
	vBox.Add(pulsarUrlEntry)
	vBox.Add(pulsarTlsInsecure)
	vBox.Add(pulsarTopicText)
	vBox.Add(pulsarTopicEntry)
	vBox.Add(pulsarJwtTokenText)
	vBox.Add(pulsarJwtTokenEntry)
	return vBox
}

func ConsumeActionView() *fyne.Container {
	subscribeButton := widget.NewButton("subscribe", func() {
	})
	closeButton := widget.NewButton("close", func() {
	})
	receiveButton := widget.NewButton("receive", func() {
	})
	return container.NewHBox(subscribeButton, closeButton, receiveButton)
}

func ConsumeView() *fyne.Container {
	return container.NewVBox(ConsumeConfigView(), ConsumeActionView(), pulsarMsgText)
}
