import express from "express";
import mongoose from "mongoose";
import "dotenv/config";
import connectToDB from "./scripts/connectToDB.js";
import path from "path";
import Student from "./models/Student.js";
import bodyParser from "body-parser";
import url from "url";
import Teacher from "./models/Teacher.js";

//express router initialization 
const __dirname = url.fileURLToPath(new URL(".", import.meta.url));
const app = express();
const urlencodedParser = bodyParser.urlencoded({ extended: false });

app.use(bodyParser.json(), urlencodedParser);

app.use(express.json());
app.use(express.urlencoded());

connectToDB();

app.post("/fsubmit", async (req, res) => {
  console.log(req.body);
  const teachId = Number(req.body.teacherselect);
  await Teacher.findOne({ teacherId: teachId }).then((t) => {
    t.messages.push(req.body.review);
    t.save();
  });
});

app.get("/ui", async (req, res) => {
  res.sendFile(path.join(__dirname + "/public/index.html"));
});

app.get("/style.css", (req, res) => {
  res.sendFile(path.join(__dirname + "/public/style.css"));
});

app.get("/onsite.js", (req, res) => {
  res.sendFile(path.join(__dirname + "/public/onsite.js"));
});

app.get("/add", async (req, res) => {
  const examplestd = {
    email: "dziki.wariacik@gmail.com",
    fullname: "dziku",
    schoolId: 2,
    classId: "6x",
    rating: 8,
    restricted: true,
  };
  const emailTaken = await Student.findOne({ email: examplestd.email });
  if (emailTaken) {
    console.log("Email already registered");
  } else {
    Student.create(examplestd);
  }
  res.send({ examplestd });
});

app.get("/createTeacher", async (req, res) => {
  const exampleteacher = {
    teacherId: 8,
    email: "mimikyu@mockup.xd",
    fullname: "disguise",
    schoolId: 2,
    subject: "cosplay",
    messages: [],
  };
  const emailTaken = await Teacher.findOne({ email: exampleteacher.email });
  if (emailTaken) {
    console.log("Email already registered");
  } else {
    Teacher.create(exampleteacher);
  }
  res.send({ exampleteacher });
});
//TODO: Test html form
//web token
//user auth

app.get("/", (req, res) => {
  res.send("Hello world");
});

app.listen(3000, () => {
  console.log("Listening on port 3000");
});
