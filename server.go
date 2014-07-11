package view

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/gostart/errs"
	"github.com/ungerik/go-dry"
)

var Port80 = Server{
	Address:      "0.0.0.0:80",
	BaseDirs:     []string{"."},
	StaticDirs:   []string{"static"},    // every StaticDir will be appended to every BaseDir to search for static files
	TemplateDirs: []string{"templates"}, // every TemplateDir will be appended to every BaseDir to search for template files
}

var Port8080 = Server{
	Address:      "0.0.0.0:8080",
	BaseDirs:     []string{"."},
	StaticDirs:   []string{"static"},    // every StaticDir will be appended to every BaseDir to search for static files
	TemplateDirs: []string{"templates"}, // every TemplateDir will be appended to every BaseDir to search for template files
}

type Server struct {
	initialized         bool
	Address             string
	ReadTimeoutMs       int
	WriteTimeoutMs      int
	NoKeepAlives        bool
	TLSCertFile         string
	TLSKeyFile          string
	IsProductionServer  bool // IsProductionServer will be set to true if localhost resolves to one of ProductionServerIPs
	ProductionServerIPs []string
	DisableCachedViews  bool
	BaseDirs            []string
	StaticDirs          []string
	TemplateDirs        []string
	RedirectSubdomains  []string // Exapmle: "www"
	SiteName            string
	CookieSecret        []byte // nil to disable cookie-encryption, or AES key 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256
	TrackSessions       bool
	SessionTracker      SessionTracker
	SessionDataStore    SessionDataStore

	Debug struct {
		ListenAndServeAt string
		Mode             bool // Will be set to true if IsProductionServer is false
		LogPaths         bool
		LogRedirects     bool
	}
}

func (server *Server) Init() error {
	if server.initialized {
		panic("view.Server already initialized")
	}

	if !server.IsProductionServer {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return err
		}
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				ip := ipNet.IP.String()
				for _, prodIP := range server.ProductionServerIPs {
					if ip == prodIP {
						server.IsProductionServer = true
						break
					}
				}
			}
		}
	}

	if !server.IsProductionServer {
		server.Debug.Mode = true
	}

	// Check if dirs exists and make them absolute

	for i := range server.BaseDirs {
		dir, err := filepath.Abs(os.ExpandEnv(server.BaseDirs[i]))
		if err != nil {
			return err
		}
		if !dry.FileIsDir(dir) {
			return errs.Format("BaseDir does not exist: %s", dir)
		}
		server.BaseDirs[i] = dir
		fmt.Println("BaseDir:", dir)
	}

	for i := range server.StaticDirs {
		server.StaticDirs[i] = os.ExpandEnv(server.StaticDirs[i])
		fmt.Println("StaticDir:", server.StaticDirs[i])
	}

	for i := range server.TemplateDirs {
		server.TemplateDirs[i] = os.ExpandEnv(server.TemplateDirs[i])
		fmt.Println("TemplateDir:", server.TemplateDirs[i])
	}

	server.initialized = true
	return nil
}

// Deploy copies all known resources (static and template files) to targetDir
func (server *Server) Deploy(targetDir string) error {
	panic("not implemented") // todo
}

func (server *Server) FindStaticFile(filename string) (filePath string, found bool, modified time.Time) {
	// todo optimize
	return dry.FileFindModified(append(server.BaseDirs, server.StaticDirs...), filename)
}

func (server *Server) FindTemplateFile(filename string) (filePath string, found bool, modified time.Time) {
	// todo optimize
	return dry.FileFindModified(append(server.BaseDirs, server.StaticDirs...), filename)
}
