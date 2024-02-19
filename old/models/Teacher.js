import mongoose from "mongoose";

const schemaTeacher = new mongoose.Schema({
    email: String,
    fullname: String,
    schoolId: Number,
    subject: String,
    messages: [String],
})

const Teacher = mongoose.model('Teacher', schemaTeacher)

export default Teacher