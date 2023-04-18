import Router from "next/router"
import toast from "react-hot-toast"

import styles from "../../../styles/Signup.module.css"

export default function Signup() {

    const handleSubmit = async (e: any) => {
        e.preventDefault();


        const name= e.target.name.value;
        const email= e.target.email.value;
         const  password= e.target.password.value;
         const confirm = e.target.confirm.value;
         const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d).{8,}$/;
         const emailRegex = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/;

         if(!emailRegex.test(email)) {
              toast.error("Invalid email");
                return;
            }
        
         if (!passwordRegex.test(password)) {
           toast.error("Password must contain at least 8 characters, including at least one uppercase letter, one lowercase letter, and one number.");
           return;
         }
       
         if (password !== confirm) {
           toast.error("Passwords do not match");
           return;
         }
       
         const data = { name, email, password };
       

        const res: Response = await fetch("http://localhost:7100/auth/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });
        

        if (res.status === 200) {
            toast.success("Sign up successful");
            await Router.push("/auth/login");
        }
    }

    return(
        <div className={styles.signupForm}>
            <h1>Sign Up</h1>
            <form onSubmit={handleSubmit}>
                <label>
                   name
                    <input type="text" id="name" name="name" placeholder="name"/>
                </label>

                <label htmlFor="email">
                    email
                    <input type="text" id="email" name="email" placeholder="email@website.com" />
                </label>

                <label>
                    password
                    <input type="password" id="password" name="password" placeholder="**********" />
                </label>

                <label>
                    confirm password
                    <input type="password" id="confirm" name="confirm" placeholder="**********" />
                </label>

                <button type="submit">Sign Up</button>
            </form>

            <button
                className={styles.loginRedirect}
                onClick={() => Router.push("/auth/login")}>
                Already have an account?
            </button>
        </div>
    )
}