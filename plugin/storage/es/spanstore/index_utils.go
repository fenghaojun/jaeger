// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanstore

import (
	"time"

	"github.com/jaegertracing/jaeger/model"
)

// returns index name with date
func indexWithDate(indexPrefix string, date time.Time) string {
	spanDate := date.UTC().Format("060102")
	return indexPrefix + spanDate + "*"
}

// returns archive index name
func archiveIndex(indexPrefix, archiveSuffix string) string {
	return indexPrefix + archiveSuffix
}

// compare two span, return true if they are the same
func compareSpans(span1, span2 *model.Span) bool {
	return span1.TraceID == span2.TraceID && span1.SpanID == span2.SpanID && span1.StartTime.Equal(span2.StartTime) && span1.Duration == span2.Duration
}

func containsSpan(spans []*model.Span, span *model.Span) bool {
	for _, tmpSpan := range spans {
		if compareSpans(tmpSpan, span) {
			return true
		}
	}
	return false
}
