
import Image from "next/image";

import styles from "../src/styles/ProfileView.module.css"

export default function ProfileView(
    props: {
        name: string,
        followers: number,
        following: number,
        postCount: number
    }
) {
    return (
        <div className={styles.container}>
            <div className={styles.card}>
                <div className={styles.imageContainer}>
                    <Image src="http://localhost:7100/images/test.jpg" alt="profile" width={300} height={300}/>
                </div>

                <div className={styles.infoContainer}>
                    <h1>@{props.name}</h1>
                    <p>{props.followers} followers</p>
                    <p>{props.following} following</p>
                    <p>{props.postCount} posts</p>
                </div>
            </div>
        </div>
    )
}