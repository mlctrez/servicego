package servicego

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kardianos/service"
)

type installer struct {
	service Service
	svc     service.Service
	log     service.Logger
	src     string
	dest    string
}

func deploy(service Service, svc service.Service) error {

	d := &installer{
		service: service,
		svc:     svc,
		log:     service.Log(),
	}

	return d.setSrcDest()
}

func (d *installer) setSrcDest() (err error) {

	if d.src, err = os.Executable(); err != nil {
		return
	}

	cfg := d.service.Config()

	serviceName := cfg.Name
	d.dest = filepath.Join(cfg.WorkingDirectory, serviceName)

	return d.checkPrevious()
}

func (d *installer) checkPrevious() (err error) {

	// Use the service manager state rather than the presence of the old
	// binary: the unit may exist without the binary, and may be stopped.
	status, err := d.svc.Status()
	if err == service.ErrNotInstalled {
		return d.installNew()
	}
	if err != nil {
		return fmt.Errorf("status %s: %w", d.service.Config().Name, err)
	}

	if status == service.StatusRunning {
		if err = d.svc.Stop(); err != nil {
			return fmt.Errorf("stop %s: %w", d.service.Config().Name, err)
		}
	}

	if err = d.svc.Uninstall(); err != nil {
		return fmt.Errorf("uninstall %s: %w", d.service.Config().Name, err)
	}

	return d.installNew()
}

func (d *installer) installNew() (err error) {
	if err = os.MkdirAll(filepath.Dir(d.dest), 0755); err != nil {
		return
	}

	if err = d.copyBinary(d.src, d.dest); err != nil {
		return
	}

	if err = d.command(d.dest, "install"); err != nil {
		return
	}

	if err = d.command(d.dest, "start"); err != nil {
		return
	}

	return
}

func (d *installer) copyBinary(src, dest string) (err error) {

	var copyFrom *os.File
	var copyTo *os.File

	if copyFrom, err = os.Open(src); err != nil {
		return err
	}
	defer func() { _ = copyFrom.Close() }()

	if copyTo, err = os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755); err != nil {
		return err
	}
	defer func() { _ = copyTo.Close() }()

	_, err = io.Copy(copyTo, copyFrom)
	return
}

func (d *installer) command(binPath string, action string) error {
	output, err := exec.Command(binPath, "-action", action).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s -action %s: %w: %s", binPath, action, err, output)
	}
	return nil
}

func ServiceName() string      { return filepath.Base(MustExecutable()) }
func ServiceDirectory() string { return fmt.Sprintf("/opt/servicego/%s", ServiceName()) }

var executableName string

func MustExecutable() string {
	if executableName != "" {
		return executableName
	}

	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	executableName = executable
	return executable
}
