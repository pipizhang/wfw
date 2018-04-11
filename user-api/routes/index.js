'use strick'

import HomeController from '../controller/home'

export default app => {

  app.get("/", HomeController.index)

}
