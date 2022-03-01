package mysqlmanage

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type (
	Variable struct {
		Name    string `json:"Variable_name" db:"Variable_name"`
		Value   string `json:"Value" db:"Value"`
		Type    string `json:"Type" db:"Type"`
		Dynamic string `json:"Dynamic" db:"Dynamic"`
	}
)

const (
	StmtShowVariables              = `SHOW VARIABLES`
	StmtSetIntegerDynamicVariables = `SET GLOBAL %s = %s`
	StmtSetStringDynamicVariables  = `SET GLOBAL %s = '%s'`
)

var (
	// global or session variables.
	// https://dev.mysql.com/doc/refman/5.7/en/dynamic-system-variables.html
	GlobalDynamicVars = map[string]string{
		"audit_log_connection_policy":                        "enumeration",
		"audit_log_exclude_accounts":                         "string",
		"audit_log_flush":                                    "boolean",
		"audit_log_include_accounts":                         "string",
		"audit_log_rotate_on_size":                           "integer",
		"audit_log_statement_policy":                         "enumeration",
		"authentication_ldap_sasl_auth_method_name":          "string",
		"authentication_ldap_sasl_bind_base_dn":              "string",
		"authentication_ldap_sasl_bind_root_dn":              "string",
		"authentication_ldap_sasl_bind_root_pwd":             "string",
		"authentication_ldap_sasl_ca_path":                   "string",
		"authentication_ldap_sasl_group_search_attr":         "string",
		"authentication_ldap_sasl_group_search_filter":       "string",
		"authentication_ldap_sasl_init_pool_size":            "integer",
		"authentication_ldap_sasl_log_status":                "integer",
		"authentication_ldap_sasl_max_pool_size":             "integer",
		"authentication_ldap_sasl_server_host":               "string",
		"authentication_ldap_sasl_server_port":               "integer",
		"authentication_ldap_sasl_tls":                       "boolean",
		"authentication_ldap_sasl_user_search_attr":          "string",
		"authentication_ldap_simple_auth_method_name":        "string",
		"authentication_ldap_simple_bind_base_dn":            "string",
		"authentication_ldap_simple_bind_root_dn":            "string",
		"authentication_ldap_simple_bind_root_pwd":           "string",
		"authentication_ldap_simple_ca_path":                 "string",
		"authentication_ldap_simple_group_search_attr":       "string",
		"authentication_ldap_simple_group_search_filter":     "string",
		"authentication_ldap_simple_init_pool_size":          "integer",
		"authentication_ldap_simple_log_status":              "integer",
		"authentication_ldap_simple_max_pool_size":           "integer",
		"authentication_ldap_simple_server_host":             "string",
		"authentication_ldap_simple_server_port":             "integer",
		"authentication_ldap_simple_tls":                     "boolean",
		"authentication_ldap_simple_user_search_attr":        "string",
		"auto_increment_increment":                           "integer",
		"auto_increment_offset":                              "integer",
		"autocommit":                                         "boolean",
		"automatic_sp_privileges":                            "boolean",
		"avoid_temporal_upgrade":                             "boolean",
		"big_tables":                                         "boolean",
		"binlog_cache_size":                                  "integer",
		"binlog_checksum":                                    "string",
		"binlog_direct_non_transactional_updates":            "boolean",
		"binlog_error_action":                                "enumeration",
		"binlog_format":                                      "enumeration",
		"binlog_group_commit_sync_delay":                     "integer",
		"binlog_group_commit_sync_no_delay_count":            "integer",
		"binlog_max_flush_queue_time":                        "integer",
		"binlog_order_commits":                               "boolean",
		"binlog_row_image=image_type":                        "enumeration",
		"binlog_rows_query_log_events":                       "boolean",
		"binlog_stmt_cache_size":                             "integer",
		"binlogging_impossible_mode":                         "enumeration",
		"block_encryption_mode":                              "string",
		"bulk_insert_buffer_size":                            "integer",
		"character_set_client":                               "string",
		"character_set_connection":                           "string",
		"character_set_database":                             "string",
		"character_set_filesystem":                           "string",
		"character_set_results":                              "string",
		"character_set_server":                               "string",
		"check_proxy_users":                                  "boolean",
		"collation_connection":                               "string",
		"collation_database":                                 "string",
		"collation_server":                                   "string",
		"completion_type":                                    "enumeration",
		"concurrent_insert":                                  "enumeration",
		"connect_timeout":                                    "integer",
		"connection_control_failed_connections_threshold":    "integer",
		"connection_control_max_connection_delay":            "integer",
		"connection_control_min_connection_delay":            "integer",
		"default_password_lifetime":                          "integer",
		"default_storage_engine":                             "enumeration",
		"default_tmp_storage_engine":                         "enumeration",
		"default_week_format":                                "integer",
		"delay_key_write":                                    "enumeration",
		"delayed_insert_limit":                               "integer",
		"delayed_insert_timeout":                             "integer",
		"delayed_queue_size":                                 "integer",
		"div_precision_increment":                            "integer",
		"end_markers_in_json":                                "boolean",
		"enforce_gtid_consistency":                           "enumeration",
		"eq_range_index_dive_limit":                          "integer",
		"event_scheduler":                                    "enumeration",
		"executed_gtids_compression_period":                  "integer",
		"expire_logs_days":                                   "integer",
		"explicit_defaults_for_timestamp":                    "boolean",
		"flush":                                              "boolean",
		"flush_time":                                         "integer",
		"foreign_key_checks":                                 "boolean",
		"ft_boolean_syntax":                                  "string",
		"general_log":                                        "boolean",
		"general_log_file":                                   "filename",
		"group_concat_max_len":                               "integer",
		"group_replication_allow_local_disjoint_gtids_join":  "boolean",
		"group_replication_allow_local_lower_version_join":   "boolean",
		"group_replication_auto_increment_increment":         "integer",
		"group_replication_bootstrap_group":                  "boolean",
		"group_replication_components_stop_timeout":          "integer",
		"group_replication_compression_threshold":            "integer",
		"group_replication_enforce_update_everywhere_checks": "boolean",
		"group_replication_flow_control_applier_threshold":   "integer",
		"group_replication_flow_control_certifier_threshold": "integer",
		"group_replication_flow_control_mode":                "enumeration",
		"group_replication_force_members":                    "string",
		"group_replication_group_name":                       "string",
		"group_replication_group_seeds":                      "string",
		"group_replication_gtid_assignment_block_size":       "integer",
		"group_replication_ip_whitelist":                     "string",
		"group_replication_local_address":                    "string",
		"group_replication_member_weight":                    "integer",
		"group_replication_poll_spin_loops":                  "integer",
		"group_replication_recovery_complete_at":             "enumeration",
		"group_replication_recovery_reconnect_interval":      "integer",
		"group_replication_recovery_retry_count":             "integer",
		"group_replication_recovery_ssl_ca":                  "string",
		"group_replication_recovery_ssl_capath":              "string",
		"group_replication_recovery_ssl_cert":                "string",
		"group_replication_recovery_ssl_cipher":              "string",
		"group_replication_recovery_ssl_crl":                 "string",
		"group_replication_recovery_ssl_crlpath":             "string",
		"group_replication_recovery_ssl_key":                 "string",
		"group_replication_recovery_ssl_verify_server_cert":  "boolean",
		"group_replication_recovery_use_ssl":                 "boolean",
		"group_replication_single_primary_mode":              "boolean",
		"group_replication_ssl_mode":                         "enumeration",
		"group_replication_start_on_boot":                    "boolean",
		"group_replication_transaction_size_limit":           "integer",
		"group_replication_unreachable_majority_timeout":     "integer",
		"gtid_executed_compression_period":                   "integer",
		"gtid_mode":                                          "enumeration",
		"gtid_purged":                                        "string",
		"host_cache_size":                                    "integer",
		"init_connect":                                       "string",
		"init_slave":                                         "string",
		"innodb_adaptive_flushing":                           "boolean",
		"innodb_adaptive_flushing_lwm":                       "integer",
		"innodb_adaptive_hash_index":                         "boolean",
		"innodb_adaptive_max_sleep_delay":                    "integer",
		"innodb_api_bk_commit_interval":                      "integer",
		"innodb_api_trx_level":                               "integer",
		"innodb_autoextend_increment":                        "integer",
		"innodb_background_drop_list_empty":                  "boolean",
		"innodb_buffer_pool_dump_at_shutdown":                "boolean",
		"innodb_buffer_pool_dump_now":                        "boolean",
		"innodb_buffer_pool_dump_pct":                        "integer",
		"innodb_buffer_pool_filename":                        "filename",
		"innodb_buffer_pool_load_abort":                      "boolean",
		"innodb_buffer_pool_load_now":                        "boolean",
		"innodb_buffer_pool_size":                            "integer",
		"innodb_change_buffer_max_size":                      "integer",
		"innodb_change_buffering":                            "enumeration",
		"innodb_change_buffering_debug":                      "integer",
		"innodb_checksum_algorithm":                          "enumeration",
		"innodb_cmp_per_index_enabled":                       "boolean",
		"innodb_commit_concurrency":                          "integer",
		"innodb_compress_debug":                              "enumeration",
		"innodb_compression_failure_threshold_pct":           "integer",
		"innodb_compression_level":                           "integer",
		"innodb_compression_pad_pct_max":                     "integer",
		"innodb_concurrency_tickets":                         "integer",
		"innodb_deadlock_detect":                             "boolean",
		"innodb_default_row_format":                          "enumeration",
		"innodb_disable_resize_buffer_pool_debug":            "boolean",
		"innodb_disable_sort_file_cache":                     "boolean",
		"innodb_fast_shutdown":                               "integer",
		"innodb_fil_make_page_dirty_debug":                   "integer",
		"innodb_file_format":                                 "string string",
		"innodb_file_format_max":                             "string",
		"innodb_file_per_table":                              "boolean",
		"innodb_fill_factor":                                 "integer",
		"innodb_flush_log_at_timeout":                        "integer",
		"innodb_flush_log_at_trx_commit":                     "enumeration",
		"innodb_flush_neighbors":                             "enumeration",
		"innodb_flush_sync":                                  "boolean",
		"innodb_flushing_avg_loops":                          "integer",
		"innodb_ft_aux_table":                                "string",
		"innodb_ft_enable_diag_print":                        "boolean",
		"innodb_ft_enable_stopword":                          "boolean",
		"innodb_ft_num_word_optimize":                        "integer",
		"innodb_ft_result_cache_limit":                       "integer",
		"innodb_ft_server_stopword_table":                    "string",
		"innodb_ft_user_stopword_table":                      "string",
		"innodb_io_capacity":                                 "integer",
		"innodb_io_capacity_max":                             "integer",
		"innodb_large_prefix":                                "boolean",
		"innodb_limit_optimistic_insert_debug":               "integer",
		"innodb_lock_wait_timeout":                           "integer",
		"innodb_log_checksum_algorithm":                      "enumeration",
		"innodb_log_checksums":                               "boolean",
		"innodb_log_compressed_pages":                        "boolean",
		"innodb_log_write_ahead_size":                        "integer",
		"innodb_lru_scan_depth":                              "integer",
		"innodb_max_dirty_pages_pct":                         "numeric",
		"innodb_max_dirty_pages_pct_lwm":                     "numeric",
		"innodb_max_purge_lag":                               "integer",
		"innodb_max_purge_lag_delay":                         "integer",
		"innodb_max_undo_log_size":                           "integer",
		"innodb_merge_threshold_set_all_debug":               "integer",
		"innodb_monitor_disable":                             "string",
		"innodb_monitor_enable":                              "string",
		"innodb_monitor_reset":                               "string",
		"innodb_monitor_reset_all":                           "string",
		"innodb_old_blocks_pct":                              "integer",
		"innodb_old_blocks_time":                             "integer",
		"innodb_online_alter_log_max_size":                   "integer",
		"innodb_optimize_fulltext_only":                      "boolean",
		"innodb_print_all_deadlocks":                         "boolean",
		"innodb_purge_batch_size":                            "integer",
		"innodb_purge_rseg_truncate_frequency":               "integer",
		"innodb_random_read_ahead":                           "boolean",
		"innodb_read_ahead_threshold":                        "integer",
		"innodb_replication_delay":                           "integer",
		"innodb_rollback_segments":                           "integer",
		"innodb_saved_page_number_debug":                     "integer",
		"innodb_spin_wait_delay":                             "integer",
		"innodb_stats_auto_recalc":                           "boolean",
		"innodb_stats_include_delete_marked":                 "boolean",
		"innodb_stats_method":                                "enumeration",
		"innodb_stats_on_metadata":                           "boolean",
		"innodb_stats_persistent":                            "boolean",
		"innodb_stats_persistent_sample_pages":               "integer",
		"innodb_stats_sample_pages":                          "integer",
		"innodb_stats_transient_sample_pages":                "integer",
		"innodb_status_output":                               "boolean",
		"innodb_status_output_locks":                         "boolean",
		"innodb_strict_mode":                                 "boolean",
		"innodb_support_xa":                                  "boolean",
		"innodb_sync_spin_loops":                             "integer",
		"innodb_table_locks":                                 "boolean",
		"innodb_thread_concurrency":                          "integer",
		"innodb_thread_sleep_delay":                          "integer",
		"innodb_tmpdir":                                      "dirname",
		"innodb_trx_purge_view_update_only_debug":            "boolean",
		"innodb_trx_rseg_n_slots_debug":                      "integer",
		"innodb_undo_log_truncate":                           "boolean",
		"innodb_undo_logs":                                   "integer",
		"interactive_timeout":                                "integer",
		"internal_tmp_disk_storage_engine":                   "enumeration",
		"join_buffer_size":                                   "integer",
		"keep_files_on_create":                               "boolean",
		"key_buffer_size":                                    "integer",
		"key_cache_age_threshold":                            "integer",
		"key_cache_block_size":                               "integer",
		"key_cache_division_limit":                           "integer",
		"keyring_aws_cmk_id":                                 "string",
		"keyring_aws_region":                                 "enumeration",
		"keyring_encrypted_file_data":                        "filename",
		"keyring_encrypted_file_password":                    "string",
		"keyring_file_data":                                  "filename",
		"keyring_okv_conf_dir":                               "dirname",
		"keyring_operations":                                 "boolean",
		"lc_messages":                                        "string",
		"lc_time_names":                                      "string",
		"local_infile":                                       "boolean",
		"lock_wait_timeout":                                  "integer",
		"log_backward_compatible_user_definitions":           "boolean",
		"log_bin_trust_function_creators":                    "boolean",
		"log_builtin_as_identified_by_password":              "boolean",
		"log_error_verbosity":                                "integer",
		"log_output":                                         "set",
		"log_queries_not_using_indexes":                      "boolean",
		"log_slow_admin_statements":                          "boolean",
		"log_slow_slave_statements":                          "boolean",
		"log_statements_unsafe_for_binlog":                   "boolean",
		"log_syslog":                                         "boolean",
		"log_syslog_facility":                                "string",
		"log_syslog_include_pid":                             "boolean",
		"log_syslog_tag":                                     "string",
		"log_throttle_queries_not_using_indexes":             "integer",
		"log_timestamps":                                     "enumeration",
		"log_warnings":                                       "integer",
		"long_query_time":                                    "numeric",
		"low_priority_updates":                               "boolean",
		"master_info_repository":                             "string",
		"master_verify_checksum":                             "boolean",
		"max_allowed_packet":                                 "integer",
		"max_binlog_cache_size":                              "integer",
		"max_binlog_size":                                    "integer",
		"max_binlog_stmt_cache_size":                         "integer",
		"max_connect_errors":                                 "integer",
		"max_connections":                                    "integer",
		"max_delayed_threads":                                "integer",
		"max_error_count":                                    "integer",
		"max_execution_time":                                 "integer",
		"max_heap_table_size":                                "integer",
		"max_insert_delayed_threads":                         "integer",
		"max_join_size":                                      "integer",
		"max_length_for_sort_data":                           "integer",
		"max_points_in_geometry":                             "integer",
		"max_prepared_stmt_count":                            "integer",
		"max_relay_log_size":                                 "integer",
		"max_seeks_for_key":                                  "integer",
		"max_sort_length":                                    "integer",
		"max_sp_recursion_depth":                             "integer",
		"max_statement_time":                                 "integer",
		"max_tmp_tables":                                     "integer",
		"max_user_connections":                               "integer",
		"max_write_lock_count":                               "integer",
		"min_examined_row_limit":                             "integer",
		"multi_range_count":                                  "integer",
		"myisam_data_pointer_size":                           "integer",
		"myisam_max_sort_file_size":                          "integer",
		"myisam_repair_threads":                              "integer",
		"myisam_sort_buffer_size":                            "integer",
		"myisam_stats_method":                                "enumeration",
		"myisam_use_mmap":                                    "boolean",
		"mysql_firewall_mode":                                "boolean",
		"mysql_firewall_trace":                               "boolean",
		"mysql_native_password_proxy_users":                  "boolean",
		"mysqlx-connect-timeout":                             "integer",
		"mysqlx_connect_timeout":                             "integer",
		"mysqlx-idle-worker-thread-timeout":                  "integer",
		"mysqlx_idle_worker_thread_timeout":                  "integer",
		"mysqlx-max-allowed-packet":                          "integer",
		"mysqlx_max_allowed_packet":                          "integer",
		"mysqlx-max-connections":                             "integer",
		"mysqlx_max_connections":                             "integer",
		"mysqlx-min-worker-threads":                          "integer",
		"mysqlx_min_worker_threads":                          "integer",
		"ndb-allow-copying-alter-table":                      "boolean",
		"ndb_autoincrement_prefetch_sz":                      "integer",
		"ndb_blob_read_batch_bytes":                          "integer",
		"ndb_blob_write_batch_bytes":                         "integer",
		"ndb_cache_check_time":                               "integer",
		"ndb_clear_apply_status":                             "boolean",
		"ndb_data_node_neighbour":                            "integer",
		"ndb_default_column_format":                          "enumeration",
		"ndb_deferred_constraints":                           "integer",
		"ndb_distribution":                                   "enumeration",
		"ndb_distribution={KEYHASH|LINHASH}":                 "enumeration",
		"ndb_eventbuffer_free_percent":                       "integer",
		"ndb_eventbuffer_max_alloc":                          "integer",
		"ndb_extra_logging":                                  "integer",
		"ndb_force_send":                                     "boolean",
		"ndb_fully_replicated":                               "boolean",
		"ndb_index_stat_enable":                              "boolean",
		"ndb_index_stat_option":                              "string",
		"ndb_join_pushdown":                                  "boolean",
		"ndb_log_bin":                                        "boolean",
		"ndb_log_binlog_index":                               "boolean",
		"ndb_log_empty_epochs":                               "boolean",
		"ndb_log_empty_update":                               "boolean",
		"ndb_log_exclusive_reads":                            "boolean",
		"ndb_log_update_minimal":                             "boolean",
		"ndb_log_updated_only":                               "boolean",
		"ndb_optimization_delay":                             "integer",
		"ndb_read_backup":                                    "boolean",
		"ndb_recv_thread_cpu_mask":                           "bitmap",
		"ndb_report_thresh_binlog_epoch_slip":                "integer",
		"ndb_report_thresh_binlog_mem_usage":                 "integer",
		"ndb_show_foreign_key_mock_tables":                   "boolean",
		"ndb_slave_last_conflict_epoch":                      "enumeration",
		"ndb_use_exact_count":                                "boolean",
		"ndb_use_transactions":                               "boolean",
		"ndbinfo_max_bytes":                                  "integer",
		"ndbinfo_max_rows":                                   "integer",
		"ndbinfo_offline":                                    "boolean",
		"ndbinfo_show_hidden":                                "boolean",
		"ndbinfo_table_prefix":                               "string",
		"net_buffer_length":                                  "integer",
		"net_read_timeout":                                   "integer",
		"net_retry_count":                                    "integer",
		"net_write_timeout":                                  "integer",
		"new":                                                "boolean",
		"offline_mode":                                       "boolean",
		"old_alter_table":                                    "boolean",
		"old_passwords":                                      "enumeration",
		"optimizer_prune_level":                              "boolean",
		"optimizer_search_depth":                             "integer",
		"optimizer_switch":                                   "set",
		"optimizer_trace":                                    "string",
		"optimizer_trace_features":                           "string",
		"optimizer_trace_limit":                              "integer",
		"optimizer_trace_max_mem_size":                       "integer",
		"optimizer_trace_offset":                             "integer",
		"parser_max_mem_size":                                "integer",
		"preload_buffer_size":                                "integer",
		"profiling":                                          "boolean",
		"profiling_history_size":                             "integer",
		"query_alloc_block_size":                             "integer",
		"query_cache_limit":                                  "integer",
		"query_cache_min_res_unit":                           "integer",
		"query_cache_size":                                   "integer",
		"query_cache_type":                                   "enumeration",
		"query_cache_wlock_invalidate":                       "boolean",
		"query_prealloc_size":                                "integer",
		"range_alloc_block_size":                             "integer",
		"range_optimizer_max_mem_size":                       "integer",
		"read_buffer_size":                                   "integer",
		"read_only":                                          "boolean",
		"read_rnd_buffer_size":                               "integer",
		"relay_log_info_repository":                          "string",
		"relay_log_purge":                                    "boolean",
		"require_secure_transport":                           "boolean",
		"rewriter_enabled":                                   "boolean",
		"rewriter_verbose":                                   "integer",
		"rpl_semi_sync_master_enabled":                       "boolean",
		"rpl_semi_sync_master_timeout":                       "integer",
		"rpl_semi_sync_master_trace_level":                   "integer",
		"rpl_semi_sync_master_wait_for_slave_count":          "integer",
		"rpl_semi_sync_master_wait_no_slave":                 "boolean",
		"rpl_semi_sync_master_wait_point":                    "enumeration",
		"rpl_semi_sync_slave_enabled":                        "boolean",
		"rpl_semi_sync_slave_trace_level":                    "integer",
		"rpl_stop_slave_timeout":                             "integer",
		"secure_auth":                                        "boolean",
		"server_id":                                          "integer",
		"session_track_gtids":                                "enumeration",
		"session_track_schema":                               "boolean",
		"session_track_state_change":                         "boolean",
		"session_track_system_variables":                     "string",
		"sha256_password_proxy_users":                        "boolean",
		"show_compatibility_56":                              "boolean",
		"show_old_temporals":                                 "boolean",
		"slave_allow_batching":                               "boolean",
		"slave_checkpoint_group=#":                           "integer",
		"slave_checkpoint_period=#":                          "integer",
		"slave_compressed_protocol":                          "boolean",
		"slave_exec_mode":                                    "enumeration",
		"slave_max_allowed_packet":                           "integer",
		"slave_net_timeout":                                  "integer",
		"slave_parallel_type":                                "enumeration",
		"slave_parallel_workers":                             "integer",
		"slave_pending_jobs_size_max":                        "integer",
		"slave_preserve_commit_order":                        "boolean",
		"slave_rows_search_algorithms=list":                  "set",
		"slave_sql_verify_checksum":                          "boolean",
		"slave_transaction_retries":                          "integer",
		"slow_launch_time":                                   "integer",
		"slow_query_log":                                     "boolean",
		"slow_query_log_file":                                "filename",
		"sort_buffer_size":                                   "integer",
		"sql_auto_is_null":                                   "boolean",
		"sql_big_selects":                                    "boolean",
		"sql_buffer_result":                                  "boolean",
		"sql_log_off":                                        "boolean",
		"sql_mode":                                           "set",
		"sql_notes":                                          "boolean",
		"sql_quote_show_create":                              "boolean",
		"sql_safe_updates":                                   "boolean",
		"sql_select_limit":                                   "integer",
		"sql_slave_skip_counter":                             "integer",
		"sql_warnings":                                       "boolean",
		"storage_engine":                                     "enumeration",
		"stored_program_cache":                               "integer",
		"super_read_only":                                    "boolean",
		"sync_binlog":                                        "integer",
		"sync_frm":                                           "boolean",
		"sync_master_info":                                   "integer",
		"sync_relay_log":                                     "integer",
		"sync_relay_log_info":                                "integer",
		"table_definition_cache":                             "integer",
		"table_open_cache":                                   "integer",
		"thread_cache_size":                                  "integer",
		"thread_pool_high_priority_connection":               "integer",
		"thread_pool_max_unused_threads":                     "integer",
		"thread_pool_prio_kickup_timer":                      "integer",
		"thread_pool_stall_limit":                            "integer",
		"time_zone":                                          "string",
		"timed_mutexes":                                      "boolean",
		"tmp_table_size":                                     "integer",
		"transaction_alloc_block_size":                       "integer",
		"tx_isolation":                                       "enumeration",
		"transaction_prealloc_size":                          "integer",
		"tx_read_only":                                       "boolean",
		"transaction_write_set_extraction":                   "enumeration",
		"unique_checks":                                      "boolean",
		"updatable_views_with_limit":                         "boolean",
		"validate_password_check_user_name":                  "boolean",
		"validate_password_dictionary_file":                  "filename",
		"validate_password_length":                           "integer",
		"validate_password_mixed_case_count":                 "integer",
		"validate_password_number_count":                     "integer",
		"validate_password_policy":                           "enumeration",
		"validate_password_special_char_count":               "integer",
		"version_tokens_session":                             "string",
		"wait_timeout":                                       "integer",
	}
)

