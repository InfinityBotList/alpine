package docs

import (
	"github.com/infinitybotlist/sysmanage-web/plugins/frontend"
	"github.com/infinitybotlist/sysmanage-web/types"
)

const ID = "docs"

func InitPlugin(c *types.PluginConfig) error {
	frontend.AddLink(c, frontend.Link{
		Title:       "View Documentation",
		Description: "See the Infinity Development Documentation",
		LinkText:    "View Docs",
		Href:        "/docs",
	})

	return nil
}
