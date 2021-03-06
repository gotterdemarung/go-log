package log

import (
	"io"
	"strings"
	"github.com/mgutz/ansi"
)

type colorFunc func(string) string
type colorFuncGen func(string) func(string) string

type AnsiPalette struct {
	BulletTrace colorFunc
	BulletDebug colorFunc
	BulletIn colorFunc
	BulletOut colorFunc
	BulletInfo colorFunc
	BulletSuccess colorFunc
	BulletWarn colorFunc
	BulletError colorFunc

	Time colorFunc
	TimePrecise colorFunc

	Tags colorFunc

	TypeNil colorFunc
	TypeBool colorFunc
	TypeString colorFunc
	TypeNumber colorFunc
}

func AnsiBuildPalette(colorFuncGen colorFuncGen) AnsiPalette {
	return AnsiPalette{
		BulletTrace: colorFuncGen("black+h"),
		BulletDebug: colorFuncGen("gray+h"),
		BulletIn: colorFuncGen("123"),
		BulletOut: colorFuncGen("123"),
		BulletInfo: colorFuncGen("25"),
		BulletSuccess: colorFuncGen("white:28"),
		BulletWarn: colorFuncGen("184"),
		BulletError: colorFuncGen("white+h:red"),

		Time: colorFuncGen("24"),
		TimePrecise: colorFuncGen("23"),

		Tags: colorFuncGen("60"),

		TypeNil: colorFuncGen("97"),
		TypeBool: colorFuncGen("103"),
		TypeString: colorFuncGen("28"),
		TypeNumber: colorFuncGen("25"),
	}
}

func nullColorGenFunc(string) func(string) string {
	return func(in string) string {
		return in
	}
}

// Common appender configuration
type AnsiAppenderOptions struct {
	Bullets bool
	Time bool
	Tags bool
	Precise bool
	Colors bool
}

func GetAnsiAppender(f io.Writer, opts AnsiAppenderOptions) Appender {
	var colorizer colorFuncGen;
	if opts.Colors {
		colorizer = ansi.ColorFunc
	} else {
		colorizer = nullColorGenFunc
	}

	pal := AnsiBuildPalette(colorizer);

	bullets := map[Type]string{}
	if opts.Bullets {
		bullets[TRACE] = " " + pal.BulletTrace("---") + " ";
		bullets[IN] = " " + pal.BulletIn("❮❮❮") + " ";
		bullets[OUT] = " " + pal.BulletOut("❯❯❯") + " ";
		bullets[DEBUG] = " " + pal.BulletDebug("DBG") + " ";
		bullets[INFO] = " " + pal.BulletInfo("INF") + " ";
		bullets[SUCCESS] = " " + pal.BulletSuccess(" ✓ ") + " ";
		bullets[WARN] = " " + pal.BulletWarn("WRN") + " ";
		bullets[ERROR] = " " + pal.BulletError("ERR") + " ";
	}

	return func(p *Packet) {
		str := ""
		bullet, ok := bullets[p.Level]
		if !ok && opts.Bullets {
			bullet = "     "
		}
		str += bullet
		if opts.Precise || opts.Time {
			str += pal.Time(p.SimpleTime())
			if opts.Precise {
				str += pal.TimePrecise("." + p.PreciseString())
			}
			str += " "
		}
		if opts.Tags {
			str += pal.Tags(strings.Join(p.Tags, " "))
			str += " "
		}

		if len(p.Values) > 0 {
			str += Substitute(p.Message + " ", func (full string, group string, separator string) string {
				if val, ok := p.Values[group]; ok {
					return SubstituteTypeHelper(
						val,
						func () string {return pal.TypeNil("null") + separator},
						func (in string) string {
							return pal.TypeString(in) + separator
						},
						func (in string) string {
							return pal.TypeNumber(in) + separator
						},
						func (in string) string {
							return pal.TypeBool(in) + separator
						},
					)
					return val.(string) + separator
				} else {
					return full
				}
			});
		} else {
			str += p.Message
		}

		f.Write([]byte(str + "\n"))
	}
}