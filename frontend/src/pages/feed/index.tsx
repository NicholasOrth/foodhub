import {GetServerSidePropsContext} from "next";
import {Post} from "../../../types/Post";
import Image from "next/image";
import Navbar from "../../../components/Navbar";

export default function Feed(props: { data: Post[] }) {
    return (
        <>
            <Navbar />
            <div>
                {props.data.map(post => (
                    <div key={post.id}>
                        <h3>{post.caption}</h3>
                        <p>@{post.username}</p>
                        <Image
                            src={"http://localhost:7100/images" + post.imgPath}
                            alt={"Missing post image"}
                            width={500}
                            height={500}
                        />
                    </div>
                ))}
                <br />
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
        const data: Post[] = await res.json();
        return {
            props: {
                data
            }
        }
    } catch (e) {
        console.log(e);
        return {
            props: {
                posts: [],
            }
        }
    }
}



