import aws from "aws-sdk";
import express from "express";

export default class ImageController {
  constructor(imageService) {
    this.s3 = new aws.S3();
    const params = {
      Bucket: "nft-metadata-images",
      ACL: "public-read",
    };

    const policy = {
      Version: "2012-10-17",
      Statement: [
        {
          Effect: "Allow",
          Principal: "*",
          Action: ["s3:GetObject"],
          Resource: ["arn:aws:s3:::nft-metadata-images/*"],
        },
      ],
    };

    this.s3.listObjects({ Bucket: "nft-metadata-images" }, (err, data) => {
      if (err.code == "NoSuchBucket")
        this.s3.createBucket(params, (err, data) => {
          if (err) {
            console.log(err);
            return;
          }
          this.s3.putBucketPolicy(
            { Bucket: "nft-metadata-images", Policy: JSON.stringify(policy) },
            (err, data) => {
              if (err) {
                console.log("Couldn't create bucket");
                console.error(err);
                return;
              }
              console.log("Bucket created successfully");
            }
          );
        });
    });

    this.router = express.Router();
    this.imageService = imageService;
  }

  defineHandlers() {
    this.router.post("/save", this.imageService.saveImage);
  }
}
