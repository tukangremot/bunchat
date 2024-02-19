package modules

import (
	"github.com/fazpass/goliath/v3/module"
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/channel"
	"github.com/tukangremot/bunchat/backend/src/modules/chat"
)

type Modules struct {
	Chat    module.ModuleInterface
	Channel module.ModuleInterface
}

func Init(app *app.App) *Modules {
	return &Modules{
		Chat:    chat.Init(app),
		Channel: channel.Init(app),
	}
}
