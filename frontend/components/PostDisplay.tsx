import {Post} from "../types/Post";

import styles from "../src/styles/PostDisplay.module.css";
import Image from "next/image";
import {useEffect, useState} from "react";
import Cookies from "universal-cookie";

const cookies = new Cookies();

export default function PostDisplay({post}: {post: Post}) {
    const [likes, setLikes] = useState(post.likes || 0);

    const likePost = async () => {
        const res = await fetch("http://localhost:7100/post/like/" + post.id, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Cookie": "jwt=" + cookies.get("jwt"),
            },
            credentials: "include",
        });
        const data = await res.json();
        setLikes(data.likes);
    }

    useEffect(() => {

    }, [likes, setLikes])

    return (
        <div className={styles.postContainer}>
            <h3>{post.caption}</h3>
            <p>@{post.username}</p>
            <Image
                src={"http://localhost:7100/" + post.imgPath}
                alt={"Missing image"}
                width={300}
                height={300}
            />
            <button
                className={styles.like}
                onClick={likePost}
            >
                {likes || 0}<p>❤</p>
            </button>
        </div>
    )
}