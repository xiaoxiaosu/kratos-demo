package localize

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type localizeKey struct{}

func I18N() middleware.Middleware {
	// 设置默认语言规则
	bundle := i18n.NewBundle(language.English)
	// 注册规则解析器
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//加载语言包
	bundle.LoadMessageFile("../../language/language.en.toml")
	bundle.LoadMessageFile("../../language/language.zh.toml")

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				accpet := tr.RequestHeader().Get("accept-language")
				locaizer := i18n.NewLocalizer(bundle, accpet)
				ctx = context.WithValue(ctx, localizeKey{}, locaizer)
			}

			return handler(ctx, req)
		}
	}
}

func Translate(ctx context.Context, translateGroup string, translateWord string) (string, error) {
	localizer := ctx.Value(localizeKey{}).(*i18n.Localizer)

	return localizer.Localize(&i18n.LocalizeConfig{
		MessageID: fmt.Sprintf("%s.%s", translateGroup, translateWord),
	})
}
