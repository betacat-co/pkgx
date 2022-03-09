package timex

const (
	TimeLayoutISO                = "2006-01-02 15:04"
	TimeLayoutNoSeconds          = "2006.01.02 15:04"
	TimeLayoutUS                 = "January 2, 2006"
	TimeLayoutFolder             = "20060102"
	TimeLayoutNoDate             = "1504011234"
	TimeLayoutReportList         = "2006/01/02"
	TimeLayoutDeviceXml          = "01/02/2006"
	TimeLayoutDeviceXmlRunAt     = "01/02/2006 15:04"
	TimeLayoutDeviceXmlMessageAt = "01/02/2006 15:04:05 AM"

	YearMonth        = "2006-01"
	YearMonthDay     = "2006-01-02"
	HourMinuteSecond = "15:04:05"
	HourMinute       = "15:04"
	DefaultLayout    = YearMonthDay + " " + HourMinuteSecond
)
