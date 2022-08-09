module github.com/zhangdapeng520/zdpgo_sim

go 1.18

require (
	github.com/zhangdapeng520/zdpgo_clearcode v0.1.4
	github.com/zhangdapeng520/zdpgo_file v1.0.9
	github.com/zhangdapeng520/zdpgo_lexers v0.1.1
	github.com/zhangdapeng520/zdpgo_pool_goroutine v0.1.1
	github.com/zhangdapeng520/zdpgo_pygments v0.1.8
	github.com/zhangdapeng520/zdpgo_type v0.1.7
	golang.org/x/text v0.3.7
)

replace (
	github.com/zhangdapeng520/zdpgo_clearcode v0.1.4 => ../zdpgo_clearcode
	github.com/zhangdapeng520/zdpgo_pygments v0.1.8 => ../zdpgo_pygments
)
