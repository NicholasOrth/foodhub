import React, {useState} from "react"

import Router from "next/router"

export default function Login() {
    const[email, setEmail] = useState('');
    const[password,setPassword] = useState('');

    const handleSubmit = (e: any)=>{
        e.preventDefault();
        console.log(email);
    }

    return(
        <div className="authform">
         <form onSubmit={handleSubmit}>
            <label htmlFor="email">email</label>
            <input value={email} onChange={(e) => setEmail(e.target.value)} type ="text" placeholder="email@whatevermail.com" id="email" name="email"/>
            <label htmlFor="password">email</label>
            <input value={password} onChange={(e) => setPassword(e.target.value)} type="passsword" placeholder="**********" id="password" name="name"/>
            <button> Log In</button>
         </form>

         <button onClick={
             () => Router.push("/auth/signup")}
         >
             Dont Have an account? Register here.
         </button>

        </div>
    )
   
} 