import {Post} from "../../../types/Post";

import Navbar from "../../../components/Navbar";

import FeedDisplay from "../../../components/FeedDisplay";
import {useEffect, useState} from "react";

export default function Feed() {
    const [posts, setPosts] = useState<any>([]);

    useEffect(() => {
        fetch("http://localhost:7100/feed", {
            method: "GET",
            credentials: "include",
        })
            .then(res => res.json())
            .then(data => {
                setPosts(data.posts);
            })
    }, [])

    return (
        <>
            <Navbar />
            {posts && <FeedDisplay posts={posts} />}
        </>
    )
}
