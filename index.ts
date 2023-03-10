import AppServer from "./src/app-server";
import ImageService from "./src/image/service";
import ImageController from "./src/image/controller";

const imageService = new ImageService();
const imageController = new ImageController(imageService, "/image");

const app = new AppServer([
  {
    handlers: imageController.defineHandlers(),
    prefix: imageController.prefix,
  },
]);

app.start();
