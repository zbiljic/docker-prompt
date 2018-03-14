package cmd

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	var suggests []prompt.Suggest
	commandArgs := excludeOptions(args)
	switch commandArgs[0] {
	case "attach":
		suggests = append(flagAttach, flagGlobal...)
	case "build":
		suggests = append(flagBuild, flagGlobal...)
	case "commit":
		suggests = append(flagCommit, flagGlobal...)
	case "cp":
		suggests = append(flagCp, flagGlobal...)
	case "create":
		suggests = append(flagCreate, flagGlobal...)
	case "deploy":
		suggests = append(flagDeploy, flagGlobal...)
	case "diff":
		suggests = append(flagDiff, flagGlobal...)
	case "events":
		suggests = append(flagEvents, flagGlobal...)
	case "exec":
		suggests = append(flagExec, flagGlobal...)
	case "export":
		suggests = append(flagExport, flagGlobal...)
	case "history":
		suggests = append(flagHistory, flagGlobal...)
	case "images":
		suggests = append(flagImages, flagGlobal...)
	case "import":
		suggests = append(flagImport, flagGlobal...)
	case "info":
		suggests = append(flagInfo, flagGlobal...)
	case "inspect":
		suggests = append(flagInspect, flagGlobal...)
	case "kill":
		suggests = append(flagKill, flagGlobal...)
	case "load":
		suggests = append(flagLoad, flagGlobal...)
	case "login":
		suggests = append(flagLogin, flagGlobal...)
	case "logout":
		suggests = append(flagLogout, flagGlobal...)
	case "logs":
		suggests = append(flagLogs, flagGlobal...)
	case "pause":
		suggests = append(flagPause, flagGlobal...)
	case "port":
		suggests = append(flagPort, flagGlobal...)
	case "ps":
		suggests = append(flagPs, flagGlobal...)
	case "pull":
		suggests = append(flagPull, flagGlobal...)
	case "push":
		suggests = append(flagPush, flagGlobal...)
	case "rename":
		suggests = append(flagRename, flagGlobal...)
	case "restart":
		suggests = append(flagRestart, flagGlobal...)
	case "rm":
		suggests = append(flagRm, flagGlobal...)
	case "rmi":
		suggests = append(flagRmi, flagGlobal...)
	case "run":
		suggests = append(flagRun, flagGlobal...)
	case "save":
		suggests = append(flagSave, flagGlobal...)
	case "search":
		suggests = append(flagSearch, flagGlobal...)
	case "start":
		suggests = append(flagStart, flagGlobal...)
	case "stats":
		suggests = append(flagStats, flagGlobal...)
	case "stop":
		suggests = append(flagStop, flagGlobal...)
	case "tag":
		suggests = append(flagTag, flagGlobal...)
	case "top":
		suggests = append(flagTop, flagGlobal...)
	case "unpause":
		suggests = append(flagUnpause, flagGlobal...)
	case "update":
		suggests = append(flagUpdate, flagGlobal...)
	case "version":
		suggests = append(flagVersion, flagGlobal...)
	case "wait":
		suggests = append(flagWait, flagGlobal...)
	default:
		suggests = optionHelp
	}

	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}

var optionHelp = []prompt.Suggest{
	{Text: "-h"},
	{Text: "--help"},
}

var flagGlobal = []prompt.Suggest{
	{Text: "--config", Description: `Location of client config files`},
	{Text: "-D", Description: `Enable debug mode`},
	{Text: "--debug", Description: `Enable debug mode`},
	{Text: "-H", Description: `Daemon socket(s) to connect to`},
	{Text: "--host", Description: `Daemon socket(s) to connect to`},
	{Text: "-l", Description: `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")`},
	{Text: "--log-level", Description: `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")`},
	{Text: "--tls", Description: `Use TLS; implied by --tlsverify`},
	{Text: "--tlscacert", Description: `Trust certs signed only by this CA`},
	{Text: "--tlscert", Description: `Path to TLS certificate file`},
	{Text: "--tlskey", Description: `Path to TLS key file`},
	{Text: "--tlsverify", Description: `Use TLS and verify the remote`},
	{Text: "-v", Description: `Print version information and quit`},
	{Text: "--version", Description: `Print version information and quit`},
}

var flagAttach = []prompt.Suggest{
	{Text: "--detach-keys", Description: `Override the key sequence for detaching a container`},
	{Text: "--no-stdin", Description: `Do not attach STDIN`},
	{Text: "--sig-proxy", Description: `Proxy all received signals to the process (default true)`},
}

