import { S3 } from "aws-sdk";
import { Request, Response } from "express";

export default class ImageService {
  private s3: any;

  constructor() {
    this.s3 = new S3();
  }

  public saveImage(req: Request, res: Response) {
    res.send("POST /image/save");
  }
}
