import React from "react"
import styles from "../../../styles/Login.module.css"

import Router from "next/router"
import {GetServerSidePropsContext} from "next";

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

            await Router.push("/feed");
        } catch (e: any) {
            console.log(e);
            Router.reload();
        }
    }

    return(
        <div className={styles.loginForm}>
            <h1>Login</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    email
                    <input type="text" id="email" name="email" placeholder="email@website.com" />
                </label>

                <label>
                    password
                    <input type="password" id="password" name="password" placeholder="**********" />
                </label>

                <button type="submit">Login</button>
            </form>

            <button
                className={styles.loginRedirect}
                onClick={() => Router.push("/auth/signup")}>
                Need an account?
            </button>
        </div>
    )
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const jwt = ctx.req.cookies.jwt;

    if (jwt) {
        return {
            redirect: {
                destination: "/feed",
                permanent: false,
            }
        }
    }

    return {
        props: {}
    }
}