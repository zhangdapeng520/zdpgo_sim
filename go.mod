module github.com/zhangdapeng520/zdpgo_sim

go 1.18

require (
	github.com/zhangdapeng520/zdpgo_clearcode v0.1.0
	github.com/zhangdapeng520/zdpgo_pygments v0.1.0
	golang.org/x/text v0.3.7
)

replace (
	github.com/zhangdapeng520/zdpgo_pygments v0.1.0 => ../zdpgo_pygments
)