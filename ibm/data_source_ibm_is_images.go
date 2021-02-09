/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	isImages = "images"
)

func dataSourceIBMISImages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISImagesRead,

		Schema: map[string]*schema.Schema{

			isImages: {
				Type:        schema.TypeList,
				Description: "List of images",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Image name",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this image",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of this image",
						},
						"visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the image is publicly visible or private to the account",
						},
						"os": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Image Operating system",
						},
						"architecture": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The operating system architecture",
						},
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this image",
						},
						isImageEncryptionKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource",
						},
						isImageEncryption: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of encryption used on the image",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISImagesRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	if userDetails.generation == 1 {
		err := classicImageList(d, meta)
		if err != nil {
			return err
		}
	} else {
		err := imageList(d, meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicImageList(d *schema.ResourceData, meta interface{}) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcclassicv1.Image{}
	for {
		listImagesOptions := &vpcclassicv1.ListImagesOptions{}
		if start != "" {
			listImagesOptions.Start = &start
		}
		availableImages, response, err := sess.ListImages(listImagesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Images %s\n%s", err, response)
		}
		start = GetNext(availableImages.Next)
		allrecs = append(allrecs, availableImages.Images...)
		if start == "" {
			break
		}
	}
	imagesInfo := make([]map[string]interface{}, 0)
	for _, image := range allrecs {

		l := map[string]interface{}{
			"name":         *image.Name,
			"id":           *image.ID,
			"status":       *image.Status,
			"crn":          *image.CRN,
			"visibility":   *image.Visibility,
			"os":           *image.OperatingSystem.Name,
			"architecture": *image.OperatingSystem.Architecture,
		}
		imagesInfo = append(imagesInfo, l)
	}
	d.SetId(dataSourceIBMISSubnetsID(d))
	d.Set(isImages, imagesInfo)
	return nil
}

func imageList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcv1.Image{}
	for {
		listImagesOptions := &vpcv1.ListImagesOptions{}
		if start != "" {
			listImagesOptions.Start = &start
		}
		availableImages, response, err := sess.ListImages(listImagesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Images %s\n%s", err, response)
		}
		start = GetNext(availableImages.Next)
		allrecs = append(allrecs, availableImages.Images...)
		if start == "" {
			break
		}
	}
	imagesInfo := make([]map[string]interface{}, 0)
	for _, image := range allrecs {

		l := map[string]interface{}{
			"name":         *image.Name,
			"id":           *image.ID,
			"status":       *image.Status,
			"crn":          *image.CRN,
			"visibility":   *image.Visibility,
			"os":           *image.OperatingSystem.Name,
			"architecture": *image.OperatingSystem.Architecture,
		}
		if image.Encryption != nil {
			l["encryption"] = *image.Encryption
		}
		if image.EncryptionKey != nil {
			l["encryption_key"] = *image.EncryptionKey.CRN
		}
		imagesInfo = append(imagesInfo, l)
	}
	d.SetId(dataSourceIBMISSubnetsID(d))
	d.Set(isImages, imagesInfo)
	return nil
}

// dataSourceIBMISImagesId returns a reasonable ID for a image list.
func dataSourceIBMISImagesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