// show variables
func ShowVariables(db *sql.DB) ([]Variable, error) {
	var vars []Variable

	rows, err := db.Query(StmtShowVariables)
	if err != nil {
		return []Variable{}, err
	}

	for rows.Next() {
		var tmpvar Variable

		err = rows.Scan(
			&tmpvar.Name,
			&tmpvar.Value,
		)

		if val, ok := GlobalDynamicVars[tmpvar.Name]; ok {
			tmpvar.Dynamic = "Yes"
			tmpvar.Type = val
		} else {
			tmpvar.Dynamic = "No"
			tmpvar.Type = "UNKNOWN"
		}

		if err != nil {
			continue
		}

		vars = append(vars, tmpvar)
	}
	return vars, nil
}

// set global or session dynamic variables
func SetDynamicVariables(db *sql.DB, variable_name string, variable_value string) error {

	var Query string

	varvalue, ok := GlobalDynamicVars[variable_name]
	if !ok {
		return errors.New(fmt.Sprintf("Not found: %s", variable_name))
	}

	switch {
	case varvalue == "integer":
		_, err := strconv.Atoi(variable_value)
		if err != nil {
			return err
		}
		Query = fmt.Sprintf(StmtSetIntegerDynamicVariables, variable_name, variable_value)
	case varvalue == "boolean":
		switch {
		case variable_value == "ON" || variable_value == "True" || variable_value == "true" || variable_value == "TRUE" || variable_value == "on" || variable_value == "On" || variable_value == "1":
			Query = fmt.Sprintf(StmtSetStringDynamicVariables, variable_name, "ON")
		case variable_value == "OFF" || variable_value == "False" || variable_value == "false" || variable_value == "FALSE" || variable_value == "off" || variable_value == "Off" || variable_value == "0":
			Query = fmt.Sprintf(StmtSetStringDynamicVariables, variable_name, "OFF")
		default:
			return errors.New(fmt.Sprintf("Not valid: %s", variable_value))
		}
	default:
		Query = fmt.Sprintf(StmtSetStringDynamicVariables, variable_name, variable_value)
	}

	_, err := db.Exec(Query)
	if err != nil {
		return err
	}
	return nil
}