var flagBuild = []prompt.Suggest{
	{Text: "--add-host", Description: `Add a custom host-to-IP mapping (host:ip)`},
	{Text: "--build-arg", Description: `Set build-time variables`},
	{Text: "--cache-from", Description: `Images to consider as cache sources`},
	{Text: "--cgroup-parent", Description: `Optional parent cgroup for the container`},
	{Text: "--compress", Description: `Compress the build context using gzip`},
	{Text: "--cpu-period", Description: `Limit the CPU CFS (Completely Fair Scheduler) period`},
	{Text: "--cpu-quota", Description: `Limit the CPU CFS (Completely Fair Scheduler) quota`},
	{Text: "-c", Description: `CPU shares (relative weight)`},
	{Text: "--cpu-shares", Description: `CPU shares (relative weight)`},
	{Text: "--cpuset-cpus", Description: `CPUs in which to allow execution (0-3, 0,1)`},
	{Text: "--cpuset-mems", Description: `MEMs in which to allow execution (0-3, 0,1)`},
	{Text: "--disable-content-trust", Description: `Skip image verification (default true)`},
	{Text: "-f", Description: `Name of the Dockerfile (Default is 'PATH/Dockerfile')`},
	{Text: "--file", Description: `Name of the Dockerfile (Default is 'PATH/Dockerfile')`},
	{Text: "--force-rm", Description: `Always remove intermediate containers`},
	{Text: "--iidfile", Description: `Write the image ID to the file`},
	{Text: "--isolation", Description: `Container isolation technology`},
	{Text: "--label", Description: `Set metadata for an image`},
	{Text: "-m", Description: `Memory limit`},
	{Text: "--memory", Description: `Memory limit`},
	{Text: "--memory-swap", Description: `Swap limit equal to memory plus swap: '-1' to enable unlimited swap`},
	{Text: "--network", Description: `Set the networking mode for the RUN instructions during build (default "default")`},
	{Text: "--no-cache", Description: `Do not use cache when building the image`},
	{Text: "--platform", Description: `Set platform if server is multi-platform capable`},
	{Text: "--pull", Description: `Always attempt to pull a newer version of the image`},
	{Text: "-q", Description: `Suppress the build output and print image ID on success`},
	{Text: "--quiet", Description: `Suppress the build output and print image ID on success`},
	{Text: "--rm", Description: `Remove intermediate containers after a successful build (default true)`},
	{Text: "--security-opt", Description: `Security options`},
	{Text: "--shm-size", Description: `Size of /dev/shm`},
	{Text: "--squash", Description: `Squash newly built layers into a single new layer`},
	{Text: "--stream", Description: `Stream attaches to server to negotiate build context`},
	{Text: "-t", Description: `Name and optionally a tag in the 'name:tag' format`},
	{Text: "--tag", Description: `Name and optionally a tag in the 'name:tag' format`},
	{Text: "--target", Description: `Set the target build stage to build.`},
	{Text: "--ulimit", Description: `Ulimit options (default [])`},
}

var flagCommit = []prompt.Suggest{
	{Text: "-a", Description: `Author (e.g., "John Hannibal Smith <hannibal@a-team.com>")`},
	{Text: "--author", Description: `Author (e.g., "John Hannibal Smith <hannibal@a-team.com>")`},
	{Text: "-c", Description: `Apply Dockerfile instruction to the created image`},
	{Text: "--change", Description: `Apply Dockerfile instruction to the created image`},
	{Text: "-m", Description: `Commit message`},
	{Text: "--message", Description: `Commit message`},
	{Text: "-p", Description: `Pause container during commit (default true)`},
	{Text: "--pause", Description: `Pause container during commit (default true)`},
}

var flagCp = []prompt.Suggest{
	{Text: "-a", Description: `Archive mode (copy all uid/gid information)`},
	{Text: "--archive", Description: `Archive mode (copy all uid/gid information)`},
	{Text: "-L", Description: `Always follow symbol link in SRC_PATH`},
	{Text: "--follow-link", Description: `Always follow symbol link in SRC_PATH`},
}

