/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalExternalConfiguration writes a value of the 'external_configuration' type to the given writer.
func MarshalExternalConfiguration(object *ExternalConfiguration, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeExternalConfiguration(object, stream)
	stream.Flush()
	return stream.Error
}

// writeExternalConfiguration writes a value of the 'external_configuration' type to the given stream.
func writeExternalConfiguration(object *ExternalConfiguration, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.labels != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("labels")
		stream.WriteObjectStart()
		stream.WriteObjectField("items")
		writeLabelList(object.labels.items, stream)
		stream.WriteObjectEnd()
		count++
	}
	if object.syncsets != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("syncsets")
		stream.WriteObjectStart()
		stream.WriteObjectField("items")
		writeSyncsetList(object.syncsets.items, stream)
		stream.WriteObjectEnd()
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalExternalConfiguration reads a value of the 'external_configuration' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalExternalConfiguration(source interface{}) (object *ExternalConfiguration, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readExternalConfiguration(iterator)
	err = iterator.Error
	return
}

// readExternalConfiguration reads a value of the 'external_configuration' type from the given iterator.
func readExternalConfiguration(iterator *jsoniter.Iterator) *ExternalConfiguration {
	object := &ExternalConfiguration{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "labels":
			value := &LabelList{}
			for {
				field := iterator.ReadObject()
				if field == "" {
					break
				}
				switch field {
				case "kind":
					text := iterator.ReadString()
					value.link = text == LabelListLinkKind
				case "href":
					text := iterator.ReadString()
					value.href = &text
				case "items":
					value.items = readLabelList(iterator)
				default:
					iterator.ReadAny()
				}
			}
			object.labels = value
		case "syncsets":
			value := &SyncsetList{}
			for {
				field := iterator.ReadObject()
				if field == "" {
					break
				}
				switch field {
				case "kind":
					text := iterator.ReadString()
					value.link = text == SyncsetListLinkKind
				case "href":
					text := iterator.ReadString()
					value.href = &text
				case "items":
					value.items = readSyncsetList(iterator)
				default:
					iterator.ReadAny()
				}
			}
			object.syncsets = value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
