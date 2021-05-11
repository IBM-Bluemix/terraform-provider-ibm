// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/awserr"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	token "github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam/token"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMCOSObject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCOSObjectCreate,
		ReadContext:   resourceIBMCOSObjectRead,
		UpdateContext: resourceIBMCOSObjectUpdate,
		DeleteContext: resourceIBMCOSObjectDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "COS object body",
			},
			"bucket_crn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "COS bucket CRN",
			},
			"bucket_location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "COS bucket location",
			},
			"content": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content_base64", "content_file"},
				Description:   "COS object content",
			},
			"content_base64": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content", "content_file"},
				Description:   "COS object content in base64 encoding",
			},
			"content_file": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content", "content_base64"},
				Description:   "COS object content file path",
			},
			"content_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "COS object content length",
			},
			"content_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "COS object content type",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"public", "private", "direct"}),
				Description:  "COS endpoint type: public, private, direct",
				Default:      "public",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "COS object MD5 hexdigest",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "COS object key",
			},
			"last_modified": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "COS object last modified date",
			},
		},
	}
}

func resourceIBMCOSObjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	bucketCRN := d.Get("bucket_crn").(string)
	bucketName := strings.Split(bucketCRN, ":bucket:")[1]
	instanceCRN := fmt.Sprintf("%s::", strings.Split(bucketCRN, ":bucket:")[0])

	bucketLocation := d.Get("bucket_location").(string)
	endpointType := d.Get("endpoint_type").(string)

	bxSession, err := m.(ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}

	s3Client, err := getS3Client(bxSession, bucketLocation, endpointType, instanceCRN)
	if err != nil {
		return diag.FromErr(err)
	}

	objectKey := d.Get("key").(string)

	// This check is to make sure new create does not
	// overwrite objects that is not managed by Terraform
	exists, err := objectExists(s3Client, bucketName, objectKey)
	if err != nil {
		return diag.FromErr(err)
	}
	if exists {
		return diag.FromErr(fmt.Errorf("error COS bucket (%s) object (%s) already exists", bucketName, objectKey))
	}

	var body io.ReadSeeker

	if v, ok := d.GetOk("content"); ok {
		content := v.(string)
		body = bytes.NewReader([]byte(content))
	} else if v, ok := d.GetOk("content_base64"); ok {
		content := v.(string)
		contentRaw, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error decoding content_base64: %s", err))
		}
		body = bytes.NewReader(contentRaw)
	} else if v, ok := d.GetOk("content_file"); ok {
		path := v.(string)
		file, err := os.Open(path)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error opening COS object file (%s): %s", path, err))
		}

		body = file
		defer func() {
			err := file.Close()
			if err != nil {
				log.Printf("[WARN] Failed closing COS object file (%s): %s", path, err)
			}
		}()
	}

	putInput := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   body,
	}

	if _, err := s3Client.PutObject(putInput); err != nil {
		return diag.FromErr(fmt.Errorf("error putting object (%s) in COS bucket (%s): %s", objectKey, bucketName, err))
	}

	objectID := getObjectId(bucketCRN, objectKey, bucketLocation)
	d.SetId(objectID)

	return resourceIBMCOSObjectRead(ctx, d, m)
}

func resourceIBMCOSObjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	objectID := d.Id()

	bucketCRN := parseObjectId(objectID, "bucketCRN")
	bucketName := parseObjectId(objectID, "bucketName")
	bucketLocation := parseObjectId(objectID, "bucketLocation")
	instanceCRN := parseObjectId(objectID, "instanceCRN")
	endpointType := d.Get("endpoint_type").(string)

	d.Set("bucket_crn", bucketCRN)
	d.Set("bucket_location", bucketLocation)

	bxSession, err := m.(ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}

	s3Client, err := getS3Client(bxSession, bucketLocation, endpointType, instanceCRN)
	if err != nil {
		return diag.FromErr(err)
	}

	objectKey := parseObjectId(objectID, "objectKey")
	headInput := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	out, err := s3Client.HeadObject(headInput)
	if err != nil {
		d.SetId("") // Set state back to empty for terraform refresh
		return diag.FromErr(fmt.Errorf("failed getting COS bucket (%s) object (%s): %w", bucketName, objectKey, err))
	}

	log.Printf("[DEBUG] Received COS object: %s", out)

	d.Set("content_length", out.ContentLength)
	d.Set("content_type", out.ContentType)
	d.Set("etag", strings.Trim(aws.StringValue(out.ETag), `"`))
	if out.LastModified != nil {
		d.Set("last_modified", out.LastModified.Format(time.RFC1123))
	} else {
		d.Set("last_modified", "")
	}

	if isContentTypeAllowed(out.ContentType) {
		getInput := s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
		}
		out, err := s3Client.GetObject(&getInput)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed getting COS object: %w", err))
		}

		buf := new(bytes.Buffer)
		bytesRead, err := buf.ReadFrom(out.Body)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed reading content of COS bucket (%s) object (%s): %w", bucketName, objectKey, err))
		}
		log.Printf("[INFO] Saving %d bytes from COS bucket (%s) object (%s)", bytesRead, bucketName, objectKey)
		d.Set("body", buf.String())
	} else {
		contentType := ""
		if out.ContentType == nil {
			contentType = "<EMPTY>"
		} else {
			contentType = aws.StringValue(out.ContentType)
		}

		log.Printf("[INFO] Ignoring body of COS bucket (%s) object (%s) with Content-Type %q", bucketName, objectKey, contentType)
	}

	d.Set("key", objectKey)

	return nil
}

func resourceIBMCOSObjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChanges("content", "content_base64", "content_file", "etag") {
		bucketCRN := d.Get("bucket_crn").(string)
		bucketName := strings.Split(bucketCRN, ":bucket:")[1]
		instanceCRN := fmt.Sprintf("%s::", strings.Split(bucketCRN, ":bucket:")[0])

		bucketLocation := d.Get("bucket_location").(string)
		endpointType := d.Get("endpoint_type").(string)

		bxSession, err := m.(ClientSession).BluemixSession()
		if err != nil {
			return diag.FromErr(err)
		}

		s3Client, err := getS3Client(bxSession, bucketLocation, endpointType, instanceCRN)
		if err != nil {
			return diag.FromErr(err)
		}

		var body io.ReadSeeker

		if v, ok := d.GetOk("content"); ok {
			content := v.(string)
			body = bytes.NewReader([]byte(content))
		} else if v, ok := d.GetOk("content_base64"); ok {
			content := v.(string)
			contentRaw, err := base64.StdEncoding.DecodeString(content)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error decoding content_base64: %s", err))
			}
			body = bytes.NewReader(contentRaw)
		} else if v, ok := d.GetOk("content_file"); ok {
			path := v.(string)
			file, err := os.Open(path)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error opening COS object file (%s): %s", path, err))
			}

			body = file
			defer func() {
				err := file.Close()
				if err != nil {
					log.Printf("[WARN] Failed closing COS object file (%s): %s", path, err)
				}
			}()
		}

		objectKey := d.Get("key").(string)

		putInput := &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   body,
		}

		if _, err := s3Client.PutObject(putInput); err != nil {
			return diag.FromErr(fmt.Errorf("error putting object (%s) in COS bucket (%s): %s", objectKey, bucketName, err))
		}

		objectID := getObjectId(bucketCRN, objectKey, bucketLocation)
		d.SetId(objectID)
	}

	return resourceIBMCOSObjectRead(ctx, d, m)
}

func resourceIBMCOSObjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	bucketCRN := d.Get("bucket_crn").(string)
	bucketName := strings.Split(bucketCRN, ":bucket:")[1]
	instanceCRN := fmt.Sprintf("%s::", strings.Split(bucketCRN, ":bucket:")[0])

	bucketLocation := d.Get("bucket_location").(string)
	endpointType := d.Get("endpoint_type").(string)

	bxSession, err := m.(ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}

	s3Client, err := getS3Client(bxSession, bucketLocation, endpointType, instanceCRN)
	if err != nil {
		return diag.FromErr(err)
	}

	objectKey := d.Get("key").(string)

	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	_, err = s3Client.DeleteObject(deleteInput)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func getCosEndpoint(bucketLocation string, endpointType string) string {
	if bucketLocation != "" {
		switch endpointType {
		case "public":
			return fmt.Sprintf("s3.%s.cloud-object-storage.appdomain.cloud", bucketLocation)
		case "private":
			return fmt.Sprintf("s3.private.%s.cloud-object-storage.appdomain.cloud", bucketLocation)
		case "direct":
			return fmt.Sprintf("s3.direct.%s.cloud-object-storage.appdomain.cloud", bucketLocation)
		default:
			return fmt.Sprintf("s3.%s.cloud-object-storage.appdomain.cloud", bucketLocation)
		}
	}
	return ""
}

func getS3Client(bxSession *bxsession.Session, bucketLocation string, endpointType string, instanceCRN string) (*s3.S3, error) {
	var s3Conf *aws.Config

	apiEndpoint := getCosEndpoint(bucketLocation, endpointType)
	apiEndpoint = envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)
	if apiEndpoint == "" {
		return nil, fmt.Errorf("the endpoint doesn't exists for given location %s and endpoint type %s", bucketLocation, endpointType)
	}

	authEndpoint, err := bxSession.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		return nil, err
	}
	authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")
	apiKey := bxSession.Config.BluemixAPIKey
	if apiKey != "" {
		s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, instanceCRN)).WithS3ForcePathStyle(true)
	}
	iamAccessToken := bxSession.Config.IAMAccessToken
	if iamAccessToken != "" {
		initFunc := func() (*token.Token, error) {
			return &token.Token{
				AccessToken:  bxSession.Config.IAMAccessToken,
				RefreshToken: bxSession.Config.IAMRefreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
				Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
			}, nil
		}
		s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, instanceCRN)).WithS3ForcePathStyle(true)
	}
	s3Sess := session.Must(session.NewSession())
	return s3.New(s3Sess, s3Conf), nil
}

// This is to prevent potential issues w/ binary files
// and generally unprintable characters
// See https://github.com/hashicorp/terraform/pull/3858#issuecomment-156856738
func isContentTypeAllowed(contentType *string) bool {
	if contentType == nil {
		return false
	}

	allowedContentTypes := []*regexp.Regexp{
		regexp.MustCompile("^text/.+"),
		regexp.MustCompile("^application/json$"),
	}

	for _, r := range allowedContentTypes {
		if r.MatchString(*contentType) {
			return true
		}
	}

	return false
}

func objectExists(s3Client *s3.S3, bucketName string, objectKey string) (bool, error) {
	headInput := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}
	_, err := s3Client.HeadObject(headInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func getObjectId(bucketCRN string, objectKey string, bucketLocation string) string {
	return fmt.Sprintf("%s:object:%s:location:%s", bucketCRN, objectKey, bucketLocation)
}

func parseObjectId(id string, info string) string {
	splitID := strings.Split(id, ":object:")
	bucketCRN := splitID[0]

	if info == "instanceCRN" {
		return fmt.Sprintf("%s::", strings.Split(bucketCRN, ":bucket:")[0])
	}

	if info == "bucketCRN" {
		return bucketCRN
	}

	if info == "bucketName" {
		return strings.Split(bucketCRN, ":bucket:")[1]
	}

	if info == "objectKey" {
		return strings.Split(splitID[1], ":location:")[0]
	}

	if info == "bucketLocation" {
		return strings.Split(splitID[1], ":location:")[1]
	}

	return parseBucketId(splitID[0], info)
}