var flagCreate = []prompt.Suggest{
	{Text: "--add-host", Description: `Add a custom host-to-IP mapping (host:ip)`},
	{Text: "-a", Description: `Attach to STDIN, STDOUT or STDERR`},
	{Text: "--attach", Description: `Attach to STDIN, STDOUT or STDERR`},
	{Text: "--blkio-weight", Description: `Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)`},
	{Text: "--blkio-weight-device", Description: `Block IO weight (relative device weight) (default [])`},
	{Text: "--cap-add", Description: `Add Linux capabilities`},
	{Text: "--cap-drop", Description: `Drop Linux capabilities`},
	{Text: "--cgroup-parent", Description: `Optional parent cgroup for the container`},
	{Text: "--cidfile", Description: `Write the container ID to the file`},
	{Text: "--cpu-period", Description: `Limit CPU CFS (Completely Fair Scheduler) period`},
	{Text: "--cpu-quota", Description: `Limit CPU CFS (Completely Fair Scheduler) quota`},
	{Text: "--cpu-rt-period", Description: `Limit CPU real-time period in microseconds`},
	{Text: "--cpu-rt-runtime", Description: `Limit CPU real-time runtime in microseconds`},
	{Text: "-c", Description: `CPU shares (relative weight)`},
	{Text: "--cpu-shares", Description: `CPU shares (relative weight)`},
	{Text: "--cpus", Description: `Number of CPUs`},
	{Text: "--cpuset-cpus", Description: `CPUs in which to allow execution (0-3, 0,1)`},
	{Text: "--cpuset-mems", Description: `MEMs in which to allow execution (0-3, 0,1)`},
	{Text: "--device", Description: `Add a host device to the container`},
	{Text: "--device-cgroup-rule", Description: `Add a rule to the cgroup allowed devices list`},
	{Text: "--device-read-bps", Description: `Limit read rate (bytes per second) from a device (default [])`},
	{Text: "--device-read-iops", Description: `Limit read rate (IO per second) from a device (default [])`},
	{Text: "--device-write-bps", Description: `Limit write rate (bytes per second) to a device (default [])`},
	{Text: "--device-write-iops", Description: `Limit write rate (IO per second) to a device (default [])`},
	{Text: "--disable-content-trust", Description: `Skip image verification (default true)`},
	{Text: "--dns", Description: `Set custom DNS servers`},
	{Text: "--dns-option", Description: `Set DNS options`},
	{Text: "--dns-search", Description: `Set custom DNS search domains`},
	{Text: "--entrypoint", Description: `Overwrite the default ENTRYPOINT of the image`},
	{Text: "-e", Description: `Set environment variables`},
	{Text: "--env", Description: `Set environment variables`},
	{Text: "--env-file", Description: `Read in a file of environment variables`},
	{Text: "--expose", Description: `Expose a port or a range of ports`},
	{Text: "--group-add", Description: `Add additional groups to join`},
	{Text: "--health-cmd", Description: `Command to run to check health`},
	{Text: "--health-interval", Description: `Time between running the check (ms|s|m|h) (default 0s)`},
	{Text: "--health-retries", Description: `Consecutive failures needed to report unhealthy`},
	{Text: "--health-start-period", Description: `Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)`},
	{Text: "--health-timeout", Description: `Maximum time to allow one check to run (ms|s|m|h) (default 0s)`},
	{Text: "--help", Description: `Print usage`},
	{Text: "-h", Description: `Container host name`},
	{Text: "--hostname", Description: `Container host name`},
	{Text: "--init", Description: `Run an init inside the container that forwards signals and reaps processes`},
	{Text: "-i", Description: `Keep STDIN open even if not attached`},
	{Text: "--interactive", Description: `Keep STDIN open even if not attached`},
	{Text: "--ip", Description: `IPv4 address (e.g., 172.30.100.104)`},
	{Text: "--ip6", Description: `IPv6 address (e.g., 2001:db8::33)`},
	{Text: "--ipc", Description: `IPC mode to use`},
	{Text: "--isolation", Description: `Container isolation technology`},
	{Text: "--kernel-memory", Description: `Kernel memory limit`},
	{Text: "-l", Description: `Set meta data on a container`},
	{Text: "--label", Description: `Set meta data on a container`},
	{Text: "--label-file", Description: `Read in a line delimited file of labels`},
	{Text: "--link", Description: `Add link to another container`},
	{Text: "--link-local-ip", Description: `Container IPv4/IPv6 link-local addresses`},
	{Text: "--log-driver", Description: `Logging driver for the container`},
	{Text: "--log-opt", Description: `Log driver options`},
	{Text: "--mac-address", Description: `Container MAC address (e.g., 92:d0:c6:0a:29:33)`},
	{Text: "-m", Description: `Memory limit`},
	{Text: "--memory", Description: `Memory limit`},
	{Text: "--memory-reservation", Description: `Memory soft limit`},
	{Text: "--memory-swap", Description: `Swap limit equal to memory plus swap: '-1' to enable unlimited swap`},
	{Text: "--memory-swappiness", Description: `Tune container memory swappiness (0 to 100) (default -1)`},
	{Text: "--mount", Description: `Attach a filesystem mount to the container`},
	{Text: "--name", Description: `Assign a name to the container`},
	{Text: "--network", Description: `Connect a container to a network (default "default")`},
	{Text: "--network-alias", Description: `Add network-scoped alias for the container`},
	{Text: "--no-healthcheck", Description: `Disable any container-specified HEALTHCHECK`},
	{Text: "--oom-kill-disable", Description: `Disable OOM Killer`},
	{Text: "--oom-score-adj", Description: `Tune host's OOM preferences (-1000 to 1000)`},
	{Text: "--pid", Description: `PID namespace to use`},
	{Text: "--pids-limit", Description: `Tune container pids limit (set -1 for unlimited)`},
	{Text: "--platform", Description: `Set platform if server is multi-platform capable`},
	{Text: "--privileged", Description: `Give extended privileges to this container`},
	{Text: "-p", Description: `Publish a container's port(s) to the host`},
	{Text: "--publish", Description: `Publish a container's port(s) to the host`},
	{Text: "-P", Description: `Publish all exposed ports to random ports`},
	{Text: "--publish-all", Description: `Publish all exposed ports to random ports`},
	{Text: "--read-only", Description: `Mount the container's root filesystem as read only`},
	{Text: "--restart", Description: `Restart policy to apply when a container exits (default "no")`},
	{Text: "--rm", Description: `Automatically remove the container when it exits`},
	{Text: "--runtime", Description: `Runtime to use for this container`},
	{Text: "--security-opt", Description: `Security Options`},
	{Text: "--shm-size", Description: `Size of /dev/shm`},
	{Text: "--stop-signal", Description: `Signal to stop a container (default "SIGTERM")`},
	{Text: "--stop-timeout", Description: `Timeout (in seconds) to stop a container`},
	{Text: "--storage-opt", Description: `Storage driver options for the container`},
	{Text: "--sysctl map", Description: `Sysctl options (default map[])`},
	{Text: "--tmpfs", Description: `Mount a tmpfs directory`},
	{Text: "-t", Description: `Allocate a pseudo-TTY`},
	{Text: "--tty", Description: `Allocate a pseudo-TTY`},
	{Text: "--ulimit ulimit", Description: `Ulimit options (default [])`},
	{Text: "-u", Description: `Username or UID (format: <name|uid>[:<group|gid>])`},
	{Text: "--user", Description: `Username or UID (format: <name|uid>[:<group|gid>])`},
	{Text: "--userns", Description: `User namespace to use`},
	{Text: "--uts", Description: `UTS namespace to use`},
	{Text: "-v", Description: `Bind mount a volume`},
	{Text: "--volume", Description: `Bind mount a volume`},
	{Text: "--volume-driver", Description: `Optional volume driver for the container`},
	{Text: "--volumes-from", Description: `Mount volumes from the specified container(s)`},
	{Text: "-w", Description: `Working directory inside the container`},
	{Text: "--workdir", Description: `Working directory inside the container`},
}

