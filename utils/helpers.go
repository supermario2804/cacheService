package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"path"
// 	"reflect"
// 	"runtime"
// 	"time"
// 	"unicode"

// 	excelize "github.com/360EntSecGroup-Skylar/excelize"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	validator "gopkg.in/go-playground/validator.v10"
// )

// type Pagination struct {
// 	Page         int   `json:"page"`
// 	Total        int64 `json:"total"`
// 	Limit        int   `json:"limit"`
// 	NextPage     bool  `json:"nextPage"`
// 	PreviousPage bool  `json:"previousPage"`
// }

// /*
// Common function to convert interface into json and send response
// */
func SendHTTPResponse(w http.ResponseWriter, model interface{}) {

	json, err := json.Marshal(model)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusRequestTimeout)
		Println(w, "Something went wrong", http.StatusRequestTimeout)
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(json))
}

// func GetPagination(count int64, Limit int, Offset int, Page int) Pagination {

// 	Println("IN () => GetPagination")

// 	Println("count:", count, "Limit =>", Limit, "Offset => ", Offset, " Page =>", Page)
// 	pagination := Pagination{
// 		NextPage:     false,
// 		Page:         Page,
// 		Total:        0,
// 		Limit:        Limit,
// 		PreviousPage: false,
// 	}

// 	pagination.Total = count

// 	Println("Total: ", count)

// 	Println("Offset+Limit ", Offset, Limit)

// 	if int64(Offset)+int64(Limit) < count {
// 		pagination.NextPage = true
// 	}

// 	if Offset >= Limit {
// 		pagination.PreviousPage = true
// 	}

// 	return pagination

// }

// /*
// ValidateBody is
// */
// func ValidateBody(requestBody io.ReadCloser, requestContainer interface{}) bool {

// 	defer requestBody.Close()

// 	body, IOErr := ioutil.ReadAll(requestBody)

// 	if IOErr != nil {
// 		PrintFatal("Error reading body: ", IOErr)
// 		return false
// 	}

// 	JSONErr := json.Unmarshal(body, &requestContainer)

// 	if JSONErr != nil {
// 		PrintFatal("Error reading body: ", JSONErr)
// 		return false
// 	}

// 	return true
// }

// /*
// DD is Dump and die
// */
// func DD(args ...interface{}) {
// 	for _, arg := range args {
// 		PrintFatal(arg)
// 	}
// 	os.Exit(0)
// }

// func GetSignedUrl(key string) (string, error) {

// 	region := os.Getenv("AWS_REGION")
// 	secret := os.Getenv("AWS_PRIVATE_KEY")
// 	public := os.Getenv("AWS_PUBLIC_KEY")
// 	bucket := os.Getenv("AWS_BUCKET")
// 	if bucket == "" || key == "" {
// 		return "", HandleError(fmt.Errorf("Key / Bucket cannot be empty"))
// 	}

// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String(region),
// 		Credentials: credentials.NewStaticCredentials(public, secret, "")},
// 	)

// 	if err != nil {
// 		return "", HandleError(err)
// 	}

// 	svc := s3.New(sess)

// 	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
// 		Bucket: aws.String(bucket),
// 		Key:    aws.String(key),
// 	})
// 	urlStr, err := req.Presign(15 * time.Minute)

// 	if err != nil {
// 		return "", HandleError(err)
// 	}

// 	return urlStr, nil
// }

// // //this logs the function name as well.
// // func HandleError(err error) error {
// // 	// notice that we're using 1, so it will actually log the where
// // 	// the error happened, 0 = this function, we don't want that.
// // 	pc, fn, line, _ := runtime.Caller(1)

// // 	PrintFatal(fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err))
// // 	return err
// // }

// func GenerateExcel(headers map[string]int, filepath string, data interface{}) error {
// 	f := excelize.NewFile()
// 	sheetName := "Sheet1"
// 	rows := InterfaceSlice(data)

// 	for k, v := range headers {

// 		axis, axisErr := excelize.CoordinatesToCellName(v, 1)
// 		if axisErr != nil {
// 			return HandleError(axisErr)
// 		}
// 		ValErr := f.SetCellValue(sheetName, axis, k)
// 		if ValErr != nil {
// 			return HandleError(ValErr)
// 		}
// 	}

// 	for i, row := range rows {
// 		rowval := reflect.ValueOf(row)
// 		rowType := reflect.TypeOf(row)

// 		for j := 0; j < rowType.NumField(); j++ {
// 			field := rowType.Field(j)
// 			fieldVal := rowval.FieldByName(field.Name)

// 			if tag, ok := field.Tag.Lookup("excel"); ok {

// 				axis, axisErr := excelize.CoordinatesToCellName(headers[tag], i+2)
// 				if axisErr != nil {
// 					return HandleError(axisErr)
// 				}
// 				ValErr := f.SetCellValue(sheetName, axis, fieldVal)
// 				if ValErr != nil {
// 					return HandleError(ValErr)
// 				}

