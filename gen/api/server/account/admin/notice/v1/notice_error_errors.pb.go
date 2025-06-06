// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package adminv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsNoticeAdminErrorReasonUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == NoticeAdminErrorReason_NOTICE_ADMIN_ERROR_REASON_UNSPECIFIED.String() && e.Code == 500
}

func ErrorNoticeAdminErrorReasonUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(500, NoticeAdminErrorReason_NOTICE_ADMIN_ERROR_REASON_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsNoticeAdminErrorReasonServer(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == NoticeAdminErrorReason_NOTICE_ADMIN_ERROR_REASON_SERVER.String() && e.Code == 500
}

func ErrorNoticeAdminErrorReasonServer(format string, args ...interface{}) *errors.Error {
	return errors.New(500, NoticeAdminErrorReason_NOTICE_ADMIN_ERROR_REASON_SERVER.String(), fmt.Sprintf(format, args...))
}

func IsNoticeAdminErrorReasonServerId(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == NoticeAdminErrorReason_NOTICE_ADMIN_ERROR_REASON_SERVER_ID.String() && e.Code == 401
}

func ErrorNoticeAdminErrorReasonServerId(format string, args ...interface{}) *errors.Error {
	return errors.New(401, NoticeAdminErrorReason_NOTICE_ADMIN_ERROR_REASON_SERVER_ID.String(), fmt.Sprintf(format, args...))
}
