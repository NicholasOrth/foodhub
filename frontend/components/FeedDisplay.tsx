import {Post} from "../types/Post";

import styles from "@/styles/Feed.module.css";

import PostDisplay from "./PostDisplay";

export default function FeedDisplay({posts}: { posts: Post[] }) {
    return (
        <div className={styles.container}>
            {posts.map((post: Post) => (
                <PostDisplay post={post} key={post.id} />
            ))}
        </div>
    )
}