var flagDeploy = []prompt.Suggest{
	{Text: "--bundle-file", Description: `Path to a Distributed Application Bundle file`},
	{Text: "-c", Description: `Path to a Compose file`},
	{Text: "--compose-file", Description: `Path to a Compose file`},
	{Text: "--prune", Description: `Prune services that are no longer referenced`},
	{Text: "--resolve-image", Description: `Query the registry to resolve image digest and supported platforms ("always"|"changed"|"never") (default "always")`},
	{Text: "--with-registry-auth", Description: `Send registry authentication details to Swarm agents`},
}

var flagDiff = []prompt.Suggest{}

var flagEvents = []prompt.Suggest{
	{Text: "-f", Description: `Filter output based on conditions provided`},
	{Text: "--filter", Description: `Filter output based on conditions provided`},
	{Text: "--format", Description: `Format the output using the given Go template`},
	{Text: "--since", Description: `Show all events created since timestamp`},
	{Text: "--until", Description: `Stream events until this timestamp`},
}

var flagExec = []prompt.Suggest{
	{Text: "-d", Description: `Detached mode: run command in the background`},
	{Text: "--detach", Description: `Detached mode: run command in the background`},
	{Text: "--detach-keys", Description: `Override the key sequence for detaching a container`},
	{Text: "-e", Description: `Set environment variables`},
	{Text: "--env", Description: `Set environment variables`},
	{Text: "-i", Description: `Keep STDIN open even if not attached`},
	{Text: "--interactive", Description: `Keep STDIN open even if not attached`},
	{Text: "--privileged", Description: `Give extended privileges to the command`},
	{Text: "-t", Description: `Allocate a pseudo-TTY`},
	{Text: "--tty", Description: `Allocate a pseudo-TTY`},
	{Text: "-u", Description: `Username or UID (format: <name|uid>[:<group|gid>])`},
	{Text: "--user", Description: `Username or UID (format: <name|uid>[:<group|gid>])`},
	{Text: "-w", Description: `Working directory inside the container`},
	{Text: "--workdir", Description: `Working directory inside the container`},
}

