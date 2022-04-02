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
	log     service.Logger
	src     string
	dest    string
}

func deploy(service Service) error {

	d := &installer{
		service: service,
		log:     service.Log(),
	}

	d.log.Info("deploying", d.service.Config().Name)

	return d.setSrcDest()
}

func (d *installer) setSrcDest() (err error) {
	d.log.Info("setSrcDest")
	if d.src, err = os.Executable(); err != nil {
		return
	}

	// TODO: configure based on service.Config.WorkingDirectory ?
	serviceName := d.service.Config().Name
	serviceDirectory := fmt.Sprintf("/opt/servicego/%s", serviceName)
	d.dest = fmt.Sprintf("%s/%s", serviceDirectory, serviceName)

	return d.checkPrevious()
}

func (d *installer) checkPrevious() (err error) {
	d.log.Info("checkPrevious")

	var stat os.FileInfo

	stat, err = os.Stat(d.dest)
	if stat != nil && !stat.IsDir() {
		return d.uninstallPrevious()
	}
	return d.installNew()
}

func (d *installer) uninstallPrevious() (err error) {
	d.log.Info("uninstallPrevious")
	if err = d.sudo(d.dest, "stop"); err != nil {
		return
	}
	if err = d.sudo(d.dest, "uninstall"); err != nil {
		return
	}
	return d.installNew()
}

func (d *installer) installNew() (err error) {
	d.log.Info("installNew")
	if err = os.MkdirAll(filepath.Dir(d.dest), 0755); err != nil {
		return
	}

	if err = d.copyBinary(d.src, d.dest); err != nil {
		return
	}

	if err = d.sudo(d.dest, "install"); err != nil {
		return
	}
	if err = d.sudo(d.dest, "start"); err != nil {
		return
	}
	d.log.Info("finished")
	return
}

func (d *installer) copyBinary(src, dest string) (err error) {

	var copyFrom *os.File
	var copyTo *os.File

	if copyFrom, err = os.Open(src); err != nil {
		return err
	}
	defer func() { _ = copyFrom.Close() }()

	if copyTo, err = os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0755); err != nil {
		return err
	}
	defer func() { _ = copyTo.Close() }()

	var written int64
	written, err = io.Copy(copyTo, copyFrom)
	if err != nil {
		return
	}
	d.log.Infof("copied %d bytes from %s to %s", written, src, dest)
	return
}

func (d *installer) sudo(binPath string, action string) error {
	output, err := exec.Command("sudo", binPath, "-action", action).CombinedOutput()
	if err != nil {
		d.log.Error(string(output))
		return err
	}
	return nil
}

func ServiceName() string { return filepath.Base(MustExecutable()) }

func MustExecutable() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return executable
}
