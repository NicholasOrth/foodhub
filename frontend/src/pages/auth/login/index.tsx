import React, {useState} from "react"


export const Index = (props: any) => {
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
        <input value={email} onChange={(e) => setEmail(e.target.value)} type ="email" placeholder="email@whatevermail.com" id="email" name="email"/>
        <label htmlFor="password">email</label>
        <input value={password} onChange={(e) => setPassword(e.target.value)} type="passsword" placeholder="**********" id="password" name="name"/>
        <button> Log In</button>
     </form>
     <button onClick={() => props.onFormSwitch('register')}>Dont Have an account? Register here. </button>
     </div>
    )
   
} 