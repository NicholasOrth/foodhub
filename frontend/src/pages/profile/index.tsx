import {GetServerSidePropsContext} from "next";

import Router from "next/router";

import styles from "../../styles/Profile.module.css";
import Navbar from "../../../components/Navbar";

export default function Profile(
    props: {data: {email: string, name: string, posts: any[]}})
{
    return (
        <>
            <Navbar />
            <div className={styles.profileContainer}>
                <h1>{props.data.name}</h1>
                <button onClick={() => Router.push("/new")}>
                    Create New Post
                </button>
            </div>
        </>
    )
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const jwt = ctx.req.cookies.jwt;

    if (!jwt) {
        return {
            redirect: {
                destination: "/auth/login",
                permanent: false,
            }
        }
    }

    const res: Response = await fetch('http://localhost:7100/user', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Cookie': "jwt=" + jwt,
        },
        credentials: 'include',
    })

    const data: {email: string, name: string, posts: any[]} = await res.json();

    return {
        props: {
            data
        }
    }
}
