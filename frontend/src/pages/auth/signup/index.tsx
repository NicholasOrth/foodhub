import React, {useState} from "react"

import Router from "next/router"

export default function Register() {
    const[email, setEmail] = useState('');
    const[password,setPassword] = useState('');
    const[name, setName]= useState('');

    const handleSubmit = (e: any)=>{
        e.preventDefault();
        console.log(email);
    }

    return(
        <div className="authform">
            <form onSubmit={handleSubmit}>
                <label htmlFor="name">Full Name</label>
                <input value ={name} name="name" id="name" placeholder="Full Name"/>
                <label htmlFor="email">email</label>
                <input value={email} onChange={(e) => setEmail(e.target.value)} type="text" placeholder="email@whatevermail.com" id="email" name="email"/>
                <label htmlFor="password">email</label>
                <input value={password} onChange={(e) => setPassword(e.target.value)} type="passsword" placeholder="**********" id="password" name="name"/>
                <button type="submit">Log In</button>
            </form>
            <button onClick={() => Router.push("/auth/login")}> Already have an account? Login here. </button>
        </div>
    )
}