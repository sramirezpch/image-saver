import express from "express";
import aws from "aws-sdk";

import ImageController from "./src/image/controller/controller.js";
import ImageService from "./src/image/service/service.js";

aws.config.update({
  region: "us-east-1",
  credentials: {
    accessKeyId: "",
    secretAccessKey: "",
  },
});

const app = express();
const port = 8080;

const imageService = new ImageService(app);
const imageController = new ImageController(imageService);
imageController.defineHandlers();

app.use("/image", imageController.router);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
