package uptimerobot

type UptimeMonitorGetMonitorsResponse struct {
	Stat       string                  `json:"stat"`
	Pagination UptimeMonitorPagination `json:"pagination"`
	Monitors   []UptimeMonitorMonitor  `json:"monitors"`
}

type UptimeMonitorPagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

type UptimeMonitorMonitor struct {
	ID                int                          `json:"id"`
	FriendlyName      string                       `json:"friendly_name"`
	URL               string                       `json:"url"`
	Type              int                          `json:"type"`
	SubType           string                       `json:"sub_type"`
	KeywordType       int                          `json:"keyword_type"`
	KeywordValue      string                       `json:"keyword_value"`
	HTTPUsername      string                       `json:"http_username"`
	HTTPPassword      string                       `json:"http_password"`
	Port              string                       `json:"port"`
	Interval          int                          `json:"interval"`
	Status            int                          `json:"status"`
	CreateDatetime    int                          `json:"create_datetime"`
	Logs              []UptimeMonitorLogs          `json:"logs"`
	AlertContacts     []UptimeMonitorAlertContacts `json:"alert_contacts"`
	SSL               int                          `json:"ssl"`
	Timeout           int                          `json:"timeout"`
	CustomHTTPStatuses string                      `json:"custom_http_statuses"`
	CustomHeader      string                       `json:"custom_header"`
}

type UptimeMonitorAlertContacts struct {
	ID         string `json:"id"`
	Threshold  int    `json:"threshold"`
	Recurrence int    `json:"recurrence"`
}

type UptimeMonitorLogs struct {
	Type     int `json:"type"`
	Datetime int `json:"datetime"`
	Duration int `json:"duration"`
}

type UptimeMonitorNewMonitorResponse struct {
	Stat    string                     `json:"stat"`
	Monitor UptimeMonitorMonitorStatus `json:"monitor"`
	Error   UptimeMonitorError         `json:"error"`
}

type UptimeMonitorError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type UptimeMonitorMonitorStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

type UptimeMonitorStatusMonitorResponse struct {
	Stat    string             `json:"stat"`
	Error   UptimeMonitorError `json:"error"`
	Monitor struct {
		ID int `json:"id"`
	} `json:"monitor"`
}

type UptimePublicStatusPage struct {
	ID           int    `json:"id"`
	FriendlyName string `json:"friendly_name"`
	Monitors     []int  `json:"monitors"`
	CustomDomain string `json:"custom_domain"`
	Password     string `json:"password"`
	Sort         int    `json:"sort"`
	Status       int    `json:"status"`
}

type UptimeStatusPageResponse struct {
	Stat                   string `json:"stat"`
	UptimePublicStatusPage struct {
		ID int `json:"id"`
	} `json:"psp"`
}

type UptimeStatusPagesResponse struct {
	Stat       string `json:"stat"`
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Total  int `json:"total"`
	} `json:"pagination"`
	StatusPages []UptimePublicStatusPage `json:"psps"`
}