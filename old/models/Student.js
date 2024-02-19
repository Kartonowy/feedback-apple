import mongoose from "mongoose";

const schemaStudent = new mongoose.Schema({
    email: String,
    fullname: String,
    schoolId: Number,
    classId: String,
    rating: Number,
    restricted: Boolean,
})

const Student = mongoose.model('Student', schemaStudent)

export default Student