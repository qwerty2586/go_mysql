package imSQL

type (
	Privileges struct {
		SelectPriv           string `json:"Select_priv" db:"Select_priv"`
		InsertPriv           string `json:"Insert_priv" db:"Insert_priv"`
		UpdatePriv           string `json:"Update_priv" db:"Update_priv"`
		DeletePriv           string `json:"Delete_priv" db:"Delete_priv"`
		CreatePriv           string `json:"Create_priv" db:"Create_priv"`
		DropPriv             string `json:"Drop_priv" db:"Drop_priv"`
		ReloadPriv           string `json:"Reload_priv" db:"Reload_priv"`
		ShutdownPriv         string `json:"Shutdown_priv" db:"Shutdown_priv"`
		ProcessPriv          string `json:"Process_priv" db:"Process_priv"`
		FilePriv             string `json:"File_priv" db:"File_priv"`
		GrantPriv            string `json:"Grant_priv" db:"Grant_priv"`
		ReferencesPriv       string `json:"References_priv" db:"References_priv"`
		IndexPriv            string `json:"Index_priv" db:"Index_priv"`
		AlterPriv            string `json:"Alter_priv" db:"Alter_priv"`
		ShowDbPriv           string `json:"Show_db_priv" db:"Show_db_priv"`
		SuperPriv            string `json:"Super_priv" db:"Super_priv"`
		CreateTmpTablePriv   string `json:"Create_tmp_table_priv" db:"Create_tmp_table_priv"`
		LockTablesPriv       string `json:"Lock_tables_priv" db:"Lock_tables_priv"`
		ExecutePriv          string `json:"Execute_priv" db:"Execute_priv"`
		ReplSlavePriv        string `json:"Repl_slave_priv" db:"Repl_slave_priv"`
		ReplClientPriv       string `json:"Repl_client_priv" db:"Repl_client_priv"`
		CreateViewPriv       string `json:"Create_view_priv" db:"Create_view_priv"`
		ShowViewPriv         string `json:"Show_view_priv" db:"Show_view_priv"`
		CreateRoutinePriv    string `json:"Create_routine_priv" db:"Create_routine_priv"`
		AlterRoutinePriv     string `json:"Alter_routine_priv" db:"Alter_routine_priv"`
		CreateUserPriv       string `json:"Create_user_priv" db:"Create_user_priv"`
		EventPriv            string `json:"Event_priv" db:"Event_priv"`
		TriggerPriv          string `json:"Trigger_priv" db:"Trigger_priv"`
		CreateTablespacePriv string `json:"Create_tablespace_priv" db:"Create_tablespace_priv"`
	}
)
