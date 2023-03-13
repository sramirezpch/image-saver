import {
  S3Client,
  PutObjectCommand,
  PutObjectCommandInput,
} from "@aws-sdk/client-s3";
import { Request, Response } from "express";

export default class ImageService {
  private s3: S3Client;

  constructor() {
    this.s3 = new S3Client({ region: "us-east-1" });
  }

  async storeImage(
    req: Request<{}, {}, { key: string; body: string }>,
    res: Response
  ) {
    const { key, body } = req.body;
    const objectParams: PutObjectCommandInput = {
      Bucket: "nft-image-saver",
      Body: body,
      Key: key,
    };

    const command = new PutObjectCommand(objectParams);
    await this.s3.send(command);

    res.send("ok");
  }
}
