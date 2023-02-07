import React, {useState} from "react"

import Router from "next/router"

export default function Login() {
    const handleSubmit = (e: any)=>{
        e.preventDefault();

        const data = {
            email: e.target.email.value,
            password: e.target.password.value,
        };
    }

    return(
        <div className="authform">
         <form onSubmit={handleSubmit}>
            <label htmlFor="email">email</label>
            <input type="text" placeholder="email@whatevermail.com" id="email" name="email"/>
             <br />

            <label htmlFor="password">password</label>
            <input type="passsword" placeholder="**********" id="password" name="name"/>
             <br/>

            <button> Log In</button>
         </form>

        <br/>

         <button onClick={
             () => Router.push("/auth/signup")}
         >
             Dont Have an account? Register here.
         </button>

        </div>
    )
   
} 