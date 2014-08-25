package kutil

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

const (
	SystemV = 1
	Upstart = 2
	Systemd = 3
)

type LinuxService struct {
	itype                          int
	name, displayName, description string
	exePath                        string
}

func NewService(itype int, name, desc, exePath string) *LinuxService {
	return &LinuxService{itype, name, fmt.Sprintf("%v service", name), desc, exePath}
}

func isUpstart() bool {
	if _, err := os.Stat("/sbin/initctl"); err == nil {
		return true
	}
	return false
}

func isSystemd() bool {
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return true
	}
	return false
}

func SupportInitType(itype int) bool {
	if itype == SystemV {
		return true
	}
	if itype == Upstart {
		return isUpstart()
	}
	if itype == Systemd {
		return isSystemd()
	}
	return false
}

func GetPreferInitType() int {
	if isSystemd() {
		return Systemd
	}
	if isUpstart() {
		return Upstart
	}
	return SystemV
}

func configPath(itype int, name string) string {
	switch itype {
	case SystemV:
		return "/etc/init.d/" + name
	case Upstart:
		return "/etc/init/" + name + ".conf"
	case Systemd:
		return "/etc/systemd/system/" + name + ".service"
	default:
		return ""
	}
}

func (s *LinuxService) Install() error {
	confPath := configPath(s.itype, s.name)
	_, err := os.Stat(confPath)
	if err == nil {
		return fmt.Errorf("Init already exists: %s", confPath)
	}

	f, err := os.Create(confPath)
	if err != nil {
		return err
	}
	defer f.Close()

	path := s.exePath

	var to = &struct {
		Display     string
		Description string
		Path        string
		Name        string
	}{
		s.displayName,
		s.description,
		path,
		s.name,
	}

	var templ string
	switch s.itype {
	case SystemV:
		templ = systemVScript
	case Upstart:
		templ = upstartScript
	case Systemd:
		templ = systemdScript
	}
	mytemp := template.Must(template.New("Script").Parse(templ))
	err = mytemp.Execute(f, to)
	if err != nil {
		return err
	}

	if s.itype == SystemV {
		if err = os.Chmod(confPath, 0755); err != nil {
			return err
		}
		for _, i := range [...]string{"2", "3", "4", "5"} {
			if err = os.Symlink(confPath, "/etc/rc"+i+".d/S50"+s.name); err != nil {
				continue
			}
		}
		for _, i := range [...]string{"0", "1", "6"} {
			if err = os.Symlink(confPath, "/etc/rc"+i+".d/K02"+s.name); err != nil {
				continue
			}
		}
	}

	if s.itype == Systemd {
		return exec.Command("systemctl", "daemon-reload").Run()
	}

	return nil
}

func (s *LinuxService) Remove() error {
	if s.itype == Systemd {
		exec.Command("systemctl", "disable", s.name+".service").Run()
	}
	if err := os.Remove(configPath(s.itype, s.name)); err != nil {
		return err
	}
	return nil
}

func (s *LinuxService) Start() error {
	if s.itype == Upstart {
		return exec.Command("initctl", "start", s.name).Run()
	}
	if s.itype == Systemd {
		return exec.Command("systemctl", "start", s.name+".service").Run()
	}
	return exec.Command("service", s.name, "start").Run()
}

func (s *LinuxService) Stop() error {
	if s.itype == Upstart {
		return exec.Command("initctl", "stop", s.name).Run()
	}
	if s.itype == Systemd {
		return exec.Command("systemctl", "stop", s.name+".service").Run()
	}
	return exec.Command("service", s.name, "stop").Run()
}

const systemVScript = `#!/bin/sh
# For RedHat and cousins:
# chkconfig: - 99 01
# description: {{.Description}}
# processname: {{.Path}}

### BEGIN INIT INFO
# Provides:          {{.Path}}
# Required-Start:
# Required-Stop:
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: {{.Display}}
# Description:       {{.Description}}
### END INIT INFO

cmd="{{.Path}}"

name="{{.Name}}"
pid_file="/var/run/$name.pid"
stdout_log="/var/log/$name.log"
stderr_log="/var/log/$name.err"

get_pid() {
    cat "$pid_file"
}

is_running() {
    [ -f "$pid_file" ] && ps $(get_pid) > /dev/null 2>&1
}

case "$1" in
    start)
        if is_running; then
            echo "Already started"
        else
            echo "Starting $name"
            $cmd >> "$stdout_log" 2>> "$stderr_log" &
            echo $! > "$pid_file"
            if ! is_running; then
                echo "Unable to start, see $stdout_log and $stderr_log"
                exit 1
            fi
        fi
    ;;
    stop)
        if is_running; then
            echo -n "Stopping $name.."
            kill $(get_pid)
            for i in {1..10}
            do
                if ! is_running; then
                    break
                fi
                echo -n "."
                sleep 1
            done
            echo
            if is_running; then
                echo "Not stopped; may still be shutting down or shutdown may have failed"
                exit 1
            else
                echo "Stopped"
                if [ -f "$pid_file" ]; then
                    rm "$pid_file"
                fi
            fi
        else
            echo "Not running"
        fi
    ;;
    restart)
        $0 stop
        if is_running; then
            echo "Unable to stop, will not attempt to start"
            exit 1
        fi
        $0 start
    ;;	
  	reload)
	    if test -s "$pid_file" ; then
	      read mypid < "$pid_file"
	      kill -HUP $mypid && echo "Reloading $name service"
	      touch "$pid_file"
	    else
	      echo "$name's PID file could not be found!"
	      exit 1
	    fi
    ;;
    status)
        if is_running; then
            echo "Running"
        else
            echo "Stopped"
            exit 1
        fi
    ;;
    *)
    echo "Usage: $0 {start|stop|restart|reload|status}"
    exit 1
    ;;
esac
exit 0`

const upstartScript = `# {{.Description}}

description     "{{.Display}}"

start on filesystem or runlevel [2345]
stop on runlevel [!2345]

#setuid username

respawn
respawn limit 10 5
umask 022

console none

pre-start script
    test -x {{.Path}} || { stop; exit 0; }
end script

# Start
exec {{.Path}}
`

const systemdScript = `[Unit]
Description={{.Description}}
ConditionFileIsExecutable={{.Path}}

[Service]
StartLimitInterval=5
StartLimitBurst=10
ExecStart={{.Path}}

[Install]
WantedBy=multi-user.target
`
