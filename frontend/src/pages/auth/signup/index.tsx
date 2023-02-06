import React from "react";

import Router from "next/router";

import styles from "@/styles/Signup.module.css"

export default function SignUp() {

    const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const data = {
            name: e.currentTarget.username.value,
            email: e.currentTarget.email.value,
            password: e.currentTarget.password.value,
        }

        if (data.password !== e.currentTarget.password2.value) {
            alert('Password must match');
            Router.reload();
            return;
        }

        const res: Response = await fetch("http://localhost:8080/user", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        });

        if (res.status === 200) {
            await Router.push('/auth/login');
            alert('Account created successfully');
        } else {
            Router.reload();
            alert('Something went wrong');
        }
    }

    return (
        <div className={styles.pageContainer}>
            <h1>Sign Up</h1>
            <form onSubmit={onSubmit}>
                <div className={styles.fieldContainer}>
                    <label htmlFor="username">Username</label>
                    <input type="text" name="username" id="username" required />
                </div>

                <div className={styles.fieldContainer}>
                    <label htmlFor="email">Email</label>
                    <input type="text" name="email" id="email" required />
                </div>

                <div className={styles.fieldContainer}>
                    <label htmlFor="password">Password</label>
                    <input type="password" name="password" id="password" required />
                </div>

                <div className={styles.fieldContainer}>
                    <label htmlFor="password2">Confirm Password</label>
                    <input type="password" name="password2" id="password2" required />
                </div>

                <div className={styles.fieldContainer}>
                    <button type="submit">Sign Up</button>
                </div>
            </form>
        </div>
    )
}