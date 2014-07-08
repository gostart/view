package view

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/gostart/errs"
	"github.com/ungerik/go-dry"
)

var Config = Configuration{
	ListenAndServeAt: "0.0.0.0:80",
	BaseDirs:         []string{"."},
	StaticDirs:       []string{"static"},    // every StaticDir will be appended to every BaseDir to search for static files
	TemplateDirs:     []string{"templates"}, // every TemplateDir will be appended to every BaseDir to search for template files
	SessionTracker:   &CookieSessionTracker{},
	SessionDataStore: NewCookieSessionDataStore(),
}

var StructTagKey = "view"

type Configuration struct {
	initialized         bool
	ListenAndServeAt    string
	IsProductionServer  bool // IsProductionServer will be set to true if localhost resolves to one of ProductionServerIPs
	ProductionServerIPs []string
	DisableCachedViews  bool
	BaseDirs            []string
	StaticDirs          []string
	TemplateDirs        []string
	RedirectSubdomains  []string // Exapmle: "www"
	SiteName            string
	CookieSecret        string
	SessionTracker      SessionTracker
	SessionDataStore    SessionDataStore
	Debug               struct {
		ListenAndServeAt string
		Mode             bool // Will be set to true if IsProductionServer is false
		LogPaths         bool
		LogRedirects     bool
	}
}

func (config *Configuration) Init() error {
	if config.initialized {
		panic("view.Config already initialized")
	}

	if !config.IsProductionServer {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return err
		}
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				ip := ipNet.IP.String()
				for _, prodIP := range Config.ProductionServerIPs {
					if ip == prodIP {
						config.IsProductionServer = true
						break
					}
				}
			}
		}
	}

	if !config.IsProductionServer {
		config.Debug.Mode = true
	}

	// Check if dirs exists and make them absolute

	for i := range Config.BaseDirs {
		dir, err := filepath.Abs(os.ExpandEnv(Config.BaseDirs[i]))
		if err != nil {
			return err
		}
		if !dry.FileIsDir(dir) {
			return errs.Format("BaseDir does not exist: %s", dir)
		}
		Config.BaseDirs[i] = dir
		fmt.Println("BaseDir:", dir)
	}

	for i := range Config.StaticDirs {
		Config.StaticDirs[i] = os.ExpandEnv(Config.StaticDirs[i])
		fmt.Println("StaticDir:", Config.StaticDirs[i])
	}

	for i := range Config.TemplateDirs {
		Config.TemplateDirs[i] = os.ExpandEnv(Config.TemplateDirs[i])
		fmt.Println("TemplateDir:", Config.TemplateDirs[i])
	}

	config.initialized = true
	return nil
}

// Deploy copies all known resources (static and template files) to targetDir
func (config *Configuration) Deploy(targetDir string) error {
	panic("not implemented") // todo
}
