package main

import (
	"alpine/plugins/foo"

	"github.com/infinitybotlist/sysmanage-web/plugins/actions"
	"github.com/infinitybotlist/sysmanage-web/plugins/frontend"
	"github.com/infinitybotlist/sysmanage-web/plugins/logger"
	"github.com/infinitybotlist/sysmanage-web/plugins/persist"
	"github.com/infinitybotlist/sysmanage-web/types"
)

var meta = types.ServerMeta{
	ConfigVersion: 1,
	Port:          10394,
	Plugins: map[string]types.Plugin{
		// Persist has no frotend, it is a backend plugin
		"persist": {
			Init: persist.InitPlugin,
		},
		"actions": {
			Init: actions.InitPlugin,
			Frontend: types.Provider{
				Provider: "@core",
			},
		},
		// Frontend has no frontend, it is a backend plugin
		"frontend": {
			Init: frontend.InitPlugin,
		},
		"foo": {
			Init: foo.InitPlugin,
			Frontend: types.Provider{
				Provider: "frontend/extplugins/foo", // This is the path to the plugin's frontend
			},
		},
		"logger": {
			Init: logger.InitPlugin,
		},
	},
	Frontend: types.FrontendConfig{
		FrontendProvider: types.Provider{
			Provider: "frontend",
			Overrides: []string{
				"package.json",
				"package-lock.json",
				"yarn.lock",
				"tsconfig.json",
				"svelte.config.js",
				"postcss.config.js",
				"tailwind.config.js",
				"vite.config.js",
				"README.md",
				"src/app.d.ts",
			},
		},
		ComponentProvider: types.Provider{
			Provider: "@core",
		},
		CorelibProvider: types.Provider{
			Provider: "@core",
		}, // Use a custom corelib provider if you want to modify the corelib
		ExtraFiles: map[string]string{
			"frontend/src/lib/images":          "$lib/images",
			"frontend/src/lib/images/logo.png": "static/favicon.png",
		},
	},
	FrontendServer: &types.FrontendServer{
		Host:                "http://localhost:10385",
		ExtraHeadersToAllow: []string{},
		Dir:                 "frontend",
		RunCommand:          "node build",
		ExtraEnv: []string{
			"NODE_ENV=production",
			"PORT=10385",
			"HOST=127.0.0.1",
		},
	},
}
