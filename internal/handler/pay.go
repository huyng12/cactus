package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/huyng12/cactus/internal/wallet"
	"gopkg.in/tucnak/telebot.v2"
)

var payRegex = regexp.MustCompile(`^/mua ([0-9]+)(.*?)$`)

var responseMessage = `Số tiền thay đổi: %d
Số tiền trong tài khoản: %d
`

func Pay(bot *telebot.Bot, w wallet.Walleter) func(msg *telebot.Message) {
	return func(msg *telebot.Message) {
		if payRegex.MatchString(msg.Text) {
			matches := payRegex.FindStringSubmatch(msg.Text)

			amount, _ := strconv.Atoi(matches[1])
			purpose := strings.Trim(matches[2], " ")

			w.DoTx(wallet.MakeTx(-amount, purpose))

			_, _ = bot.Send(msg.Chat, fmt.Sprintf(responseMessage, amount, w.GetBalance()))
		}
	}
}
