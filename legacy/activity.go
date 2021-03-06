// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package legacy

import (
	"encoding/json"
)

// Activity represents a report activity
type Activity struct {
	ID                *uint64         `json:"id"`
	Editable          *bool           `json:"editable"`
	IsInternal        *bool           `json:"is_internal"`
	Message           *string         `json:"message"`
	MarkdownMessage   *string         `json:"markdown_message"`
	AutomatedResponse *bool           `json:"automated_response"`
	CreatedAt         *Timestamp      `json:"created_at"`
	UpdatedAt         *Timestamp      `json:"updated_at"`
	RawActor          json.RawMessage `json:"actor"`
	AssignedUser      *User           `json:"assigned_user"`
	GeniusExecutionID *string         `json:"genius_execution_id"` // TODO: This may be wrong type
	FileName          *string         `json:"file_name"`
	ExpiringURL       *string         `json:"expiring_url"`
	Type              *string         `json:"type"`
}
