import React from 'react';

import Router from "next/router"

import Head from "next/head"

export default function home() {
  return (
      <>
        <Head>
          <title>FoodHub</title>
        </Head>
        <div>
          <br />
          <p> FoodHub </p>
          <br />

          <button onClick={
            () => Router.push("/auth/login")}
          >
            Log In
          </button>
          <br />

          <button onClick={
            () => Router.push("/auth/signup")}
          >
            Sign Up
          </button>
          <br />

        </div>
      </>

  )

}