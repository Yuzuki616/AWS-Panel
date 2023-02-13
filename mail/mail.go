package mail

import (
	"github.com/Yuzuki616/Aws-Panel/conf"
	"gopkg.in/gomail.v2"
	"strings"
)

const tmpl = "<div style=\"background: #eee\">\n    <table width=\"600\" border=\"0\" align=\"center\" cellpadding=\"0\" cellspacing=\"0\">\n        <tbody>\n        <tr>\n            <td>\n                <div style=\"background:#fff\">\n                    <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                        <thead>\n                        <tr>\n                            <td valign=\"middle\" style=\"padding-left:30px;background-color:#415A94;color:#fff;padding:20px 40px;font-size: 21px;\">{{$name}}</td>\n                        </tr>\n                        </thead>\n                        <tbody>\n                        <tr style=\"padding:40px 40px 0 40px;display:table-cell\">\n                            <td style=\"font-size:24px;line-height:1.5;color:#000;margin-top:40px\">邮箱验证码</td>\n                        </tr>\n                        <tr>\n                            <td style=\"font-size:14px;color:#333;padding:24px 40px 0 40px\">\n                                尊敬的用户您好！\n                                <br />\n                                <br />\n                                您的验证码是：{{$code}}，请在 5 分钟内进行验证。如果该验证码不为您本人申请，请无视。\n                            </td>\n                        </tr>\n                        <tr style=\"padding:40px;display:table-cell\">\n                        </tr>\n                        </tbody>\n                    </table>\n                </div>\n                <div>\n                    <table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n                        <tbody>\n                        <tr>\n                            <td style=\"padding:20px 40px;font-size:12px;color:#999;line-height:20px;background:#f7f7f7\"><a href=\"https://github.com/Yuzuki616/AWS-Panel\" style=\"font-size:14px;color:#929292\">AWS-Panel</a></td>\n                        </tr>\n                        </tbody>\n                    </table>\n                </div></td>\n        </tr>\n        </tbody>\n    </table>\n</div>"

var mail *gomail.Dialer

func Init() {
	if !conf.Config.EnableMailVerify {
		return
	}
	c := conf.Config.MailConfig
	mail = gomail.NewDialer(c.Host, c.Port, c.Email, c.Password)
}

func SendMail(to, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Config.MailConfig.Email)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "邮件验证码")
	m.SetBody("text/html", strings.ReplaceAll(tmpl,
		"{{$code}}", code))
	err := mail.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
