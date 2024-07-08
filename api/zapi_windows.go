// mksyscall_windows.pl api.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package api

import (
	"os"
	"syscall"
	"unsafe"
	trc "github.com/ibmdb/go_ibm_db/log2"
	"fmt"
)

var (
	mododbc32 = syscall.NewLazyDLL(GetDllName())

	procSQLAllocHandle     = mododbc32.NewProc("SQLAllocHandle")
	procSQLBindCol         = mododbc32.NewProc("SQLBindCol")
	procSQLBindParameter   = mododbc32.NewProc("SQLBindParameter")
	procSQLCloseCursor     = mododbc32.NewProc("SQLCloseCursor")
	procSQLDescribeColW    = mododbc32.NewProc("SQLDescribeColW")
	procSQLDescribeParam   = mododbc32.NewProc("SQLDescribeParam")
	procSQLDisconnect      = mododbc32.NewProc("SQLDisconnect")
	procSQLDriverConnectW  = mododbc32.NewProc("SQLDriverConnectW")
	procSQLEndTran         = mododbc32.NewProc("SQLEndTran")
	procSQLExecute         = mododbc32.NewProc("SQLExecute")
	procSQLFetch           = mododbc32.NewProc("SQLFetch")
	procSQLFetchScroll     = mododbc32.NewProc("SQLFetchScroll")
	procSQLFreeHandle      = mododbc32.NewProc("SQLFreeHandle")
	procSQLGetData         = mododbc32.NewProc("SQLGetData")
	procSQLGetDiagRecW     = mododbc32.NewProc("SQLGetDiagRecW")
	procSQLNumParams       = mododbc32.NewProc("SQLNumParams")
	procSQLNumResultCols   = mododbc32.NewProc("SQLNumResultCols")
	procSQLPrepareW        = mododbc32.NewProc("SQLPrepareW")
	procSQLExecDirectW     = mododbc32.NewProc("SQLExecDirectW")
	procSQLRowCount        = mododbc32.NewProc("SQLRowCount")
	procSQLSetEnvAttr      = mododbc32.NewProc("SQLSetEnvAttr")
	procSQLSetConnectAttrW = mododbc32.NewProc("SQLSetConnectAttrW")
	procSQLColAttribute    = mododbc32.NewProc("SQLColAttribute")
	procSQLMoreResults     = mododbc32.NewProc("SQLMoreResults")
	procSQLSetStmtAttr     = mododbc32.NewProc("SQLSetStmtAttr")
	procSQLSetStmtAttrW    = mododbc32.NewProc("SQLSetStmtAttrW")
	procSQLGetStmtAttr     = mododbc32.NewProc("SQLGetStmtAttr")
	procSQLSetPos	       = mododbc32.NewProc("SQLSetPos")
	procSQLBulkOperations  = mododbc32.NewProc("SQLBulkOperations")
	procSQLCreateDb        = mododbc32.NewProc("SQLCreateDbW")
	procSQLDropDb          = mododbc32.NewProc("SQLDropDbW")
)

func GetDllName() string {
	trc.Trace1("api/zapi_windows.go GetDllName()")
	if winArch := os.Getenv("PROCESSOR_ARCHITECTURE"); winArch == "x86" {
	    trc.Trace1("winArch is x86 and db2cli.dll")
		return "db2cli.dll"
	} else {
	    trc.Trace1("db2cli64.dll")
		return "db2cli64.dll"
	}
}

func SQLAllocHandle(handleType SQLSMALLINT, inputHandle SQLHANDLE, outputHandle *SQLHANDLE) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLAllocHandle() - ENTRY")
	trc.Trace1(fmt.Sprintf("handleType=%d", handleType))

	r0, _, _ := syscall.Syscall(procSQLAllocHandle.Addr(), 3, uintptr(handleType), uintptr(inputHandle), uintptr(unsafe.Pointer(outputHandle)))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLAllocHandle()- EXIT")
	return
}

func SQLBindCol(statementHandle SQLHSTMT, columnNumber SQLUSMALLINT, targetType SQLSMALLINT, targetValuePtr []byte, bufferLength SQLLEN, vallen *SQLLEN) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLBindCol() - ENTRY")
	trc.Trace1(fmt.Sprintf("columnNumber=%d, targetType=%d, bufferLength=%d", columnNumber, targetType, bufferLength))

	r0, _, _ := syscall.Syscall6(procSQLBindCol.Addr(), 6, uintptr(statementHandle), uintptr(columnNumber), uintptr(targetType), uintptr(unsafe.Pointer(&targetValuePtr[0])), uintptr(bufferLength), uintptr(unsafe.Pointer(vallen)))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLBindCol()- EXIT")
	return
}

