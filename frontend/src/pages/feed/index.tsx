import {GetServerSidePropsContext} from "next";
import {Post} from "../../../types/Post";

import Navbar from "../../../components/Navbar";
import Image from "next/image";

export default function Feed(props: any) {
    return (
        <>
            <Navbar />
            <div>
                {props.data.posts.map((post: Post) => (
                    <div key={post.id}>
                        <h1>{post.caption}</h1>
                        <p>@{post.username}</p>
                        <Image
                            src={"http://localhost:7100/" + post.imgPath}
                            alt={"Missing image"}
                            width={300}
                            height={300}
                        />
                    </div>
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



