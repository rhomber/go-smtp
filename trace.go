package smtp

type TransmitMode string
type SmtpFacility string

const (
	TransmitModeTx TransmitMode = "TX"
	TransmitModeRx TransmitMode = "RX"

	SmtpFacilityHelo      SmtpFacility = "HELO"
	SmtpFacilityStartTLS  SmtpFacility = "STARTTLS"
	SmtpFacilityVerify    SmtpFacility = "VRFY"
	SmtpFacilityAuth      SmtpFacility = "AUTH"
	SmtpFacilityMail      SmtpFacility = "MAIL"
	SmtpFacilityRcpt      SmtpFacility = "RCPT"
	SmtpFacilityData      SmtpFacility = "DATA"
	SmtpFacilityDataClose SmtpFacility = "DATA_CLOSE"
	SmtpFacilityReset     SmtpFacility = "RSET"
	SmtpFacilityNoOp      SmtpFacility = "NOOP"
	SmtpFacilityQuit      SmtpFacility = "QUIT"
)

type TraceEmitter interface {
	Emit(mode TransmitMode, facility SmtpFacility, msg string, code int)
}

func (c *Client) traceTx(facility SmtpFacility, msg string) {
	c.trace(TransmitModeTx, facility, msg, 0)
}

func (c *Client) traceRx(facility SmtpFacility, code int, msg string) {
	c.trace(TransmitModeRx, facility, msg, code)
}

func (c *Client) trace(mode TransmitMode, facility SmtpFacility, msg string, code int) {
	if c.traceEmitter != nil {
		c.traceEmitter.Emit(mode, facility, msg, code)
	}
}
