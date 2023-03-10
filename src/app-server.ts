import express, { Express, Router } from "express";

export default class AppServer {
  private app: Express;

  constructor(controllers: any[]) {
    this.app = express();

    for (const controller of controllers) {
      this.app.use(controller.prefix, controller.handlers);
    }
  }

  public start() {
    this.app.listen(process.env.PORT || 8080, () => {
      console.log("Server is up!");
    });
  }
}
