package errorresponses

type ErrorRespones struct {
	LogSysNo int    `json:"log_sys_no"`
	Message  string `json:"message"`
	Success  bool   `json:"success"`
}
