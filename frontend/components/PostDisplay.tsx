import Image from "next/image";

import styles from "../src/styles/PostDisplay.module.css"

export default function PostDisplay(props: {post: any}) {
    return (
        <div className={styles.postContainer}>
            <Image
                src={"http://localhost:7100/" + props.post.imgPath}
                alt={"Missing Post Image"}
                width={100}
                height={100}
                className={styles.postImage}
            />
            <p>{props.post.caption}</p>
        </div>
    )
}

