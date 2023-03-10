import Router from "next/router"

import styles from "../../../styles/Signup.module.css"
import {GetServerSidePropsContext} from "next";

export default function Register() {

    const handleSubmit = async (e: any) => {
        e.preventDefault();

        const data = {
            name: e.target.name.value,
            email: e.target.email.value,
            password: e.target.password.value,
        };

        if (data.password !== e.target.confirm.value) {
            alert("Passwords do not match");
            return
        }

        const res: Response = await fetch("http://localhost:7100/auth/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        if (res.status === 200) {
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

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const jwt = ctx.req.cookies.jwt;

    if (jwt) {
        return {
            redirect: {
                destination: "/profile",
                permanent: false,
            }
        }
    }

    return {
        props: {}
    }
}