var flagExport = []prompt.Suggest{
	{Text: "-o", Description: `Write to a file, instead of STDOUT`},
	{Text: "--output", Description: `Write to a file, instead of STDOUT`},
}

var flagHistory = []prompt.Suggest{
	{Text: "--format", Description: `Pretty-print images using a Go template`},
	{Text: "-H", Description: `Print sizes and dates in human readable format (default true)`},
	{Text: "--human", Description: `Print sizes and dates in human readable format (default true)`},
	{Text: "--no-trunc", Description: `Don't truncate output`},
	{Text: "-q", Description: `Only show numeric IDs`},
	{Text: "--quiet", Description: `Only show numeric IDs`},
}

var flagImages = []prompt.Suggest{
	{Text: "-a", Description: `Show all images (default hides intermediate images)`},
	{Text: "--all", Description: `Show all images (default hides intermediate images)`},
	{Text: "--digests", Description: `Show digests`},
	{Text: "-f", Description: `Filter output based on conditions provided`},
	{Text: "--filter", Description: `Filter output based on conditions provided`},
	{Text: "--format", Description: `Pretty-print images using a Go template`},
	{Text: "--no-trunc", Description: `Don't truncate output`},
	{Text: "-q", Description: `Only show numeric IDs`},
	{Text: "--quiet", Description: `Only show numeric IDs`},
}

var flagImport = []prompt.Suggest{
	{Text: "-c", Description: `Apply Dockerfile instruction to the created image`},
	{Text: "--change", Description: `Apply Dockerfile instruction to the created image`},
	{Text: "-m", Description: `Set commit message for imported image`},
	{Text: "--message", Description: `Set commit message for imported image`},
}

var flagInfo = []prompt.Suggest{
	{Text: "-f", Description: `Format the output using the given Go template`},
	{Text: "--format", Description: `Format the output using the given Go template`},
}

var flagInspect = []prompt.Suggest{
	{Text: "-f", Description: `Format the output using the given Go template`},
	{Text: "--format", Description: `Format the output using the given Go template`},
	{Text: "-s", Description: `Display total file sizes if the type is container`},
	{Text: "--size", Description: `Display total file sizes if the type is container`},
	{Text: "--type", Description: `Return JSON for specified type`},
}

var flagKill = []prompt.Suggest{
	{Text: "-s", Description: `Signal to send to the container (default "KILL")`},
	{Text: "--signal", Description: `Signal to send to the container (default "KILL")`},
}

var flagLoad = []prompt.Suggest{
	{Text: "-i", Description: `Read from tar archive file, instead of STDIN`},
	{Text: "--input", Description: `Read from tar archive file, instead of STDIN`},
	{Text: "-q", Description: `Suppress the load output`},
	{Text: "--quiet", Description: `Suppress the load output`},
}

var flagLogin = []prompt.Suggest{
	{Text: "-p", Description: `Password`},
	{Text: "--password", Description: `Password`},
	{Text: "--password-stdin", Description: `Take the password from stdin`},
	{Text: "-u", Description: `Username`},
	{Text: "--username", Description: `Username`},
}

var flagLogout = []prompt.Suggest{}

var flagLogs = []prompt.Suggest{
	{Text: "--details", Description: `Show extra details provided to logs`},
	{Text: "-f", Description: `Follow log output`},
	{Text: "--follow", Description: `Follow log output`},
	{Text: "--since", Description: `Show logs since timestamp (e.g. 2013-01-02T13:23:37) or relative (e.g. 42m for 42 minutes)`},
	{Text: "--tail", Description: `Number of lines to show from the end of the logs (default "all")`},
	{Text: "-t", Description: `Show timestamps`},
	{Text: "--timestamps", Description: `Show timestamps`},
	{Text: "--until", Description: `Show logs before a timestamp (e.g. 2013-01-02T13:23:37) or relative (e.g. 42m for 42 minutes)`},
}

var flagPause = []prompt.Suggest{}

var flagPort = []prompt.Suggest{}

