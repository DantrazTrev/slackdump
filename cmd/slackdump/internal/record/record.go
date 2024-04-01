package record

import (
	"context"
	_ "embed"
	"errors"
	"strings"
	"time"

	"github.com/rusq/fsadapter"
	"github.com/rusq/slackdump/v3"
	"github.com/rusq/slackdump/v3/cmd/slackdump/internal/cfg"
	"github.com/rusq/slackdump/v3/cmd/slackdump/internal/golang/base"
	"github.com/rusq/slackdump/v3/internal/chunk"
	"github.com/rusq/slackdump/v3/internal/chunk/control"
	"github.com/rusq/slackdump/v3/internal/chunk/transform/fileproc"
	"github.com/rusq/slackdump/v3/internal/structures"
	"github.com/rusq/slackdump/v3/logger"
)

//go:embed assets/record.md
var mdRecord string

var CmdRecord = &base.Command{
	Run:         RunRecord,
	UsageLine:   "slackdump record [flags] [link1[ link 2[ link N]]]",
	Short:       "record the dump of the workspace or individual conversations",
	Long:        mdRecord,
	FlagMask:    cfg.OmitUserCacheFlag | cfg.OmitCacheDir,
	RequireAuth: true,
	PrintFlags:  true,
}

const zipExt = ".ZIP"

func StripZipExt(s string) string {
	if strings.HasSuffix(strings.ToUpper(s), zipExt) {
		return s[:len(s)-len(zipExt)]
	}
	return s
}

var (
	errNoOutput = errors.New("output directory is required")
)

func RunRecord(ctx context.Context, cmd *base.Command, args []string) error {
	list, err := structures.NewEntityList(args)
	if err != nil {
		base.SetExitStatus(base.SUserError)
		return err
	}

	cfg.Output = StripZipExt(cfg.Output)
	if cfg.Output == "" {
		base.SetExitStatus(base.SInvalidParameters)
		return errNoOutput
	}

	cd, err := chunk.CreateDir(cfg.Output)
	if err != nil {
		base.SetExitStatus(base.SGenericError)
		return err
	}

	sess, err := cfg.SlackdumpSession(ctx)
	if err != nil {
		base.SetExitStatus(base.SInitializationError)
		return err
	}
	lg := logger.FromContext(ctx)
	stream := sess.Stream(
		slackdump.OptLatest(time.Time(cfg.Latest)),
		slackdump.OptOldest(time.Time(cfg.Oldest)),
		slackdump.OptResultFn(resultLogger(lg)),
	)
	dl, stop := fileproc.NewDownloader(
		ctx,
		cfg.DownloadFiles,
		sess.Client(),
		fsadapter.NewDirectory(cd.Name()),
		lg,
	)
	defer stop()
	// we are using the same file subprocessor as the mattermost export.
	subproc := fileproc.NewExport(fileproc.STmattermost, dl)
	ctrl := control.New(cd, stream, control.WithLogger(lg), control.WithFiler(subproc))
	if err := ctrl.Run(ctx, list); err != nil {
		base.SetExitStatus(base.SApplicationError)
		return err
	}
	lg.Printf("Recorded workspace data to %s", cd.Name())

	return nil
}

func resultLogger(lg logger.Interface) func(sr slackdump.StreamResult) error {
	return func(sr slackdump.StreamResult) error {
		lg.Printf("%s", sr)
		return nil
	}
}
