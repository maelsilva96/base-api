package utils

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/loggingingestion"
	"time"
)

type LogConsole struct {
}

type OracleLogWriter struct {
	LogId           string
	ApplicationName string
	loggingClient   loggingingestion.LoggingClient
}

func (lw *OracleLogWriter) Write(data []byte) (n int, err error) {
	go func() {
		err = lw.SendLog(string(data))
		if err != nil {
			panic(err)
		}
	}()
	return len(data), err
}

func (lw *LogConsole) Write(data []byte) (n int, err error) {
	fmt.Println(string(data))
	return len(data), nil
}

func (lw *OracleLogWriter) SendLog(data string) error {
	_, err := lw.loggingClient.PutLogs(context.Background(), loggingingestion.PutLogsRequest{
		LogId: common.String(lw.LogId),
		PutLogsDetails: loggingingestion.PutLogsDetails{
			LogEntryBatches: []loggingingestion.LogEntryBatch{
				{
					Type:   common.String(fmt.Sprintf("%s.logRequest", lw.ApplicationName)),
					Source: common.String(lw.ApplicationName),
					Entries: []loggingingestion.LogEntry{
						{
							Id:   common.String(uuid.NewString()),
							Time: &common.SDKTime{Time: time.Now()},
							Data: common.String(data),
						},
					},
					Defaultlogentrytime: &common.SDKTime{Time: time.Now()},
				},
			},
			Specversion: common.String("1.0"),
		},
		TimestampOpcAgentProcessing: &common.SDKTime{Time: time.Now()},
	})
	return err
}

func NewOracleLogWriter(applicationName string, logId string, ociConfigProvider *common.ConfigurationProvider) *OracleLogWriter {
	lw := &OracleLogWriter{
		ApplicationName: applicationName,
		LogId:           logId,
	}
	var err error
	lw.loggingClient, err = loggingingestion.NewLoggingClientWithConfigurationProvider(*ociConfigProvider)
	if err != nil {
		panic(err)
	}
	return lw
}