var flagPs = []prompt.Suggest{
	{Text: "-a", Description: `Show all containers (default shows just running)`},
	{Text: "--all", Description: `Show all containers (default shows just running)`},
	{Text: "-f", Description: `Filter output based on conditions provided`},
	{Text: "--filter", Description: `Filter output based on conditions provided`},
	{Text: "--format", Description: `Pretty-print containers using a Go template`},
	{Text: "-n", Description: `Show n last created containers (includes all states) (default -1)`},
	{Text: "--last", Description: `Show n last created containers (includes all states) (default -1)`},
	{Text: "-l", Description: `Show the latest created container (includes all states)`},
	{Text: "--latest", Description: `Show the latest created container (includes all states)`},
	{Text: "--no-trunc", Description: `Don't truncate output`},
	{Text: "-q", Description: `Only display numeric IDs`},
	{Text: "--quiet", Description: `Only display numeric IDs`},
	{Text: "-s", Description: `Display total file sizes`},
	{Text: "--size", Description: `Display total file sizes`},
}

var flagPull = []prompt.Suggest{
	{Text: "-a", Description: `Download all tagged images in the repository`},
	{Text: "--all-tags", Description: `Download all tagged images in the repository`},
	{Text: "--disable-content-trust", Description: `Skip image verification (default true)`},
	{Text: "--platform", Description: `Set platform if server is multi-platform capable`},
}

var flagPush = []prompt.Suggest{
	{Text: "--disable-content-trust", Description: `Skip image signing (default true)`},
}

var flagRename = []prompt.Suggest{}

var flagRestart = []prompt.Suggest{
	{Text: "-t", Description: `Seconds to wait for stop before killing the container (default 10)`},
	{Text: "--time", Description: `Seconds to wait for stop before killing the container (default 10)`},
}

var flagRm = []prompt.Suggest{
	{Text: "-f", Description: `Force the removal of a running container (uses SIGKILL)`},
	{Text: "--force", Description: `Force the removal of a running container (uses SIGKILL)`},
	{Text: "-l", Description: `Remove the specified link`},
	{Text: "--link", Description: `Remove the specified link`},
	{Text: "-v", Description: `Remove the volumes associated with the container`},
	{Text: "--volumes", Description: `Remove the volumes associated with the container`},
}

var flagRmi = []prompt.Suggest{
	{Text: "-f", Description: `Force removal of the image`},
	{Text: "--force", Description: `Force removal of the image`},
	{Text: "--no-prune", Description: `Do not delete untagged parents`},
}

