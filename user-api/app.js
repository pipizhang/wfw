import express from 'express'
import router from './routes/index'

const app = express()
router(app)

app.listen(8083, () => {
  console.log("app listening on port 8083")
});

