module golang.org/x/tools/gopls

go 1.18

require (
	github.com/google/go-cmp v0.6.0
	github.com/goplus/gop v1.1.4-0.20240113075819-d3f76cee389e
	github.com/goplus/mod v0.12.2-0.20240107203906-5044606d0c51
	github.com/jba/printsrc v0.2.2
	github.com/jba/templatecheck v0.7.0
	github.com/qiniu/x v1.13.2
	github.com/sergi/go-diff v1.1.0
	golang.org/x/mod v0.14.0
	golang.org/x/sync v0.6.0
	golang.org/x/sys v0.16.0
	golang.org/x/text v0.14.0
	golang.org/x/tools v0.17.0
	golang.org/x/vuln v0.0.0-20230110180137-6ad3e3d07815
	gopkg.in/yaml.v3 v3.0.1
	honnef.co/go/tools v0.4.6
	mvdan.cc/gofumpt v0.4.0
	mvdan.cc/xurls/v2 v2.5.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/google/safehtml v0.1.0 // indirect
	github.com/goplus/gox v1.13.2 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
	golang.org/x/exp/typeparams v0.0.0-20221212164502-fae10dda9338 // indirect
)

replace golang.org/x/tools => ../