var flagRun = []prompt.Suggest{
	{Text: "--add-host", Description: `Add a custom host-to-IP mapping (host:ip)`},
	{Text: "-a", Description: `Attach to STDIN, STDOUT or STDERR`},
	{Text: "--attach", Description: `Attach to STDIN, STDOUT or STDERR`},
	{Text: "--blkio-weight", Description: `Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)`},
	{Text: "--blkio-weight-device", Description: `Block IO weight (relative device weight) (default [])`},
	{Text: "--cap-add", Description: `Add Linux capabilities`},
	{Text: "--cap-drop", Description: `Drop Linux capabilities`},
	{Text: "--cgroup-parent", Description: `Optional parent cgroup for the container`},
	{Text: "--cidfile", Description: `Write the container ID to the file`},
	{Text: "--cpu-period", Description: `Limit CPU CFS (Completely Fair Scheduler) period`},
	{Text: "--cpu-quota", Description: `Limit CPU CFS (Completely Fair Scheduler) quota`},
	{Text: "--cpu-rt-period", Description: `Limit CPU real-time period in microseconds`},
	{Text: "--cpu-rt-runtime", Description: `Limit CPU real-time runtime in microseconds`},
	{Text: "-c", Description: `CPU shares (relative weight)`},
	{Text: "--cpu-shares", Description: `CPU shares (relative weight)`},
	{Text: "--cpus decimal", Description: `Number of CPUs`},
	{Text: "--cpuset-cpus", Description: `CPUs in which to allow execution (0-3, 0,1)`},
	{Text: "--cpuset-mems", Description: `MEMs in which to allow execution (0-3, 0,1)`},
	{Text: "-d", Description: `Run container in background and print container ID`},
	{Text: "--detach", Description: `Run container in background and print container ID`},
	{Text: "--detach-keys", Description: `Override the key sequence for detaching a container`},
	{Text: "--device", Description: `Add a host device to the container`},
	{Text: "--device-cgroup-rule", Description: `Add a rule to the cgroup allowed devices list`},
	{Text: "--device-read-bps", Description: `Limit read rate (bytes per second) from a device (default [])`},
	{Text: "--device-read-iops", Description: `Limit read rate (IO per second) from a device (default [])`},
	{Text: "--device-write-bps", Description: `Limit write rate (bytes per second) to a device (default [])`},
	{Text: "--device-write-iops", Description: `Limit write rate (IO per second) to a device (default [])`},
	{Text: "--disable-content-trust", Description: `Skip image verification (default true)`},
	{Text: "--dns", Description: `Set custom DNS servers`},
	{Text: "--dns-option", Description: `Set DNS options`},
	{Text: "--dns-search", Description: `Set custom DNS search domains`},
	{Text: "--entrypoint", Description: `Overwrite the default ENTRYPOINT of the image`},
	{Text: "-e", Description: `Set environment variables`},
	{Text: "--env", Description: `Set environment variables`},
	{Text: "--env-file", Description: `Read in a file of environment variables`},
	{Text: "--expose", Description: `Expose a port or a range of ports`},
	{Text: "--group-add", Description: `Add additional groups to join`},
	{Text: "--health-cmd", Description: `Command to run to check health`},
	{Text: "--health-interval", Description: `Time between running the check (ms|s|m|h) (default 0s)`},
	{Text: "--health-retries", Description: `Consecutive failures needed to report unhealthy`},
	{Text: "--health-start-period", Description: `Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)`},
	{Text: "--health-timeout", Description: `Maximum time to allow one check to run (ms|s|m|h) (default 0s)`},
	{Text: "--help", Description: `Print usage`},
	{Text: "-h", Description: `Container host name`},
	{Text: "--hostname", Description: `Container host name`},
	{Text: "--init", Description: `Run an init inside the container that forwards signals and reaps processes`},
	{Text: "-i", Description: `Keep STDIN open even if not attached`},
	{Text: "--interactive", Description: `Keep STDIN open even if not attached`},
	{Text: "--ip", Description: `IPv4 address (e.g., 172.30.100.104)`},
	{Text: "--ip6", Description: `IPv6 address (e.g., 2001:db8::33)`},
	{Text: "--ipc", Description: `IPC mode to use`},
	{Text: "--isolation", Description: `Container isolation technology`},
	{Text: "--kernel-memory", Description: `Kernel memory limit`},
	{Text: "-l", Description: `Set meta data on a container`},
	{Text: "--label", Description: `Set meta data on a container`},
	{Text: "--label-file", Description: `Read in a line delimited file of labels`},
	{Text: "--link", Description: `Add link to another container`},
	{Text: "--link-local-ip", Description: `Container IPv4/IPv6 link-local addresses`},
	{Text: "--log-driver", Description: `Logging driver for the container`},
	{Text: "--log-opt", Description: `Log driver options`},
	{Text: "--mac-address", Description: `Container MAC address (e.g., 92:d0:c6:0a:29:33)`},
	{Text: "-m", Description: `Memory limit`},
	{Text: "--memory", Description: `Memory limit`},
	{Text: "--memory-reservation", Description: `Memory soft limit`},
	{Text: "--memory-swap", Description: `Swap limit equal to memory plus swap: '-1' to enable unlimited swap`},
	{Text: "--memory-swappiness", Description: `Tune container memory swappiness (0 to 100) (default -1)`},
	{Text: "--mount", Description: `Attach a filesystem mount to the container`},
	{Text: "--name", Description: `Assign a name to the container`},
	{Text: "--network", Description: `Connect a container to a network (default "default")`},
	{Text: "--network-alias", Description: `Add network-scoped alias for the container`},
	{Text: "--no-healthcheck", Description: `Disable any container-specified HEALTHCHECK`},
	{Text: "--oom-kill-disable", Description: `Disable OOM Killer`},
	{Text: "--oom-score-adj", Description: `Tune host's OOM preferences (-1000 to 1000)`},
	{Text: "--pid", Description: `PID namespace to use`},
	{Text: "--pids-limit", Description: `Tune container pids limit (set -1 for unlimited)`},
	{Text: "--platform", Description: `Set platform if server is multi-platform capable`},
	{Text: "--privileged", Description: `Give extended privileges to this container`},
	{Text: "-p", Description: `Publish a container's port(s) to the host`},
	{Text: "--publish", Description: `Publish a container's port(s) to the host`},
	{Text: "-P", Description: `Publish all exposed ports to random ports`},
	{Text: "--publish-all", Description: `Publish all exposed ports to random ports`},
	{Text: "--read-only", Description: `Mount the container's root filesystem as read only`},
	{Text: "--restart", Description: `Restart policy to apply when a container exits (default "no")`},
	{Text: "--rm", Description: `Automatically remove the container when it exits`},
	{Text: "--runtime", Description: `Runtime to use for this container`},
	{Text: "--security-opt", Description: `Security Options`},
	{Text: "--shm-size", Description: `Size of /dev/shm`},
	{Text: "--sig-proxy", Description: `Proxy received signals to the process (default true)`},
	{Text: "--stop-signal", Description: `Signal to stop a container (default "SIGTERM")`},
	{Text: "--stop-timeout", Description: `Timeout (in seconds) to stop a container`},
	{Text: "--storage-opt", Description: `Storage driver options for the container`},
	{Text: "--sysctl map", Description: `Sysctl options (default map[])`},
	{Text: "--tmpfs", Description: `Mount a tmpfs directory`},
	{Text: "-t", Description: `Allocate a pseudo-TTY`},
	{Text: "--tty", Description: `Allocate a pseudo-TTY`},
	{Text: "--ulimit ulimit", Description: `Ulimit options (default [])`},
	{Text: "-u", Description: `Username or UID (format: <name|uid>[:<group|gid>])`},
	{Text: "--user", Description: `Username or UID (format: <name|uid>[:<group|gid>])`},
	{Text: "--userns", Description: `User namespace to use`},
	{Text: "--uts", Description: `UTS namespace to use`},
	{Text: "-v", Description: `Bind mount a volume`},
	{Text: "--volume", Description: `Bind mount a volume`},
	{Text: "--volume-driver", Description: `Optional volume driver for the container`},
	{Text: "--volumes-from", Description: `Mount volumes from the specified container(s)`},
	{Text: "-w", Description: `Working directory inside the container`},
	{Text: "--workdir", Description: `Working directory inside the container`},
}

