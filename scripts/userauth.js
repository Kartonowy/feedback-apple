import "dotenv/config";
import { Jwt } from "jsonwebtoken";
import express from "express";
import bcrypt from "bcrypt";

const auth = express.Router();

function verifyJWT(req, res, next) {
  const token = req.header["x-access-token"]?.split(" ")[1];

  if (token) {
    jwt.verify(token, process.env.JWT_SECRET, (err, decoded) => {
      if (err)
        return res.json({
          isLogged: false,
          message: "Failed to verify",
        });
      req.user = {};
      req.user.id = decoded.id;
      req.user.username = decoded.username;
      next;
    });
  } else {
    res.json({ message: "Incorrect token given", isLogged: false });
  }
}
