import "dotenv/config"
import mongoose from "mongoose"

const connectToDB = async () => {
    try {
        mongoose.connect(process.env.MONGO_URL, 
            { useNewUrlParser: true}
        )
        console.log("MongoDB connected successfully!")
    } catch (err) {
        console.log(err.message)
        process.exit(1)
    }
}

export default connectToDB