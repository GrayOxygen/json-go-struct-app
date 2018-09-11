package main

import (
	"flag"

	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Constants
const htmlAbout = `</b>  
<div class="panel panel-default">
	<div class="panel-heading">
		<h3 class="panel-title">
        欢迎使用<b>JSON转非嵌套Struct工具!
		</h3>
	</div>
	<div class="panel-body">
		Powered By Json-To-Go.js Go-Astilectron <br> 
	</div>
</div>
<div class="panel panel-default">
	<div class="panel-heading">
		<h3 class="panel-title">
        关于我
		</h3>
	</div>
	<li class="list-group-item">
		<span class="badge">GitHub</span>
		https://github.com/GrayOxygen
	</li>
	<li class="list-group-item">
		<span class="badge">Blog</span>
		https://grayoxygen.github.io/ShineOxygenBlog
	</li>
 
</div>
`

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap-4.0.0
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug: *debug,
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astilectron.PtrStr("选项"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label: astilectron.PtrStr("关于我"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "about", htmlAbout, func(m *bootstrap.MessageIn) {
								// Unmarshal payload
								var s string
								if err := json.Unmarshal(m.Payload, &s); err != nil {
									astilog.Error(errors.Wrap(err, "unmarshaling payload failed"))
									return
								}
								astilog.Infof("About modal has been displayed and payload is %s!", s)
							}); err != nil {
								astilog.Error(errors.Wrap(err, "sending about event failed"))
							}
							return
						},
					},

					{Role: astilectron.MenuItemRoleClose},
				},
			},
			{
				Label: astilectron.PtrStr("编辑"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Role: astilectron.MenuItemRoleClose},
					{Role: astilectron.MenuItemRoleCopy},
					{Role: astilectron.MenuItemRoleCut},
					{Role: astilectron.MenuItemRoleDelete},
					{Role: astilectron.MenuItemRoleEditMenu},
					{Role: astilectron.MenuItemRoleForceReload},
					{Role: astilectron.MenuItemRoleMinimize},
					{Role: astilectron.MenuItemRolePaste},
					{Role: astilectron.MenuItemRolePasteAndMatchStyle},
					{Role: astilectron.MenuItemRoleQuit},
					{Role: astilectron.MenuItemRoleRedo},
					{Role: astilectron.MenuItemRoleReload},
					{Role: astilectron.MenuItemRoleResetZoom},
					{Role: astilectron.MenuItemRoleSelectAll},
					{Role: astilectron.MenuItemRoleToggleDevTools},
					{Role: astilectron.MenuItemRoleToggleFullScreen},
					{Role: astilectron.MenuItemRoleUndo},
					{Role: astilectron.MenuItemRoleWindowMenu},
					{Role: astilectron.MenuItemRoleZoomOut},
					{Role: astilectron.MenuItemRoleZoomIn},
				},
			},
		},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			//go func() {
			//	time.Sleep(5 * time.Second)
			//	if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
			//		astilog.Error(errors.Wrap(err, "sending check.out.menu event failed"))
			//	}
			//}()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center: astilectron.PtrBool(true),
				//Height:          astilectron.PtrInt(1080),
				//Width:           astilectron.PtrInt(1920),
				//Fullscreen: astilectron.PtrBool(true),
				Fullscreenable: astilectron.PtrBool(true),
				Resizable:      astilectron.PtrBool(true),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap-4.0.0 failed"))
	}
}
