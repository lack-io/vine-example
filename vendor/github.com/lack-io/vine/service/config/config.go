// Copyright 2020 lack
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

package config

import (
	"github.com/lack-io/vine/service/config/reader"
	"github.com/lack-io/vine/service/config/source"
	"github.com/lack-io/vine/service/config/source/file"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// provide the reader.Values interface
	reader.Values
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
	// Stop the config loader/watcher
	Close() error
	// Load config sources
	Load(source ...source.Source) error
	// Force a source changeset sync
	Sync() error
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
}

// Watcher is the config watcher
type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

var (
	// Default Config Manager
	DefaultConfig, _ = NewConfig()
)

// NewConfig returns new config
func NewConfig(opts ...Option) (Config, error) {
	return newConfig(opts...)
}

// Return config as raw json
func Bytes() []byte {
	return DefaultConfig.Bytes()
}

// Return config as a map
func Map() map[string]interface{} {
	return DefaultConfig.Map()
}

// Scan values to a go type
func Scan(v interface{}) error {
	return DefaultConfig.Scan(v)
}

// Force a source changeset sync
func Sync() error {
	return DefaultConfig.Sync()
}

// Get a value from the config
func Get(path ...string) reader.Value {
	return DefaultConfig.Get(path...)
}

// Load config sources
func Load(source ...source.Source) error {
	return DefaultConfig.Load(source...)
}

// Watch a value for changes
func Watch(path ...string) (Watcher, error) {
	return DefaultConfig.Watch(path...)
}

// LoadFile is short hand for creating a file source and loading it
func LoadFile(path string) error {
	return Load(file.NewSource(
		file.WithPath(path),
	))
}
