package project

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd/mutagen/forward"
	"github.com/mutagen-io/mutagen/cmd/mutagen/sync"
	"github.com/mutagen-io/mutagen/pkg/filesystem/locking"
	"github.com/mutagen-io/mutagen/pkg/identifier"
	"github.com/mutagen-io/mutagen/pkg/project"
)

func listMain(command *cobra.Command, arguments []string) error {
	// Compute the name of the configuration file and change our working
	// directory to the path in which the file resides.
	var configurationFileName string
	if len(arguments) == 0 {
		configurationFileName = project.DefaultConfigurationFileName
	} else if len(arguments) == 1 {
		// Parse the target into directory and file name.
		var directory string
		directory, configurationFileName = filepath.Split(arguments[0])
		if configurationFileName == "" {
			return errors.New("empty configuration file name")
		}

		// Switch to the directory (if it's not the current directory).
		if directory != "" {
			if err := os.Chdir(directory); err != nil {
				return errors.Wrap(err, "unable to switch to target directory")
			}
		}
	} else {
		return errors.New("invalid number of arguments")
	}

	// Compute the lock path.
	lockPath := configurationFileName + project.LockFileExtension

	// Track whether or not we should remove the lock file on return.
	var removeLockFileOnReturn bool

	// Create a locker and defer its closure and potential removal. On Windows
	// systems, we have to handle this removal after the file is closed.
	locker, err := locking.NewLocker(lockPath, 0600)
	if err != nil {
		return errors.Wrap(err, "unable to create project locker")
	}
	defer func() {
		locker.Close()
		if removeLockFileOnReturn && runtime.GOOS == "windows" {
			os.Remove(lockPath)
		}
	}()

	// Acquire the project lock and defer its release and potential removal. On
	// Windows systems, we can't remove the lock file if it's locked or even
	// just opened, so we handle removal for Windows systems after we close the
	// lock file (see above). In this case, we truncate the lock file before
	// releasing it to ensure that any other process that opens or acquires the
	// lock file before we manage to remove it will simply see an empty lock
	// file, which it will ignore or attempt to remove.
	if err := locker.Lock(true); err != nil {
		return errors.Wrap(err, "unable to acquire project lock")
	}
	defer func() {
		if removeLockFileOnReturn {
			if runtime.GOOS == "windows" {
				locker.Truncate(0)
			} else {
				os.Remove(lockPath)
			}
		}
		locker.Unlock()
	}()

	// Read the project identifier from the lock file. If the lock file is
	// empty, then we can assume that we created it when we created the lock and
	// just remove it.
	buffer := &bytes.Buffer{}
	if length, err := buffer.ReadFrom(locker); err != nil {
		return errors.Wrap(err, "unable to read project lock")
	} else if length == 0 {
		removeLockFileOnReturn = true
		return errors.New("project not running")
	}
	projectIdentifier := buffer.String()

	// Ensure that the project identifier is valid.
	if !identifier.IsValid(projectIdentifier) {
		return errors.New("invalid project identifier found in project lock")
	}

	// Compute the label selector that we're going to use to list sessions.
	labelSelector := fmt.Sprintf("%s=%s", project.LabelKey, projectIdentifier)

	// List forwarding sessions.
	fmt.Println("Forwarding sessions:")
	if err := forward.ListWithLabelSelector(labelSelector, listConfiguration.long); err != nil {
		return errors.Wrap(err, "unable to list forwarding session(s)")
	}

	// Print an empty line.
	fmt.Println()

	// List synchronization sessions.
	fmt.Println("Synchronization sessions:")
	if err := sync.ListWithLabelSelector(labelSelector, listConfiguration.long); err != nil {
		return errors.Wrap(err, "unable to list synchronization session(s)")
	}

	// Success.
	return nil
}

var listCommand = &cobra.Command{
	Use:          "list [<configuration-file>]",
	Short:        "List project sessions",
	RunE:         listMain,
	SilenceUsage: true,
}

var listConfiguration struct {
	// help indicates whether or not help information should be shown for the
	// command.
	help bool
	// long indicates whether or not to use long-format listing.
	long bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := listCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&listConfiguration.help, "help", "h", false, "Show help information")

	// Wire up list flags.
	flags.BoolVarP(&listConfiguration.long, "long", "l", false, "Show detailed session information")
}
