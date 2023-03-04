import express from "express";

const app = express();
const port = 8080;

app.post("/save-image", (req, res) => {
  res.json({ message: "Request received!" });
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
