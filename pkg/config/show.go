package config

import (
	"fmt"

	"github.com/spf13/pflag"
	pb "github.com/vesoft-inc/nebula-agent/pkg/proto"
	"github.com/vesoft-inc/nebula-br/pkg/storage"
)

type ShowConfig struct {
	Backend *pb.Backend
}

func (s *ShowConfig) ParseFlags(flags *pflag.FlagSet) error {
	var err error
	s.Backend, err = storage.ParseFromFlags(flags)
	if err != nil {
		return fmt.Errorf("parse storage flags failed: %w", err)
	}
	return nil
}
