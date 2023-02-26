import React from "react"
import styles from "../../../styles/Login.module.css"

import Router from "next/router"

export default function Login() {
    const handleSubmit = async (e: any)=> {
        e.preventDefault();

        const data = {
            email: e.target.email.value,
            password: e.target.password.value,
        };

        try {
            await fetch("http://localhost:7100/auth/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify(data),
            });

            await Router.push("/");
        } catch (e: any) {
            console.log(e);
            Router.reload();
        }

    }

    return(
        <div className={styles.authform}>
         <form onSubmit={handleSubmit}>
            <label htmlFor="email">email</label>
            <input type="text" placeholder="email@whatevermail.com" id="email" name="email"/>
             <br />

            <label htmlFor="password">password</label>
            <input type="password" placeholder="**********" id="password" name="password"/>
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