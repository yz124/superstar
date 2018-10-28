package conf

import "time"

/*
前面是含义，后面是 go 的表示值,多种表示,逗号","分割
年　 06,2006
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
周几 Mon,Monday
时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
时区字母缩写 MST
您看出规律了么！哦是的，你发现了，这里面没有一个是重复的，所有的值表示都唯一对应一个时间部分。
并且涵盖了很多格式组合。
*/
// 时间格式化字符串
const SysTimeform string = "2006-01-02 15:04:05"
const SysTimeformShort string = "2006-01-02"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")
