import {GetServerSidePropsContext} from "next";
import {Post} from "../../../types/Post";

import Navbar from "../../../components/Navbar";

import FeedDisplay from "../../../components/FeedDisplay";

export default function Feed(props: any) {
    return (
        <>
            <Navbar />
            <FeedDisplay posts={props.data.posts} />
        </>
    )
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const jwt = ctx.req.cookies.jwt;

    if (!jwt) {
        return {
            redirect: {
                destination: "/",
                permanent: false,
            }
        }
    }

    const res: Response = await fetch("http://localhost:7100/feed", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Cookie": "jwt=" + jwt,
        },
        credentials: "include",
    })

    try {
        const data: { posts: Post[] } = await res.json();
        return {
            props: {
                data
            }
        }
    } catch (e) {
        return {
            props: {
                posts: [],
            }
        }
    }
}



