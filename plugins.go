package main

import (
	"alpine/plugins/docs"

	"github.com/infinitybotlist/sysmanage-web/plugins/actions"
	"github.com/infinitybotlist/sysmanage-web/plugins/frontend"
	"github.com/infinitybotlist/sysmanage-web/plugins/logger"
	"github.com/infinitybotlist/sysmanage-web/plugins/persist"
	"github.com/infinitybotlist/sysmanage-web/types"
)

var meta = types.ServerMeta{
	ConfigVersion: 1,
	Port:          10394,
	Plugins: []types.Plugin{
		// Persist has no frotend, it is a backend plugin
		{
			ID:   persist.ID,
			Init: persist.InitPlugin,
		},
		{
			ID:   actions.ID,
			Init: actions.InitPlugin,
			Frontend: types.Provider{
				Provider: "@core",
			},
		},
		// Frontend has no frontend, it is a backend plugin
		{
			ID:   frontend.ID,
			Init: frontend.InitPlugin,
		},
		// Example of a custom plugin
		{
			ID:   docs.ID,
			Init: docs.InitPlugin,
			Frontend: types.Provider{
				Provider: "frontend/extplugins/foo", // This is the path to the plugin's frontend
			},
		},
		{
			ID:   logger.ID,
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
	// Enable this for using a node server
	//
	/*FrontendServer: &types.FrontendServer{
		Host:                "http://localhost:10385",
		ExtraHeadersToAllow: []string{},
		Dir:                 "frontend",
		RunCommand:          "node build",
		ExtraEnv: []string{
			"NODE_ENV=production",
			"PORT=10385",
			"HOST=127.0.0.1",
		},
	},*/
}
