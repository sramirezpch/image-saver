import { Router } from "express";

import BaseController from "../../base-controller";
import ImageService from "../service";

export default class ImageController extends BaseController {
  private router: Router;
  private imageService: ImageService;

  constructor(readonly _imageService: ImageService, readonly prefix: string) {
    super(prefix);
    this.imageService = _imageService;
    this.router = Router();
  }

  public defineHandlers(): Router {
    super.defineHandlers();
    this.router.post("/save", this.imageService.storeImage);

    return this.router;
  }
}
