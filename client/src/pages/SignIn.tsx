<<<<<<< HEAD
import React from 'react'
import Inputform from '../components/Inputform'
import { useState } from 'react';
const SignIn = () => 
{
  const[username,setusername]=useState("")
  const[password,setpassword]=useState("")

  const submit=(e:any)=>
  {
    e.preventDefault();
    console.log("username:" + username);
    console.log("password:" + password);
  }
  const passwordhide=(e:any)=>
  {

  }

  return (
    <form onSubmit = {submit}>
    <text className = "p"> Username</text>
    <div>
    <input type = "text" className = "container" 
    onChange = {(event)=>setusername(event.target.value)}></input>
    </div>
    <text className = "p">Password</text>
    <div>
    <input type = "text" className = "container"
    onChange = {(event)=>{setpassword(event.target.value)}}></input>
    </div>
    <button className = "button" 
    type = "submit">submit</button>
   </form>
  )
}
export default SignIn
=======
import { Link } from "react-router-dom";

const SignIn = () => {
  return (
    <div className="mt-14 mx-auto w-full  bg-white rounded-lg shadow sm:max-w-md xl:p-0">
      Sign in
    </div>
  );
};

export default SignIn;
>>>>>>> ab2283133e606e0df77b9cc9e74b6e4beffaecc8
