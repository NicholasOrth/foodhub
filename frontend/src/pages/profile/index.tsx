import {GetServerSidePropsContext} from "next";

import Image from "next/image";
import Router from "next/router";

export default function Profile(
    props: {data: {email: string, name: string, posts: any[]}})
{
    return (
        <div>
            <h1>Profile</h1>
            <p>{props.data.name}</p>
            <p>{props.data.email}</p>
            {
                props.data.posts.map((post: any) => {
                    return (
                        <div key={post.ID}>
                            <p>{post.caption}</p>
                            <Image
                                src={"http://localhost:7100/" + post.imgPath}
                                alt={"Missing Post Image"}
                                width={100}
                                height={100}
                            />
                        </div>
                    )
                })
            }
            <button onClick={() => Router.push("/new")}>
                Create New Post
            </button>
        </div>
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
