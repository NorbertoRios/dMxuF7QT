package repository

import (
	"genx-go/core/device/interfaces"
	"genx-go/message"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//NewDeviceStateRepository returns new mysql repository
func NewDeviceStateRepository(_connectionString string) *DeviceStateRepository {
	_connection, err := gorm.Open(mysql.Open(_connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Error connecting to raw database:" + err.Error())
	}
	return &DeviceStateRepository{
		Connection: _connection,
	}
}

//DeviceStateRepository represents mysql repository
type DeviceStateRepository struct {
	Connection *gorm.DB
}

//Save ...
func (db *DeviceStateRepository) Save(device ...interfaces.IDevice) error {
	return nil
}

//Load returns device last message
func (db *DeviceStateRepository) Load(identity string) interface{} {
	return nil
}

func (db *DeviceStateRepository) save(message *message.Message) (uint64, error) {
	return uint64(0), nil
}

func (db *DeviceStateRepository) createMessageHistoryTable(tableName string) error {
	return db.Connection.Exec("CREATE TABLE IF NOT EXISTS  raw_data.`" + tableName + "` ( " +
		"`Id` bigint(20) NOT NULL AUTO_INCREMENT, " +
		"`DevId` varchar(100) NOT NULL, " +
		"`EntryData` blob, " +
		"`ParsedEntryData` blob, " +
		"`Time` datetime NOT NULL, " +
		"`RecievedTime` datetime NOT NULL, " +
		"`ReportClass` varchar(100) DEFAULT NULL, " +
		"`ReportType` int(11) DEFAULT NULL, " +
		"`Reason` varchar(5) DEFAULT NULL, " +
		"`Latitude` double DEFAULT NULL COMMENT 'degrees', " +
		"`Longitude` double DEFAULT NULL COMMENT 'degrees', " +
		"`Speed` double DEFAULT NULL, " +
		"`ValidFix` int(11) DEFAULT NULL, " +
		"`Altitude` double DEFAULT NULL, " +
		"`Heading` double DEFAULT NULL, " +
		"`IgnitionState` int(11) DEFAULT NULL, " +
		"`Odometer` int(10) DEFAULT NULL COMMENT 'm', " +
		"`Satellites` tinyint(3) unsigned DEFAULT NULL, " +
		"`Supply` int(10) DEFAULT NULL, " +
		"`GPIO` int(10) DEFAULT NULL COMMENT 'Input ports state', " +
		"`Relay` int(10) DEFAULT NULL COMMENT 'Output ports state', " +
		"`msg_id` binary(16) DEFAULT NULL, " +
		"`Extra` text, " +
		"`BatteryLow` double DEFAULT NULL, " +
		" PRIMARY KEY (`Id`,`Time`,`DevId`), " +
		"KEY `IX_RecievedTime` (`RecievedTime`,`DevId`) " +
		")" +
		"ENGINE = INNODB " +
		"AVG_ROW_LENGTH = 8192 " +
		"CHARACTER SET utf8 " +
		"COLLATE utf8_general_ci " +
		"PARTITION BY RANGE (to_days(Time)) " +
		"(" +
		"PARTITION p180201 VALUES LESS THAN (737091) ENGINE = InnoDB, " +
		"PARTITION p180301 VALUES LESS THAN (737119) ENGINE = InnoDB, " +
		"PARTITION p180401 VALUES LESS THAN (737150) ENGINE = InnoDB, " +
		"PARTITION p180501 VALUES LESS THAN (737180) ENGINE = InnoDB, " +
		"PARTITION p180601 VALUES LESS THAN (737211) ENGINE = InnoDB, " +
		"PARTITION p180701 VALUES LESS THAN (737241) ENGINE = InnoDB, " +
		"PARTITION p180801 VALUES LESS THAN (737272) ENGINE = InnoDB, " +
		"PARTITION p180901 VALUES LESS THAN (737303) ENGINE = InnoDB, " +
		"PARTITION p181001 VALUES LESS THAN (737333) ENGINE = InnoDB, " +
		"PARTITION p181101 VALUES LESS THAN (737364) ENGINE = InnoDB, " +
		"PARTITION p181201 VALUES LESS THAN (737394) ENGINE = InnoDB, " +
		"PARTITION p190101 VALUES LESS THAN (737425) ENGINE = InnoDB, " +
		"PARTITION p190201 VALUES LESS THAN (737456) ENGINE = InnoDB, " +
		"PARTITION p_cur VALUES LESS THAN MAXVALUE ENGINE = InnoDB " +
		");").Error
}
