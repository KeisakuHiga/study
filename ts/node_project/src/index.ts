import dotenv from "dotenv";
import express from "express";

dotenv.config();

const app = express();
const port = process.env.PORT;
app.get("/", (_, res) => {
  res.send("The sedulous hyena ate the antelope!");
});
app.listen(port, () => {
  return console.log(`server is listening on ${port}!!`);
});
