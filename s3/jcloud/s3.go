package jcloud

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"time"
)

/**
* @Author: He Bingchang
* @Date: 2020/8/27
* @Description:
 */

var (
	JCloudS3 jcloudS3
)

type jcloudS3 struct {
	AccessKey string
	SecretKey string
	EndPoint  string
	Bucket    string
	Client    *s3.S3
	Prefix    string
}

func (s *jcloudS3) Init() {
	s.EndPoint = "https://s3.jcloud.sjtu.edu.cn"
	s.AccessKey = "5da1483fe3a4418db8c3e54d0715f343"
	s.SecretKey = "3ce4c753756c4da2b51612a42bcfdc79"
	s.Bucket = "913093352c884308a37a6d700984b013-survey"

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(s.AccessKey, s.SecretKey, ""),
		Endpoint:         aws.String(s.EndPoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	s.Client = s3Client
	s.Prefix = "survey/"
}

func (s *jcloudS3) Put(data []byte, name string, ACL ...string) error {
	objName := aws.String(s.Prefix + name)
	r := bytes.NewReader(data)
	newObj := &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    objName,
		Body:   aws.ReadSeekCloser(r),
	}

	if len(ACL) > 0 {
		newObj.ACL = aws.String(ACL[0])
	}
	_, err := s.Client.PutObject(newObj)

	return err
}

func (s *jcloudS3) Get(name string) ([]byte, error) {
	objName := aws.String(s.Prefix + name)
	newObj := &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    objName,
	}
	output, err := s.Client.GetObject(newObj)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(output.Body)

	return body, err
}

func (s *jcloudS3) Delete(name string) error {
	objName := aws.String(s.Prefix + name)
	newObj := &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    objName,
	}
	_, err := s.Client.DeleteObject(newObj)
	return err
}

func (s *jcloudS3) PreSign(name string, path string, contentType *string) (string, error) {
	objName := aws.String(s.Prefix + path)
	cd := aws.String("attachment;filename=" + name)
	r, _ := s.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket:                     aws.String(s.Bucket),
		Key:                        objName,
		ResponseContentDisposition: cd,
		ResponseContentType:        contentType,
	})
	urlStr, err := r.Presign(15 * time.Minute)

	return urlStr, err
}

func (s *jcloudS3) PreSignWithInline(name string, path string, contentType *string) (string, error) {
	objName := aws.String(s.Prefix + path)
	cd := aws.String("inline;filename=" + name)
	r, _ := s.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket:                     aws.String(s.Bucket),
		Key:                        objName,
		ResponseContentDisposition: cd,
		ResponseContentType:        contentType,
	})
	urlStr, err := r.Presign(15 * time.Minute)

	return urlStr, err
}