func SQLBindParameter(statementHandle SQLHSTMT, parameterNumber SQLUSMALLINT, inputOutputType SQLSMALLINT, valueType SQLSMALLINT, parameterType SQLSMALLINT, columnSize SQLULEN, decimalDigits SQLSMALLINT, parameterValue SQLPOINTER, bufferLength SQLLEN, ind *SQLLEN) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLBindParameter() - ENTRY")
	trc.Trace1(fmt.Sprintf("parameterNumber=%d, inputOutputType=%d, valueType=%d, parameterType=%d, columnSize=%d, decimalDigits=%d, parameterValue=%x, bufferLength=%d", parameterNumber, inputOutputType, valueType, parameterType, columnSize, decimalDigits, parameterValue, bufferLength))

	r0, _, _ := syscall.Syscall12(procSQLBindParameter.Addr(), 10, uintptr(statementHandle), uintptr(parameterNumber), uintptr(inputOutputType), uintptr(valueType), uintptr(parameterType), uintptr(columnSize), uintptr(decimalDigits), uintptr(parameterValue), uintptr(bufferLength), uintptr(unsafe.Pointer(ind)), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLBindParameter()- EXIT")
	return
}

func SQLCloseCursor(statementHandle SQLHSTMT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLCloseCursor() - ENTRY")

	r0, _, _ := syscall.Syscall(procSQLCloseCursor.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLCloseCursor() - EXIT")
	return
}

func SQLDescribeCol(statementHandle SQLHSTMT, columnNumber SQLUSMALLINT, columnName *SQLWCHAR, bufferLength SQLSMALLINT, nameLengthPtr *SQLSMALLINT, dataTypePtr *SQLSMALLINT, columnSizePtr *SQLULEN, decimalDigitsPtr *SQLSMALLINT, nullablePtr *SQLSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLDescribeCol() - ENTRY")
	trc.Trace1(fmt.Sprintf("columnNumber=%d, columnName=%s, bufferLength=%d, nameLengthPtr=0x%x, dataTypePtr=0x%x, columnSizePtr=0x%x, decimalDigitsPtr=%x, nullablePtr=0x%x", columnNumber, *columnName, bufferLength, nameLengthPtr, dataTypePtr, columnSizePtr, decimalDigitsPtr, nullablePtr))

	r0, _, _ := syscall.Syscall9(procSQLDescribeColW.Addr(), 9, uintptr(statementHandle), uintptr(columnNumber), uintptr(unsafe.Pointer(columnName)), uintptr(bufferLength), uintptr(unsafe.Pointer(nameLengthPtr)), uintptr(unsafe.Pointer(dataTypePtr)), uintptr(unsafe.Pointer(columnSizePtr)), uintptr(unsafe.Pointer(decimalDigitsPtr)), uintptr(unsafe.Pointer(nullablePtr)))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLDescribeCol()- EXIT")
	return
}

func SQLDescribeParam(statementHandle SQLHSTMT, parameterNumber SQLUSMALLINT, dataTypePtr *SQLSMALLINT, parameterSizePtr *SQLULEN, decimalDigitsPtr *SQLSMALLINT, nullablePtr *SQLSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLDescribeParam() - ENTRY")
	trc.Trace1(fmt.Sprintf("parameterNumber=%d, dataTypePtr=%x, parameterSizePtr=%x, decimalDigitsPtr=%x, nullablePtr=%x", parameterNumber, dataTypePtr, parameterSizePtr, decimalDigitsPtr, nullablePtr))

	r0, _, _ := syscall.Syscall6(procSQLDescribeParam.Addr(), 6, uintptr(statementHandle), uintptr(parameterNumber), uintptr(unsafe.Pointer(dataTypePtr)), uintptr(unsafe.Pointer(parameterSizePtr)), uintptr(unsafe.Pointer(decimalDigitsPtr)), uintptr(unsafe.Pointer(nullablePtr)))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLDescribeParam()- EXIT")
	return
}

func SQLDisconnect(connectionHandle SQLHDBC) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLDisconnect() - ENTRY")

	r0, _, _ := syscall.Syscall(procSQLDisconnect.Addr(), 1, uintptr(connectionHandle), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLDisconnect()- EXIT")
	return
}

func SQLDriverConnect(connectionHandle SQLHDBC, windowHandle SQLHWND, inConnectionString *SQLWCHAR, stringLength1 SQLSMALLINT, outConnectionString *SQLWCHAR, bufferLength SQLSMALLINT, stringLength2Ptr *SQLSMALLINT, driverCompletion SQLUSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLDriverConnect()- ENTRY")
	trc.Trace1(fmt.Sprintf("inConnectionString=%s, stringLength1=%d, outConnectionString=%s, bufferLength=%d, stringLength2Ptr=%x, driverCompletion=%d", inConnectionString, stringLength1, outConnectionString, bufferLength, stringLength2Ptr, driverCompletion))

	r0, _, _ := syscall.Syscall9(procSQLDriverConnectW.Addr(), 8, uintptr(connectionHandle), uintptr(windowHandle), uintptr(unsafe.Pointer(inConnectionString)), uintptr(stringLength1), uintptr(unsafe.Pointer(outConnectionString)), uintptr(bufferLength), uintptr(unsafe.Pointer(stringLength2Ptr)), uintptr(driverCompletion), 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLDriverConnect()- EXIT")
	return
}

func SQLEndTran(handleType SQLSMALLINT, handle SQLHANDLE, completionType SQLSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLEndTran() - ENTRY")
	trc.Trace1(fmt.Sprintf("completionType=%d", completionType))

	r0, _, _ := syscall.Syscall(procSQLEndTran.Addr(), 3, uintptr(handleType), uintptr(handle), uintptr(completionType))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLEndTran()- EXIT")
	return
}

func SQLExecute(statementHandle SQLHSTMT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLExecute() - ENTRY")

	r0, _, _ := syscall.Syscall(procSQLExecute.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLExecute()- EXIT")
	return
}

func SQLFetch(statementHandle SQLHSTMT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLFetch() - ENTRY")

	r0, _, _ := syscall.Syscall(procSQLFetch.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLFetch()- EXIT")
	return
}

func SQLFetchScroll(statementHandle SQLHSTMT, fetchOrientation SQLUSMALLINT, fetchOffset SQLLEN) (ret SQLRETURN) {
        trc.Trace1("api/zapi_windows.go SQLFetchScroll() - ENTRY")
	trc.Trace1(fmt.Sprintf("fetchOrientation=%d, fetchOffset=%d", fetchOrientation, fetchOffset))

	r0, _, _ := syscall.Syscall(procSQLFetch.Addr(), 3, uintptr(statementHandle), uintptr(fetchOrientation), uintptr(fetchOffset))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLFetchScroll()- EXIT")
	return
}

func SQLFreeHandle(handleType SQLSMALLINT, handle SQLHANDLE) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLFreeHandle() - ENTRY")
	trc.Trace1(fmt.Sprintf("handleType=%d",  handleType))

	r0, _, _ := syscall.Syscall(procSQLFreeHandle.Addr(), 2, uintptr(handleType), uintptr(handle), 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLFreeHandle()- EXIT")
	return
}

func SQLGetData(statementHandle SQLHSTMT, colOrParamNum SQLUSMALLINT, targetType SQLSMALLINT, targetValuePtr SQLPOINTER, bufferLength SQLLEN, vallen *SQLLEN) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLGetData() - ENTRY")
	trc.Trace1(fmt.Sprintf("colOrParamNum=%d, targetType=%d, targetValuePtr=%x, bufferLength=%d, vallen=%d", colOrParamNum, targetType, targetValuePtr, bufferLength, vallen))

	r0, _, _ := syscall.Syscall6(procSQLGetData.Addr(), 6, uintptr(statementHandle), uintptr(colOrParamNum), uintptr(targetType), uintptr(targetValuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(vallen)))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLGetData()- EXIT")
	return
}

func SQLGetDiagRec(handleType SQLSMALLINT, handle SQLHANDLE, recNumber SQLSMALLINT, sqlState *SQLWCHAR, nativeErrorPtr *SQLINTEGER, messageText *SQLWCHAR, bufferLength SQLSMALLINT, textLengthPtr *SQLSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLGetDiagRec() - ENTRY")
	trc.Trace1(fmt.Sprintf("handleType=%d, recNumber=%d, sqlState=%s, nativeErrorPtr=%x, messageText=%s, bufferLength=%d, textLengthPtr=%x", handleType, recNumber, sqlState, nativeErrorPtr, messageText, bufferLength, textLengthPtr))

	r0, _, _ := syscall.Syscall9(procSQLGetDiagRecW.Addr(), 8, uintptr(handleType), uintptr(handle), uintptr(recNumber), uintptr(unsafe.Pointer(sqlState)), uintptr(unsafe.Pointer(nativeErrorPtr)), uintptr(unsafe.Pointer(messageText)), uintptr(bufferLength), uintptr(unsafe.Pointer(textLengthPtr)), 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLGetDiagRec()- EXIT")
	return
}

func SQLNumParams(statementHandle SQLHSTMT, parameterCountPtr *SQLSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLNumParams() - ENTRY")
	trc.Trace1(fmt.Sprintf("parameterCountPtr=%x", parameterCountPtr))

	r0, _, _ := syscall.Syscall(procSQLNumParams.Addr(), 2, uintptr(statementHandle), uintptr(unsafe.Pointer(parameterCountPtr)), 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLNumParams()- EXIT")
	return
}

func SQLNumResultCols(statementHandle SQLHSTMT, columnCountPtr *SQLSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLNumResultCols() - ENTRY")
	trc.Trace1(fmt.Sprintf("columnCountPtr=%x", columnCountPtr))

	r0, _, _ := syscall.Syscall(procSQLNumResultCols.Addr(), 2, uintptr(statementHandle), uintptr(unsafe.Pointer(columnCountPtr)), 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLNumResultCols()- EXIT")
	return
}

func SQLPrepare(statementHandle SQLHSTMT, statementText *SQLWCHAR, textLength SQLINTEGER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLPrepare() - ENTRY")
	trc.Trace1(fmt.Sprintf("statementText=%s, textLength=%d", statementText, textLength))

	r0, _, _ := syscall.Syscall(procSQLPrepareW.Addr(), 3, uintptr(statementHandle), uintptr(unsafe.Pointer(statementText)), uintptr(textLength))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLPrepare()- EXIT")
	return
}

func SQLRowCount(statementHandle SQLHSTMT, rowCountPtr *SQLLEN) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLRowCount() - ENTRY")
	trc.Trace1(fmt.Sprintf("rowCountPtr=0x%x", rowCountPtr))

	r0, _, _ := syscall.Syscall(procSQLRowCount.Addr(), 2, uintptr(statementHandle), uintptr(unsafe.Pointer(rowCountPtr)), 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLRowCount()- EXIT")
	return
}

func SQLSetEnvAttr(environmentHandle SQLHENV, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLSetEnvAttr() - ENTRY")
	trc.Trace1(fmt.Sprintf("attribute=%d, valuePtr=0x%x, stringLength=%d", attribute, valuePtr, stringLength))

	r0, _, _ := syscall.Syscall6(procSQLSetEnvAttr.Addr(), 4, uintptr(environmentHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLSetEnvAttr()- EXIT")
	return
}

func SQLSetConnectAttr(connectionHandle SQLHDBC, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLSetConnectAttr() - ENTRY")
	trc.Trace1(fmt.Sprintf("attribute=%d, valuePtr=0x%x, stringLength=%d", attribute, valuePtr, stringLength))

	r0, _, _ := syscall.Syscall6(procSQLSetConnectAttrW.Addr(), 4, uintptr(connectionHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLSetConnectAttr()- EXIT")
	return
}

func SQLColAttribute(statementHandle SQLHSTMT, ColumnNumber SQLUSMALLINT, FieldIdentifier SQLUSMALLINT, CharacterAttributePtr SQLPOINTER, BufferLength SQLSMALLINT, StringLengthPtr *SQLSMALLINT, NumericAttributePtr SQLPOINTER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLColAttribute() - ENTRY")
	trc.Trace1(fmt.Sprintf("ColumnNumber=%d, FieldIdentifier=%d, CharacterAttributePtr=0x%x, BufferLength=%d, StringLengthPtr=0x%x, NumericAttributePtr=0x%x", ColumnNumber, FieldIdentifier, CharacterAttributePtr, BufferLength, StringLengthPtr, NumericAttributePtr))

	r0, _, _ := syscall.Syscall9(procSQLColAttribute.Addr(), 7, uintptr(statementHandle), uintptr(ColumnNumber), uintptr(FieldIdentifier), uintptr(CharacterAttributePtr), uintptr(BufferLength), uintptr(unsafe.Pointer(StringLengthPtr)), uintptr(NumericAttributePtr), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLColAttribute()- EXIT")
	return
}

func SQLMoreResults(statementHandle SQLHSTMT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLMoreResults() - ENTRY")

	r0, _, _ := syscall.Syscall(procSQLMoreResults.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLMoreResults()- EXIT")
	return
}

func SQLSetStmtAttr(statementHandle SQLHSTMT, attribute SQLINTEGER, valuePtr SQLPOINTER, stringLength SQLINTEGER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLSetStmtAttr() - ENTRY")
	trc.Trace1(fmt.Sprintf("attribute=%d, valuePtr=0x%x, stringLength=%d", attribute, valuePtr, stringLength))

	r0, _, _ := syscall.Syscall6(procSQLSetStmtAttr.Addr(), 4, uintptr(statementHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLSetStmtAttr()- EXIT")
	return
}

func SQLGetStmtAttr(environmentHandle SQLHSTMT, attribute SQLINTEGER, targetValuePtr []byte, stringLength SQLINTEGER) (ret SQLRETURN) {
        trc.Trace1("api/zapi_windows.go SQLGetStmtAttr() - ENTRY")
	trc.Trace1(fmt.Sprintf("attribute=%d, stringLength=%d", attribute, stringLength))

	r0, _, _ := syscall.Syscall6(procSQLGetStmtAttr.Addr(), 4, uintptr(environmentHandle), uintptr(attribute), uintptr(unsafe.Pointer(&targetValuePtr[0])), uintptr(stringLength), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLGetStmtAttr()- EXIT")
	return
}

func SQLSetPos(statementHandle SQLHSTMT, rowNumber SQLSETPOSIROW, operation SQLUSMALLINT, lockType SQLUSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLSetPos() - ENTRY")
	trc.Trace1(fmt.Sprintf("rowNumber=%d, operation=%d, lockType=%d", rowNumber, operation, lockType))

	r0, _, _ := syscall.Syscall6(procSQLSetPos.Addr(), 4, uintptr(statementHandle), uintptr(rowNumber), uintptr(operation), uintptr(lockType), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLSetPos()- EXIT")
	return
}

func SQLBulkOperations(statementHandle SQLHSTMT, operation SQLSMALLINT) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLBulkOperations() - ENTRY")
	trc.Trace1(fmt.Sprintf("operation=%d", operation))

	r0, _, _ := syscall.Syscall6(procSQLBulkOperations.Addr(), 2, uintptr(statementHandle), uintptr(operation), 0, 0, 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLBulkOperations()- EXIT")
	return
}

func SQLCreateDb(connectionHandle SQLHDBC, dbnamePtr *SQLWCHAR, dbnameLen SQLINTEGER, codeSetPtr *SQLWCHAR, codeSetLen SQLINTEGER, modePtr *SQLWCHAR, modeLen SQLINTEGER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLCreateDb() - ENTRY")
	trc.Trace1(fmt.Sprintf("dbnamePtr=%x,dbnameLen=%d, codeSetPtr=%x, codeSetLen=%d, modePtr=%x, modeLen=%d", dbnamePtr, dbnameLen, codeSetPtr, codeSetLen, modePtr, modeLen))

	r0, _, _ := syscall.Syscall9(procSQLCreateDb.Addr(), 7, uintptr(connectionHandle), uintptr(unsafe.Pointer(dbnamePtr)), uintptr(dbnameLen), uintptr(unsafe.Pointer(codeSetPtr)), uintptr(codeSetLen), uintptr(unsafe.Pointer(modePtr)), uintptr(modeLen), 0, 0)
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLCreateDb()- EXIT")
	return
}

func SQLDropDb(connectionHandle SQLHDBC, dbnamePtr *SQLWCHAR, dbnameLen SQLINTEGER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLDropDb() - ENTRY")
	trc.Trace1(fmt.Sprintf("dbnamePtr=%x, dbnameLen=%d", dbnamePtr, dbnameLen))

	r0, _, _ := syscall.Syscall(procSQLDropDb.Addr(), 3, uintptr(connectionHandle), uintptr(unsafe.Pointer(dbnamePtr)), uintptr(dbnameLen))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLDropDb()- EXIT")
	return
}

func SQLExecDirect(statementHandle SQLHSTMT, statementText *SQLWCHAR, textLength SQLINTEGER) (ret SQLRETURN) {
	trc.Trace1("api/zapi_windows.go SQLExecDirect() - ENTRY")
	trc.Trace1(fmt.Sprintf("statementText=%s, textLength=%d", statementText, textLength))

	r0, _, _ := syscall.Syscall(procSQLExecDirectW.Addr(), 3, uintptr(statementHandle), uintptr(unsafe.Pointer(statementText)), uintptr(textLength))
	ret = SQLRETURN(r0)

	trc.Trace1(fmt.Sprintf("ret = %d", ret))
	trc.Trace1("api/zapi_windows.go SQLExecDirect()- EXIT")
	return
}
