package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

//计划任务函数
func crontab(self *openwechat.Self, friends openwechat.Friends, groups openwechat.Groups) {
	//每天上午10点00分00秒
	spec_noon := "00 30 12 * * ?"
	spec_off := "00 00 18 * * ?"
	spec_list := []string{"00 00 10 * * ?", "00 00 11 * * ?", "00 00 14 * * ?", "00 00 15 * * ?", "00 00 16 * * ?", "00 00 17 * * ?"}
	c := cron.New(cron.WithSeconds())
	for _, v := range spec_list {
		c.AddFunc(v, func() {
			qrImg, err := os.Open("test.jpg")
			if err != nil {
				log.Println(err)
			}
			defer qrImg.Close()
			// self.SendImageToGroup(groups.GetByNickName("污都⑦狼live long群"), qrImg)
			// self.SendTextToGroup(groups.GetByNickName("污都⑦狼live long群"), "@所有人 小殷机器人温馨提示：\n快起来走动走动，扭扭腰，活动脖子")
			self.SendImageToFriend(friends.GetByRemarkName("思羽"), qrImg)
			self.SendTextToFriend(friends.GetByRemarkName("思羽"), "小殷机器人温馨提示：\n快起来走动走动，扭扭腰，活动脖子")
		})
	}
	c.AddFunc(spec_off, func() {
		qrImg, err := os.Open("xiaban.jpg")
		if err != nil {
			log.Println(err)
		}
		defer qrImg.Close()
		// self.SendImageToGroup(groups.GetByNickName("污都⑦狼live long群"), qrImg)
		// self.SendTextToGroup(groups.GetByNickName("污都⑦狼live long群"), "@所有人 小殷机器人温馨提示：\n下班时间到了，请准时下班，好好吃饭")
		self.SendImageToFriend(friends.GetByRemarkName("思羽"), qrImg)
		self.SendTextToFriend(friends.GetByRemarkName("思羽"), "小殷机器人温馨提示：\n下班时间到了，请准时下班，好好吃饭")
	})
	c.AddFunc(spec_noon, func() {
		qrImg, err := os.Open("mid.jpg")
		if err != nil {
			log.Println(err)
		}
		defer qrImg.Close()
		self.SendImageToFriend(friends.GetByRemarkName("思羽"), qrImg)
		self.SendTextToFriend(friends.GetByRemarkName("思羽"), "小殷机器人温馨提示：\n午睡时间快到了，请准备睡觉，好梦呦")
	})
	c.Start()
	select {}
}

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	fmt.Println(groups, err)

	go crontab(self, friends, groups)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
