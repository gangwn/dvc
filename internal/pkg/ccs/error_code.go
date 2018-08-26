package ccs

type CCSErrorCode int32

const (
	CCSErrorCode_Success   CCSErrorCode = 0
	CCSErrorCode_AlreadyInMeeting  CCSErrorCode = 1
	CCSErrorCode_NotInMeeting  CCSErrorCode = 1
)