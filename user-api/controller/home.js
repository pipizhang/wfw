'use strick'

import BaseController from './base'

class HomeController extends BaseController {

  constructor() {
    super()
  }

  async index(req, res, next) {
    res.send("It works")
  }

}

export default new HomeController
