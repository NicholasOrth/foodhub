import {GetServerSidePropsContext} from "next";

import Navbar from "../../../components/Navbar";
import FeedDisplay from "../../../components/FeedDisplay";
import Router from "next/router";

export default function Profile(
    props: {data: {email: string, name: string, posts: any[]}})
{
    const logout = async () => {
        await fetch("http://localhost:7100/auth/logout", {
            method: "POST",
            credentials: "include",
        })

        await Router.push("/auth/login");
    };

    return (
        <>
            <Navbar />
            <FeedDisplay posts={props.data.posts} />
            <button onClick={() => { logout(); }}>Logout</button>
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

    const res: Response = await fetch('http://localhost:7100/user/me', {
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
