package bar

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

type Bar struct {
	conn       *xgb.Conn
	rootWindow xproto.Window
}

func New() *Bar {
	conn, err := xgb.NewConn()
	if err != nil {
		panic(err)
	}
	return &Bar{
		conn:       conn,
		rootWindow: xproto.Setup(conn).DefaultScreen(conn).Root,
	}
}

func (b *Bar) Close() {
	b.conn.Close()
}

func (b *Bar) Update(data []byte) {
	xproto.ChangeProperty(
		b.conn,
		xproto.PropModeReplace,
		b.rootWindow,
		xproto.AtomWmName,
		xproto.AtomString,
		8,
		uint32(len(data)),
		data,
	)
}
