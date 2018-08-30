module github.com/ahstn/atlas

require (
	9fans.net/go v0.0.0-20180727211846-5d4fa602e1e8 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.10 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/acroca/go-symbols v0.0.0-20180523203557-953befd75e22 // indirect
	github.com/apex/log v1.0.0
	github.com/briandowns/spinner v0.0.0-20180529140538-1567cd82701b
	github.com/cosiner/argv v0.0.0-20170225145430-13bacc38a0a5 // indirect
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/davidrjenni/reftools v0.0.0-20180509164333-3813a62570d2 // indirect
	github.com/derekparker/delve v1.1.0 // indirect
	github.com/docker/distribution v2.6.0-rc.1.0.20180327202408-83389a148052+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.3.0
	github.com/docker/go-units v0.3.3 // indirect
	github.com/fatih/color v1.7.0
	github.com/fatih/gomodifytags v0.0.0-20180826164257-7987f52a7108 // indirect
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/gogo/protobuf v1.1.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/google/go-cmp v0.2.0 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2 // indirect
	github.com/haya14busa/goplay v1.0.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/josharian/impl v0.0.0-20180228163738-3d0f908298c4 // indirect
	github.com/karrick/godirwalk v1.7.3 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.3 // indirect
	github.com/mdempsky/gocode v0.0.0-20180727200127-00e7f5ac290a // indirect
	github.com/onsi/ginkgo v1.6.0 // indirect
	github.com/onsi/gomega v1.4.1 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/peterh/liner v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ramya-rao-a/go-outline v0.0.0-20170803230019-9e9d089bb61a // indirect
	github.com/rogpeppe/godef v0.0.0-20170920080713-b692db1de522 // indirect
	github.com/sirupsen/logrus v1.0.6 // indirect
	github.com/skratchdot/open-golang v0.0.0-20160302144031-75fb7ed4208c // indirect
	github.com/spf13/cobra v0.0.3 // indirect
	github.com/spf13/pflag v1.0.2 // indirect
	github.com/sqs/goreturns v0.0.0-20180302073349-83e02874ec12 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.2.2
	github.com/urfave/cli v1.20.0
	github.com/uudashr/gopkgs v1.3.2 // indirect
	golang.org/x/arch v0.0.0-20180516175055-5de9028c2478 // indirect
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac // indirect
	golang.org/x/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	golang.org/x/net v0.0.0-20180719180050-a680a1efc54d
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	golang.org/x/sys v0.0.0-20180821140842-3b58ed4ad339 // indirect
	golang.org/x/text v0.3.0 // indirect
	golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2 // indirect
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52 // indirect
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 // indirect
	google.golang.org/grpc v1.14.0 // indirect
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gopkg.in/kyokomi/emoji.v1 v1.5.1
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.2.1
	gotest.tools v2.1.0+incompatible // indirect
)

//replace github.com/docker/docker v1.13.1 => github.com/docker/engine v0.0.0-20180816081446-320063a2ad06
// github.com/docker/engine v18.06.1-ce
replace github.com/docker/docker => github.com/docker/engine v0.0.0-20180816081446-320063a2ad06

// github.com/docker/distribution master
replace github.com/docker/distribution => github.com/docker/distribution v2.6.0-rc.1.0.20180820212402-02bf4a2887a4+incompatible
