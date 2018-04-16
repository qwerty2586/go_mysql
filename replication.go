package imSQL

type (
	// show master status
	Master struct {
		File           string `json:"file" db:"file"`
		Position       uint64 `json:"position" db:"position"`
		BinlogDoDb     string `json:"binlog_do_db" db:"binlog_do_db"`
		BinlogIgnoreDb string `json:"binlog_ignore_db" db:"binlog_ignore_db"`
		ExecutedGTID   string `json:"executed_gtid_set" db:"executed_gtid_set"`
	}
	// change master to
	Slave struct {
		MasterBind                string   `json:"master_bind" db:"master_bind"`
		MasterHost                string   `json:"master_host" db:"master_host"`
		MasterUser                string   `json:"master_user" db:"master_user"`
		MasterPassword            string   `json:"master_password" db:"master_password"`
		MasterPort                uint64   `json:"master_port" db:"master_port"`
		MasterConnectRetry        uint64   `json:"master_connect_retry" db:"master_connect_retry"`
		MasterRetryCount          uint64   `json:"master_retry_count" db:"master_retry_count"`
		MasterDelay               uint64   `json:"master_delay" db"master_delay"`
		MasterHeartbeatPeriod     uint64   `json:"master_heartbeat_period" db:"master_heartbeat_period"`
		MasterLogFile             string   `json:"master_log_file" db:"master_log_file"`
		MasterLogPos              uint64   `json:"master_log_pos" db:"master_log_pos"`
		MasterAutoPosition        uint64   `json:"master_auto_position" db:"master_auto_position"`
		RelayLogFile              string   `json:"relay_log_file" db:"relay_log_file"`
		RelayLogPos               uint64   `json:"relay_log_pos" db:"relay_log_pos"`
		MasterSSL                 uint64   `json:"master_ssl" db:"master_ssl"`
		MasterSSLCa               string   `json:"master_ssl_ca" db:"master_ssl_ca"`
		MasterSSLCapath           string   `json:"master_ssl_capath" db:"master_ssl_capath"`
		MasterSSLCert             string   `json:"master_ssl_cert" db:"master_ssl_cert"`
		MasterSSLCrl              string   `json:"master_ssl_crl" db:"master_ssl_crl"`
		MasterSSLCrlpath          string   `json:"master_ssl_crlpath" db:"master_ssl_crlpath"`
		MasterSSLKey              string   `json:"master_ssl_key" db:"master_ssl_key"`
		MasterSSLCipher           string   `json:"master_ssl_cipher" db:"master_ssl_cipher"`
		MasterSSLVerifyServerCert uint64   `json:"master_ssl_verify_server_cert" db:"master_ssl_verify_server_cert"`
		MasterTLSVersion          string   `json:"master_tls_version" db:"master_tls_version"`
		IgnoreServerIds           []string `json:"ignore_server_ids" db:"ignore_server_ids"`
	}
)

// new slave struct.
func NewSlave(host string, port uint64, username string, password string, filename string, filepos uint64) (*Slave, error) {

	slave := new(Slave)

	slave.MasterHost = host
	slave.MasterPort = port
	slave.MasterUser = username
	slave.MasterPassword = password
	slave.MasterLogFile = filename
	slave.MasterLogPos = filepos

	slave.MasterBind = ""
	slave.MasterConnectRetry = 60
	slave.MasterRetryCount = 86400
	slave.MasterDelay = 0
	slave.MasterHeartbeatPeriod = 0
	slave.MasterAutoPosition = 0

	slave.RelayLogFile = ""
	slave.RelayLogPos = 0

	slave.MasterSSL = 0
	slave.MasterSSLCa = ""
	slave.MasterSSLCapath = ""
	slave.MasterSSLCapath = ""
	slave.MasterSSLCert = ""
	slave.MasterSSLCrl = ""
	slave.MasterSSLCrlpath = ""
	slave.MasterSSLKey = ""
	slave.MasterSSLCipher = ""
	slave.MasterSSLVerifyServerCert = 0
	slave.MasterTLSVersion = ""
	slave.IgnoreServerIds = []string{}

	return slave, nil
}
