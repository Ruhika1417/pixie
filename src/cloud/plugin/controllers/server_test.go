/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package controllers_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"px.dev/pixie/src/cloud/plugin/controllers"
	"px.dev/pixie/src/cloud/plugin/pluginpb"
	"px.dev/pixie/src/cloud/plugin/schema"
	"px.dev/pixie/src/shared/services/pgtest"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	err := testMain(m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Got error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func testMain(m *testing.M) error {
	viper.Set("jwt_signing_key", "key0")
	s := bindata.Resource(schema.AssetNames(), schema.Asset)
	testDB, teardown, err := pgtest.SetupTestDB(s)
	if err != nil {
		return fmt.Errorf("failed to start test database: %w", err)
	}

	defer teardown()
	db = testDB

	if c := m.Run(); c != 0 {
		return fmt.Errorf("some tests failed with code: %d", c)
	}
	return nil
}

func mustLoadTestData(db *sqlx.DB) {
	db.MustExec(`DELETE FROM data_retention_plugin_releases`)
	db.MustExec(`DELETE FROM plugin_releases`)

	insertRelease := `INSERT INTO plugin_releases(name, id, description, logo, version, data_retention_enabled) VALUES ($1, $2, $3, $4, $5, $6)`
	db.MustExec(insertRelease, "test_plugin", "test-plugin", "This is a test plugin", "logo1", "0.0.1", "true")
	db.MustExec(insertRelease, "test_plugin", "test-plugin", "This is a newer test plugin", "logo2", "0.0.2", "true")
	db.MustExec(insertRelease, "test_plugin", "test-plugin", "This is the newest test plugin", "logo3", "0.0.3", "true")
	db.MustExec(insertRelease, "another_plugin", "another-plugin", "This is another plugin", "anotherLogo", "0.0.1", "false")
	db.MustExec(insertRelease, "another_plugin", "another-plugin", "This is another new plugin", "anotherLogo2", "0.0.2", "false")

	insertRetentionRelease := `INSERT INTO data_retention_plugin_releases(plugin_id, version, configurations, preset_scripts, documentation_url, default_export_url, allow_custom_export_url) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	db.MustExec(insertRetentionRelease, "test-plugin", "0.0.1", controllers.Configurations(map[string]string{"license_key": "This is what we use to authenticate"}), controllers.PresetScripts([]*controllers.PresetScript{
		&controllers.PresetScript{
			Name:              "http data",
			Description:       "This is a script to get http data",
			DefaultFrequencyS: 10,
			Script:            "script",
		},
		&controllers.PresetScript{
			Name:              "http data 2",
			Description:       "This is a script to get http data 2",
			DefaultFrequencyS: 20,
			Script:            "script 2",
		},
	}), "http://test-doc-url", "http://test-export-url", true)
	db.MustExec(insertRetentionRelease, "test-plugin", "0.0.2", controllers.Configurations(map[string]string{"license_key2": "This is what we use to authenticate 2"}), controllers.PresetScripts([]*controllers.PresetScript{
		&controllers.PresetScript{
			Name:              "dns data",
			Description:       "This is a script to get dns data",
			DefaultFrequencyS: 10,
			Script:            "dns script",
		},
		&controllers.PresetScript{
			Name:              "dns data 2",
			Description:       "This is a script to get dns data 2",
			DefaultFrequencyS: 20,
			Script:            "dns script 2",
		},
	}), "http://test-doc-url2", "http://test-export-url2", true)
	db.MustExec(insertRetentionRelease, "test-plugin", "0.0.3", controllers.Configurations(map[string]string{"license_key3": "This is what we use to authenticate 3"}), nil, "http://test-doc-url3", "http://test-export-url3", true)
}

func TestServer_GetPlugins(t *testing.T) {
	mustLoadTestData(db)

	s := controllers.New(db, "test")
	resp, err := s.GetPlugins(context.Background(), &pluginpb.GetPluginsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 2, len(resp.Plugins))
	assert.ElementsMatch(t, []*pluginpb.Plugin{
		&pluginpb.Plugin{
			Name:             "test_plugin",
			ID:               "test-plugin",
			LatestVersion:    "0.0.3",
			RetentionEnabled: true,
			Description:      "This is the newest test plugin",
			Logo:             "logo3",
		},
		&pluginpb.Plugin{
			Name:             "another_plugin",
			ID:               "another-plugin",
			LatestVersion:    "0.0.2",
			RetentionEnabled: false,
			Description:      "This is another new plugin",
			Logo:             "anotherLogo2",
		},
	}, resp.Plugins)
}

func TestServer_GetPluginsWithKind(t *testing.T) {
	mustLoadTestData(db)

	s := controllers.New(db, "test")
	resp, err := s.GetPlugins(context.Background(), &pluginpb.GetPluginsRequest{Kind: pluginpb.PLUGIN_KIND_RETENTION})
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 1, len(resp.Plugins))
	assert.ElementsMatch(t, []*pluginpb.Plugin{
		&pluginpb.Plugin{
			Name:             "test_plugin",
			ID:               "test-plugin",
			LatestVersion:    "0.0.3",
			RetentionEnabled: true,
			Description:      "This is the newest test plugin",
			Logo:             "logo3",
		},
	}, resp.Plugins)
}

func TestServer_GetRetentionPluginConfig(t *testing.T) {
	mustLoadTestData(db)

	s := controllers.New(db, "test")
	resp, err := s.GetRetentionPluginConfig(context.Background(), &pluginpb.GetRetentionPluginConfigRequest{
		ID:      "test-plugin",
		Version: "0.0.2",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, &pluginpb.GetRetentionPluginConfigResponse{
		Configurations: map[string]string{
			"license_key2": "This is what we use to authenticate 2",
		},
		DocumentationURL:     "http://test-doc-url2",
		DefaultExportURL:     "http://test-export-url2",
		AllowCustomExportURL: true,
		PresetScripts: []*pluginpb.GetRetentionPluginConfigResponse_PresetScript{
			&pluginpb.GetRetentionPluginConfigResponse_PresetScript{
				Name:              "dns data",
				Description:       "This is a script to get dns data",
				DefaultFrequencyS: 10,
				Script:            "dns script",
			},
			&pluginpb.GetRetentionPluginConfigResponse_PresetScript{
				Name:              "dns data 2",
				Description:       "This is a script to get dns data 2",
				DefaultFrequencyS: 20,
				Script:            "dns script 2",
			},
		},
	}, resp)
}
