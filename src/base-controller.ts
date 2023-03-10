export default class BaseController {
  prefix: string;

  constructor(readonly _prefix: string) {
    this.prefix = _prefix;
  }

  defineHandlers() {
    console.log(`Defining route handlers for controller at ${this.prefix}`);
  }
}
