// Copyright 2021, OpenTelemetry Authors
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
// +build !windows

package smartagentextension

import (
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestBundleDirDefault(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.Nil(t, err)

	factory := NewFactory()
	factories.Extensions[typeStr] = factory
	cfg, err := configtest.LoadConfigFile(
		t, path.Join(".", "testdata", "config.yaml"), factories,
	)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	require.GreaterOrEqual(t, len(cfg.Extensions), 1)

	allSettingsConfig := cfg.Extensions["smartagent/default_settings"]
	require.NotNil(t, allSettingsConfig)

	saConfigProvider, ok := allSettingsConfig.(SmartAgentConfigProvider)
	require.True(t, ok)

	require.Equal(t, "/usr/lib/splunk-otel-collector/agent-bundle", saConfigProvider.SmartAgentConfig().BundleDir)
}