var flagSave = []prompt.Suggest{
	{Text: "-o", Description: `Write to a file, instead of STDOUT`},
	{Text: "--output", Description: `Write to a file, instead of STDOUT`},
}

var flagSearch = []prompt.Suggest{
	{Text: "-f", Description: `Filter output based on conditions provided`},
	{Text: "--filter", Description: `Filter output based on conditions provided`},
	{Text: "--format", Description: `Pretty-print search using a Go template`},
	{Text: "--limit", Description: `Max number of search results (default 25)`},
	{Text: "--no-trunc", Description: `Don't truncate output`},
}

var flagStart = []prompt.Suggest{
	{Text: "-a", Description: `Attach STDOUT/STDERR and forward signals`},
	{Text: "--attach", Description: `Attach STDOUT/STDERR and forward signals`},
	{Text: "--checkpoint", Description: `Restore from this checkpoint`},
	{Text: "--checkpoint-dir", Description: `Use a custom checkpoint storage directory`},
	{Text: "--detach-keys", Description: `Override the key sequence for detaching a container`},
	{Text: "-i", Description: `Attach container's STDIN`},
	{Text: "--interactive", Description: `Attach container's STDIN`},
}

var flagStats = []prompt.Suggest{
	{Text: "-a", Description: `Show all containers (default shows just running)`},
	{Text: "--all", Description: `Show all containers (default shows just running)`},
	{Text: "--format", Description: `Pretty-print images using a Go template`},
	{Text: "--no-stream", Description: `Disable streaming stats and only pull the first result`},
	{Text: "--no-trunc", Description: `Do not truncate output`},
}

var flagStop = []prompt.Suggest{
	{Text: "-t", Description: `Seconds to wait for stop before killing it (default 10)`},
	{Text: "--time", Description: `Seconds to wait for stop before killing it (default 10)`},
}

var flagTag = []prompt.Suggest{}

var flagTop = []prompt.Suggest{}

var flagUnpause = []prompt.Suggest{}

var flagUpdate = []prompt.Suggest{
	{Text: "--blkio-weight", Description: `Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)`},
	{Text: "--cpu-period", Description: `Limit CPU CFS (Completely Fair Scheduler) period`},
	{Text: "--cpu-quota", Description: `Limit CPU CFS (Completely Fair Scheduler) quota`},
	{Text: "--cpu-rt-period", Description: `Limit the CPU real-time period in microseconds`},
	{Text: "--cpu-rt-runtime", Description: `Limit the CPU real-time runtime in microseconds`},
	{Text: "--cpu-shares", Description: `CPU shares (relative weight)`},
	{Text: "--cpus decimal", Description: `Number of CPUs`},
	{Text: "--cpuset-cpus", Description: `CPUs in which to allow execution (0-3, 0,1)`},
	{Text: "--cpuset-mems", Description: `MEMs in which to allow execution (0-3, 0,1)`},
	{Text: "--kernel-memory", Description: `Kernel memory limit`},
	{Text: "--memory", Description: `Memory limit`},
	{Text: "--memory-reservation", Description: `Memory soft limit`},
	{Text: "--memory-swap", Description: `Swap limit equal to memory plus swap: '-1' to enable unlimited swap`},
	{Text: "--restart", Description: `Restart policy to apply when a container exits`},
}

var flagVersion = []prompt.Suggest{
	{Text: "-f", Description: `Format the output using the given Go template`},
	{Text: "--format", Description: `Format the output using the given Go template`},
}

var flagWait = []prompt.Suggest{}
