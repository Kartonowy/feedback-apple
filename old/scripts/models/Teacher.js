import mongoose from "mongoose";

const schemaTeacher = new mongoose.Schema({
  teacherId: Number,
  email: String,
  fullname: String,
  schoolId: Number,
  subject: String,
  messages: [String],
});

const Teacher = mongoose.model("Teacher", schemaTeacher, "teachers");

export default Teacher;
