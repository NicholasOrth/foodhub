import {GetServerSidePropsContext} from "next";
import {Post} from "../../../types/Post";

import Navbar from "../../../components/Navbar";

import PostDisplay from "../../../components/PostDisplay";

import styles from "../../styles/Feed.module.css";

export default function Feed(props: any) {
    return (
        <>
            <Navbar />
            <div className={styles.container}>
                {props.data.posts.map((post: Post) => (
                    <PostDisplay post={post} key={post.id} />
                ))}
            </div>
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



