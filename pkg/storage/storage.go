package storage

import (
	"fmt"
	"net/url"

	"go.uber.org/zap"
)

type ExternalStorage interface {
	SetBackupName(name string)
	BackupPreCommand() []string
	BackupStorageCommand(src string, host string, spaceID string) string
	BackupMetaCommand(src []string) string
	BackupMetaFileCommand(src string) []string
	RestoreMetaFileCommand(file string, dst string) []string
	RestoreMetaCommand(src []string, dst string) (string, []string)
	RestoreStorageCommand(host string, spaceID []string, dst string) string
	RestoreMetaPreCommand(dst string) string
	RestoreStoragePreCommand(dst string) string
	CheckCommand() string
	ListBackupCommand() ([]string, error)
	URI() string
}

func NewExternalStorage(storageUrl string, log *zap.Logger, maxConcurrent int, args string) (ExternalStorage, error) {
	u, err := url.Parse(storageUrl)
	if err != nil {
		return nil, err
	}

	log.Info("parsed external storage", zap.String("schema", u.Scheme), zap.String("path", u.Path))

	switch u.Scheme {
	case "local":
		return NewLocalBackedStore(u.Path, log, maxConcurrent, args), nil
	case "s3":
		return NewS3BackendStore(storageUrl, log, maxConcurrent, args), nil
	case "oss":
		return NewOSSBackendStore(storageUrl, log, maxConcurrent, args), nil
	case "hdfs":
		return NewHDFSBackendStore(storageUrl, log, maxConcurrent, args), nil
	default:
		return nil, fmt.Errorf("Unsupported Backend Storage Types")
	}
}
