import React from 'react';

import styles from "../styles/Home.module.css";

import Head from "next/head"
import Router from "next/router";
import {GetServerSidePropsContext} from "next";

export default function Home() {
  return (
      <>
        <Head>
          <title>FoodHub</title>
        </Head>
        <div className={styles.container}>
          <div className={styles.content}>
              <h1>Welcome to FoodHub</h1>
              <h3>The world&apos;s best social media site for food.</h3>
          </div>
          <div className={styles.spacer}></div>
          <div className={styles.authLinks}>
              <button
                  className={styles.loginBtn}
                  onClick={() => Router.push("/auth/login")}>
                  Login
              </button>
              <p>or</p>
              <button
                  className={styles.signupBtn}
                  onClick={() => Router.push("auth/signup")}>
                  Sign Up
              </button>
          </div>
        </div>
      </>
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