// 			} else {
// 				return HandleError(fmt.Errorf("Particular field has not valid tag : %+v", field))
// 			}
// 		}
// 	}
// 	dir, _ := path.Split(filepath)
// 	dirErr := os.MkdirAll(dir, 0755)
// 	if dirErr != nil {
// 		return dirErr
// 	}

// 	// Save xlsx file by the given path.
// 	if err := f.SaveAs(filepath); err != nil {
// 		fmt.Println(err)
// 	}
// 	return nil
// }

// // AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// // and will set file info like content type and encryption on the uploaded file.
// func AddFileToS3(fileDir string) error {
// 	region := os.Getenv("AWS_REGION")
// 	secret := os.Getenv("AWS_PRIVATE_KEY")
// 	public := os.Getenv("AWS_PUBLIC_KEY")
// 	awsBucket := os.Getenv("AWS_BUCKET")
// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String(region),
// 		Credentials: credentials.NewStaticCredentials(public, secret, "")},
// 	)

// 	if err != nil {
// 		return HandleError(err)
// 	}
// 	// Open the file for use
// 	file, err := os.Open(fileDir)
// 	if err != nil {
// 		return HandleError(err)
// 	}
// 	defer file.Close()

// 	// Get file size and read the file content into a buffer
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		return HandleError(err)
// 	}

// 	var size int64 = fileInfo.Size()
// 	buffer := make([]byte, size)
// 	file.Read(buffer)

// 	// Config settings: this is where you choose the bucket, filename, content-type etc.
// 	// of the file you're uploading.
// 	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
// 		Bucket: aws.String(awsBucket),
// 		Key:    aws.String(fileDir),
// 		ACL:    aws.String("private"),
// 		Body:   bytes.NewReader(buffer),
// 	})

// 	if err != nil {
// 		return HandleError(err)
// 	}

// 	return nil
// }

// func InterfaceSlice(slice interface{}) []interface{} {
// 	s := reflect.ValueOf(slice)
// 	if s.Kind() != reflect.Slice {
// 		panic("InterfaceSlice() given a non-slice type")
// 	}

// 	ret := make([]interface{}, s.Len())

// 	for i := 0; i < s.Len(); i++ {
// 		ret[i] = s.Index(i).Interface()
// 	}

// 	return ret
// }

// func ValidateInputs(dataset interface{}) (error, bool, map[string]string) {
// 	validate := validator.New()
// 	errors := make(map[string]string)
// 	err := validate.Struct(dataset)

// 	if err != nil {

// 		if err, ok := err.(*validator.InvalidValidationError); ok {
// 			return err, false, errors
// 		}
// 		datasetPtr := reflect.ValueOf(dataset)
// 		datasetVal := reflect.Indirect(datasetPtr)
// 		datasetType := datasetVal.Type()

// 		for _, Valerr := range err.(validator.ValidationErrors) {
// 			field, _ := datasetType.FieldByName(Valerr.StructField())
// 			name := Valerr.StructField()
// 			errors[name] = field.Tag.Get("valerr")
// 		}

// 		return nil, false, errors
// 	}
// 	return nil, true, errors
// }

// func ValidatePassword(password string) map[string]string {
// 	var uppercasePresent bool
// 	var lowercasePresent bool
// 	var numberPresent bool
// 	var specialCharPresent bool
// 	const minPassLength = 8
// 	const maxPassLength = 64
// 	var passLen int
// 	errorList := make(map[string]string)

// 	for _, ch := range password {
// 		switch {
// 		case unicode.IsNumber(ch):
// 			numberPresent = true
// 			passLen++
// 		case unicode.IsUpper(ch):
// 			uppercasePresent = true
// 			passLen++
// 		case unicode.IsLower(ch):
// 			lowercasePresent = true
// 			passLen++
// 		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
// 			specialCharPresent = true
// 			passLen++
// 		case ch == ' ':
// 			passLen++
// 		}
// 	}

// 	if !lowercasePresent {
// 		errorList["lowercase"] = "lowercase letter missing"
// 	}
// 	if !uppercasePresent {
// 		errorList["uppercase"] = "uppercase letter missing"
// 	}
// 	if !numberPresent {
// 		errorList["number"] = "atleast one numeric character required"
// 	}
// 	if !specialCharPresent {
// 		errorList["specialChar"] = "special character missing"
// 	}
// 	if !(minPassLength <= passLen && passLen <= maxPassLength) {
// 		errorList["minlength"] = fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength)
// 	}

// 	return errorList
// }

type ApiResponse struct {
	Success bool                   `json:"success"`
	Status  int                    `json:"statusCode"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Errors  map[string]string      `json:"errors,omitempty"`
}
