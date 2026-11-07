package servicego

import (
	"net"
	"os"

	"github.com/kardianos/service"
)

// EnableReadyNotify makes systemd treat the service as started only after it
// calls Ready. It sets Type=notify through kardianos' SystemdScript option, so
// the unit written by -action deploy needs no drop-ins or extra files.
//
//	func (s *svc) Config() *service.Config {
//	  return servicego.EnableReadyNotify(s.DefaultConfig.Config())
//	}
//
// The service must call servicego.Ready once it is able to do its work, or
// systemd will fail the start after TimeoutStartSec.
func EnableReadyNotify(config *service.Config) *service.Config {
	if config.Option == nil {
		config.Option = service.KeyValue{}
	}
	config.Option["SystemdScript"] = systemdNotifyScript
	return config
}

// Ready tells systemd the service is ready. It does nothing when the process
// was not started by systemd with Type=notify (NOTIFY_SOCKET unset).
func Ready() error {
	return Notify("READY=1")
}

// Notify sends a raw sd_notify state string, such as "STATUS=...".
func Notify(state string) error {
	socket := os.Getenv("NOTIFY_SOCKET")
	if socket == "" {
		return nil
	}
	if socket[0] == '@' {
		socket = "\x00" + socket[1:]
	}
	conn, err := net.DialUnix("unixgram", nil, &net.UnixAddr{Name: socket, Net: "unixgram"})
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Write([]byte(state))
	return err
}

// systemdNotifyScript is kardianos/service v1.2.2's systemd template with
// Type=notify added. Keep in sync when bumping kardianos/service.
const systemdNotifyScript = `[Unit]
Description={{.Description}}
ConditionFileIsExecutable={{.Path|cmdEscape}}
{{range $i, $dep := .Dependencies}} 
{{$dep}} {{end}}

[Service]
Type=notify
StartLimitInterval=5
StartLimitBurst=10
ExecStart={{.Path|cmdEscape}}{{range .Arguments}} {{.|cmd}}{{end}}
{{if .ChRoot}}RootDirectory={{.ChRoot|cmd}}{{end}}
{{if .WorkingDirectory}}WorkingDirectory={{.WorkingDirectory|cmdEscape}}{{end}}
{{if .UserName}}User={{.UserName}}{{end}}
{{if .ReloadSignal}}ExecReload=/bin/kill -{{.ReloadSignal}} "$MAINPID"{{end}}
{{if .PIDFile}}PIDFile={{.PIDFile|cmd}}{{end}}
{{if and .LogOutput .HasOutputFileSupport -}}
StandardOutput=file:{{.LogDirectory}}/{{.Name}}.out
StandardError=file:{{.LogDirectory}}/{{.Name}}.err
{{- end}}
{{if gt .LimitNOFILE -1 }}LimitNOFILE={{.LimitNOFILE}}{{end}}
{{if .Restart}}Restart={{.Restart}}{{end}}
{{if .SuccessExitStatus}}SuccessExitStatus={{.SuccessExitStatus}}{{end}}
RestartSec=120
EnvironmentFile=-/etc/sysconfig/{{.Name}}

{{range $k, $v := .EnvVars -}}
Environment={{$k}}={{$v}}
{{end -}}

[Install]
WantedBy=multi-user.target
`
