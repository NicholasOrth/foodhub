import React, { useEffect} from 'react';
import Router from "next/router"

export default function home() {
  useEffect(() => {
    document.title = "FoodHub";
  }, []);

  return (
      <>
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