import {Post} from "../types/Post";

import styles from "@/styles/Feed.module.css";

import PostDisplay from "./PostDisplay";

export default function FollowerDisplay({followers}: { followers: Post[] }) {
    return (
        <div className={styles.container}>
            {followers.map((post: Post) => (
                <PostDisplay post={post} key={post.id} />
            ))}
        </div>
    )
}
